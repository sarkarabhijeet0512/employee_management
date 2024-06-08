package employee

import (
	"context"
	model "employee_management/utils/models"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	Repo Repository
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository) *Service {
	return &Service{
		conf: conf,
		log:  log,
		Repo: Repo,
	}
}

// The UpsertEmployeeRegistration function is a method within a service that either inserts
// a new employee registration or updates an existing one if mobile number is already registered,
// based on the provided employee details.It interacts with a dbrepository to perform database
// operations and returns an error if any issues arise during the process.
func (s *Service) UpsertEmployeeRegistration(ctx context.Context, employee *Employee) error {
	return s.Repo.upsertEmployeeRegistration(ctx, employee)
}

// The FetchEmployeeByID function, part of a service, retrieves an employee's information by
// their unique identifier. It utilizes a dbrepository to execute the database query, returning
// the employee's data or an error if the operation fails.
func (s *Service) FetchEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	return s.Repo.fetchEmployeeByID(ctx, id)
}

// The FetchALLEmployeeByFilter function, belonging to a service, retrieves multiple employees
// based on a provided filters. It utilizes a dbrepository to execute the query and returns a slice
// of employee data along with pagination information or an error if the operation fails.
func (s *Service) FetchALLEmployeeByFilter(ctx context.Context, employeeFilter model.EmployeeFilter) ([]Employee, *model.Pagination, error) {
	return s.Repo.fetchALLEmployeeByFilter(ctx, employeeFilter)
}

// The UpdateEmployeeByID function, part of a service, updates an existing employee's information
// using their unique identifier. It interacts with a dbrepository to execute the update operation
// and returns an error if the process encounters any issues.
func (s *Service) UpdateEmployeeByID(ctx context.Context, employee *Employee) error {
	return s.Repo.updateEmployeeByID(ctx, employee)
}

// The SoftDeleteEmployeeByID function, within a service, performs a soft delete operation on an
// employee record based on the provided employee details. It interacts with a dbrepository to
// execute the soft delete operation and returns either the deleted employee data or an error if
// the operation encounters any issues.
func (s *Service) SoftDeleteEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	return s.Repo.softDeleteEmployeeByID(ctx, id)
}
