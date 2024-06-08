package employee

import (
	"context"
	model "employee_management/utils/models"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertEmployeeRegistration(context.Context, *Employee) error
	fetchEmployeeByID(context.Context, int) (*Employee, error)
	fetchALLEmployeeByFilter(context.Context, model.EmployeeFilter) (*Employee, *model.Pagination, error)
	updateEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error)
	softDeleteEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error)
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"employeedb"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	log *logrus.Logger
	db  *pg.DB
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
	_, err := r.db.ModelContext(ctx, req).OnConflict("(mobile,email) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) fetchEmployeeByID(ctx context.Context, id int) (*Employee, error) {

	return nil, nil
}
func (r *PGRepo) fetchALLEmployeeByFilter(dCtx context.Context, req model.EmployeeFilter) (res *Employee, pagination *model.Pagination, err error) {
	return nil, nil, nil
}

func (r *PGRepo) updateEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error) {
	return nil, nil
}

func (r *PGRepo) softDeleteEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error) {
	return nil, nil
}
