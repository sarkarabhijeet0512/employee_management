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

func (s *Service) UpsertEmployeeRegistration(dCtx context.Context, req *Employee) error {
	return s.Repo.upsertEmployeeRegistration(dCtx, req)
}

func (s *Service) FetchEmployeeByID(dCtx context.Context, id int) (res *Employee, err error) {
	return s.Repo.fetchEmployeeByID(dCtx, id)
}

func (s *Service) FetchALLEmployeeByFilter(dCtx context.Context, req model.EmployeeFilter) (res []Employee, pagination *model.Pagination, err error) {
	return s.Repo.fetchALLEmployeeByFilter(dCtx, req)
}

func (s *Service) UpdateEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error) {
	return s.Repo.updateEmployeeByID(dCtx, req)
}

func (s *Service) SoftDeleteEmployeeByID(dCtx context.Context, req Employee) (res *Employee, err error) {
	return s.Repo.softDeleteEmployeeByID(dCtx, req)
}
