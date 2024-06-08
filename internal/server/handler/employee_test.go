package handler

import (
	"employee_management/pkg/employee"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestEmployeeHandler_EmployeeRegistration(t *testing.T) {
	type fields struct {
		log             *logrus.Logger
		employeeService *employee.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EmployeeHandler{
				log:             tt.fields.log,
				employeeService: tt.fields.employeeService,
			}
			h.EmployeeRegistration(tt.args.c)
		})
	}
}

func TestEmployeeHandler_FetchEmployee(t *testing.T) {
	type fields struct {
		log             *logrus.Logger
		employeeService *employee.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EmployeeHandler{
				log:             tt.fields.log,
				employeeService: tt.fields.employeeService,
			}
			h.FetchEmployee(tt.args.c)
		})
	}
}

func TestEmployeeHandler_FetchALLEmployee(t *testing.T) {
	type fields struct {
		log             *logrus.Logger
		employeeService *employee.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EmployeeHandler{
				log:             tt.fields.log,
				employeeService: tt.fields.employeeService,
			}
			h.FetchALLEmployee(tt.args.c)
		})
	}
}

func TestEmployeeHandler_UpdateEmployee(t *testing.T) {
	type fields struct {
		log             *logrus.Logger
		employeeService *employee.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EmployeeHandler{
				log:             tt.fields.log,
				employeeService: tt.fields.employeeService,
			}
			h.UpdateEmployee(tt.args.c)
		})
	}
}

func TestEmployeeHandler_DeleteEmployee(t *testing.T) {
	type fields struct {
		log             *logrus.Logger
		employeeService *employee.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &EmployeeHandler{
				log:             tt.fields.log,
				employeeService: tt.fields.employeeService,
			}
			h.DeleteEmployee(tt.args.c)
		})
	}
}
