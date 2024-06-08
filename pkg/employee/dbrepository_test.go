package employee_test

import (
	"context"
	"employee_management/config"
	"employee_management/internal/server"
	"employee_management/internal/server/handler"
	"employee_management/pkg/employee"
	"employee_management/utils"
	"employee_management/utils/initialize"
	model "employee_management/utils/models"
	"fmt"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
	),
)

type (
	// User represents the user entity
	Employee struct {
		ID        int       `json:"id" pg:"id,pk"`
		Name      string    `json:"name" pg:"name" binding:"required"`
		Mobile    string    `json:"mobile" pg:"mobile,unique" binding:"required"`
		Position  string    `json:"position" pg:"position"`
		Salary    float64   `json:"salary" pg:"salary"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
		// Mutex     *sync.Mutex
	}
)

var router *gin.Engine
var params server.Options

func setupMockServer() (router *gin.Engine) {
	gin.SetMode(gin.TestMode)
	app := fx.New(
		fx.Provide(
			initialize.NewDB,
		),
		config.Module,
		initialize.Module,
		handler.Module,
		server.Module,
		employee.Module,
		Module,
		// Run app forever
		fx.Populate(&params),
	)
	app.Start(context.TODO())
	defer app.Stop(context.TODO())
	router = server.SetupRouter(&params)
	return
}
func init() {
	router = setupMockServer()
}
func setupTest() *PGRepo {
	log := logrus.New()
	repo := &PGRepo{
		Log: log,
		Db:  params.PostgresDB,
	}
	return repo
}

type Repository interface {
	upsertEmployeeRegistration(context.Context, *Employee) error
	fetchEmployeeByID(context.Context, int) (*Employee, error)
	fetchALLEmployeeByFilter(context.Context, model.EmployeeFilter) ([]Employee, *model.Pagination, error)
	softDeleteEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error)
	updateEmployeeByID(ctx context.Context, employee *Employee) error
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"employeedb"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	Log   *logrus.Logger
	Db    *pg.DB
	Mutex sync.RWMutex
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {

	Repo = &PGRepo{
		Log: i.Log,
		Db:  i.DB,
	}

	return
}

func (r *PGRepo) upsertEmployeeRegistration(ctx context.Context, req *Employee) error {
	utils.SetGenericFieldValue(req)
	_, err := r.Db.ModelContext(ctx, req).OnConflict("(mobile) DO UPDATE").Insert()
	return err
}
func (r *PGRepo) fetchEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	employee := &Employee{}
	err := r.Db.ModelContext(ctx, employee).Where("id=?", id).Select()
	return employee, err
}

func (r *PGRepo) fetchALLEmployeeByFilter(ctx context.Context, filter model.EmployeeFilter) ([]Employee, *model.Pagination, error) {
	var (
		employees []Employee
		p         model.Pagination
		count     int
		err       error
		emptyTime time.Time
	)
	query := r.Db.ModelContext(ctx, &employees)
	if filter.ID != 0 {
		query.Where("id=?", filter.ID)
	}
	if filter.Name != "" {
		query.Where(`name ILIKE '%` + filter.Name + `%'`)
	}
	if filter.Position != "" {
		query.Where(`position ILIKE '%` + filter.Position + `%'`)
	}
	if filter.CreatedAt != emptyTime {
		query.Where("date(created_at)=?", filter.CreatedAt.Format("2006-01-02"))
	}
	if filter.IsActive != nil {
		query.Where("is_active=?", filter.IsActive)
	}
	if filter.Salary != 0 {
		query.Where("salary=?", filter.Salary)
	}
	if filter.Limit != -1 {
		count, err = query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).Order("id desc").SelectAndCount(&employees)
	} else {
		err = query.Select()
	}
	if err != nil {
		return nil, &p, err
	}
	//hack as orm not returning error in case of array of struct
	if len(employees) == 0 {
		return nil, &p, pg.ErrNoRows
	}
	p.TotalDataCount = count
	p.CurrentPage = filter.Page
	p.TotalPages = int(math.Ceil(float64(count) / float64(filter.Limit)))
	return employees, &p, err
}

func (r *PGRepo) updateEmployeeByID(ctx context.Context, employee *Employee) error {

	query := r.Db.ModelContext(ctx, employee)
	if employee.Name != "" {
		query.Set("name =? ", employee.Name)
	}
	if employee.Position != "" {
		query.Set("position =?", employee.Position)
	}
	if employee.Mobile != "" {
		query.Set("mobile =?", employee.Mobile)
	}
	if employee.Salary != 0 {
		query.Set("salary=?", employee.Salary)
	}
	query.Set("updated_at=?", time.Now())
	r.Mutex.Lock()
	_, err := query.Where("id=?", employee.ID).Returning("*").Update()
	r.Mutex.Unlock()
	return err
}

func (r *PGRepo) softDeleteEmployeeByID(ctx context.Context, req Employee) (*Employee, error) {
	employee := &Employee{}
	r.Mutex.Lock()
	_, err := r.Db.ModelContext(ctx, &employee).Set("is_active=?", false).Where("id = ?", employee.ID).Update(&req)
	r.Mutex.Unlock()
	return employee, err
}

func TestUpsertEmployeeRegistration(t *testing.T) {
	repo := setupTest()
	tx, err := repo.Db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	// defer tx.Close()

	ctx := context.Background()
	// defer repo.Db.Close()
	req := &Employee{
		Name:     "John Doe",
		Position: "Developer",
		Mobile:   "1234567890",
		Salary:   50000,
		IsActive: true,
	}
	fmt.Println("Before upsertEmployeeRegistration")
	err = repo.upsertEmployeeRegistration(ctx, req)
	if err != nil {
		t.Fatalf("Error upserting employee registration: %v", err)
	}
	fmt.Println("After upsertEmployeeRegistration")
	tx.Rollback()
}

func TestFetchALLEmployeeByFilter(t *testing.T) {
	repo := setupTest()
	tx, err := repo.Db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Close()

	ctx := context.Background()
	// defer repo.Db.Close()
	req := model.EmployeeFilter{
		Name:     "John Doe",
		Position: "Developer",
		Salary:   50000,
	}
	fmt.Println("Before upsertEmployeeRegistration")
	data, pagination, err := repo.fetchALLEmployeeByFilter(ctx, req)
	if err != nil {
		t.Fatalf("Error upserting employee registration: %v", err)
	}
	if len(data) == 0 {
		t.Error("Expected at least one employee, but got none.")
	}

	if pagination.TotalPages <= 0 {
		t.Error("Expected positive total pages, but got non-positive value.")
	}
	fmt.Println("After upsertEmployeeRegistration")
	tx.Rollback()
}
