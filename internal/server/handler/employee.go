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
	"github.com/go-pg/pg/v10"
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
		err = er.New(err, er.InvalidRequestBody).SetStatus(http.StatusBadRequest)
		return
	}
	err = h.employeeService.UpsertEmployeeRegistration(dCtx, req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
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
		err = er.New(err, er.InvalidRequestBody).SetStatus(http.StatusBadRequest)
		return
	}
	data, err := h.employeeService.FetchEmployeeByID(dCtx, employeeID)
	switch err {
	case pg.ErrNoRows:
		err = nil
		res.Message = "No data found"
		res.Success = true
		res.Data = data
		c.JSON(http.StatusOK, res)
		return
	case nil:
		res.Message = "Fetched Sucessfully"
		res.Success = true
		res.Data = data
		c.JSON(http.StatusOK, res)
		return
	default:
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
}

func (h *EmployeeHandler) FetchALLEmployee(c *gin.Context) {
	var (
		err        error
		dCtx       = context.Background()
		req        = model.EmployeeFilter{}
		res        = model.GenericRes{}
		employeeID = 0
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", employeeID).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.InvalidRequestBody).SetStatus(http.StatusBadRequest)
		return
	}
	data, pagination, err := h.employeeService.FetchALLEmployeeByFilter(dCtx, req)
	switch err {
	case pg.ErrNoRows:
		err = nil
		res.Message = "No data found"
		res.Success = true
		res.Data = []employee.Employee{}
		c.JSON(http.StatusOK, res)
		return
	case nil:
		res.Message = "Fetched Sucessfully"
		res.Success = true
		res.Data = data
		res.Meta = pagination
		c.JSON(http.StatusOK, res)
		return
	default:
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	var (
		err  error
		dCtx = context.Background()
		req  = &model.UpdateEmployeeReq{}
		res  = model.GenericRes{}
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("span", req).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.InvalidRequestBody).SetStatus(http.StatusBadRequest)
		return
	}

	employeeID, err := strconv.Atoi(fmt.Sprint(c.Param("id")))
	if err != nil {
		h.log.WithField("span", employeeID).Info("error while converting string to int: " + err.Error())
		err = er.New(err, er.InvalidRequestBody).SetStatus(http.StatusBadRequest)
		return
	}
	newReq := &employee.Employee{
		ID:       employeeID,
		Name:     req.Name,
		Mobile:   req.Mobile,
		Position: req.Position,
		Salary:   req.Salary,
	}
	err = h.employeeService.UpdateEmployeeByID(dCtx, newReq)
	switch err {
	case pg.ErrNoRows:
		err = nil
		res.Message = "Employee Not found"
		res.Success = true
		c.JSON(http.StatusOK, res)
		return
	case nil:
		res.Message = "Updated Sucessfully"
		res.Success = true
		res.Data = req
		c.JSON(http.StatusOK, res)
		return
	default:
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
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
