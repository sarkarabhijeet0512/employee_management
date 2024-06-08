package handler

import (
	"context"
	"employee_management/er"
	"employee_management/pkg/employee"
	"fmt"
	"net/http"
	"strconv"

	model "employee_management/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EmployeeHandler struct {
	log             *logrus.Logger
	employeeService *employee.Service
}

func newEmployeeHandler(
	log *logrus.Logger,
	employeeService *employee.Service,
) *EmployeeHandler {
	return &EmployeeHandler{
		log,
		employeeService,
	}
}

func (h *EmployeeHandler) EmployeeRegistration(c *gin.Context) {
	var (
		err  error
		res  = model.GenericRes{}
		req  = &employee.Employee{}
		dCtx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", res).Warn(err.Error())
			return
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	err = h.employeeService.UpsertEmployeeRegistration(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Registration Sucessfully Done"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}

func (h *EmployeeHandler) FetchEmployee(c *gin.Context) {
	var (
		err        error
		dCtx       = context.Background()
		res        = model.GenericRes{}
		employeeID = 0
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", employeeID).Warn(err.Error())
		}
	}()

	employeeID, err = strconv.Atoi(fmt.Sprint(c.Param("id")))
	if err != nil {
		h.log.WithField("span", employeeID).Info("error while converting string to int: " + err.Error())
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	data, err := h.employeeService.FetchEmployeeByID(dCtx, employeeID)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Message = "Registration Sucessfully Done"
	res.Success = true
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (h *EmployeeHandler) FetchALLEmployee(c *gin.Context) {
	var (
		err        error
		dCtx       = context.Background()
		res        = model.GenericRes{}
		employeeID = 0
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", employeeID).Warn(err.Error())
		}
	}()

	employeeID, err = strconv.Atoi(fmt.Sprint(c.Param("id")))
	if err != nil {
		h.log.WithField("span", employeeID).Info("error while converting string to int: " + err.Error())
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	data, err := h.employeeService.FetchEmployeeByID(dCtx, employeeID)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Message = "Fetched Sucessfully"
	res.Success = true
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	var (
		err  error
		dCtx = context.Background()
		req  = employee.Employee{}
		res  = model.GenericRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", req).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	data, err := h.employeeService.UpdateEmployeeByID(dCtx, req)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Message = "Updated Successfully"
	res.Success = true
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	var (
		err  error
		dCtx = context.Background()
		req  = employee.Employee{}
		res  = model.GenericRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", req).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	Data, err := h.employeeService.SoftDeleteEmployeeByID(dCtx, req)
	if err != nil {
		err = er.New(err, er.UserNotFound).SetStatus(http.StatusNotFound)
		return
	}
	res.Message = "Sucessfully Deleted"
	res.Success = true
	res.Data = Data
	c.JSON(http.StatusOK, res)
}
