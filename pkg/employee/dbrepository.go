package employee

import (
	"context"
	"employee_management/utils"
	model "employee_management/utils/models"
	"math"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

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
	log   *logrus.Logger
	db    *pg.DB
	mutex sync.RWMutex
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {

	Repo = &PGRepo{
		log: i.Log,
		db:  i.DB,
	}

	return
}

func (r *PGRepo) upsertEmployeeRegistration(ctx context.Context, req *Employee) error {
	utils.SetGenericFieldValue(req)
	_, err := r.db.ModelContext(ctx, req).OnConflict("(mobile) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) fetchEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	employee := &Employee{}
	err := r.db.ModelContext(ctx, employee).Where("id=?", id).Select()
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
	query := r.db.ModelContext(ctx, &employees)
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

	query := r.db.ModelContext(ctx, employee)
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
	r.mutex.Lock()
	_, err := query.Where("id=?", employee.ID).Returning("*").Update()
	r.mutex.Unlock()
	return err
}

func (r *PGRepo) softDeleteEmployeeByID(ctx context.Context, req Employee) (*Employee, error) {
	employee := &Employee{}
	r.mutex.Lock()
	_, err := r.db.ModelContext(ctx, &employee).Set("is_active=?", false).Where("id = ?", employee.ID).Update(&req)
	r.mutex.Unlock()
	return employee, err
}
