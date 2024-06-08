package server

import (
	"employee_management/internal/server/mw"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/v1/api/")
	// Authentication apis
	r.Use(mw.ErrorHandlerX(o.Log))
	r.POST("/employee", o.EmployeeHandler.EmployeeRegistration)
	r.GET("/employee/:id", o.EmployeeHandler.FetchEmployee)
	r.GET("/employee", o.EmployeeHandler.FetchALLEmployee)
	r.PUT("/employee/:id", o.EmployeeHandler.UpdateEmployee)
	r.DELETE("/employee/:id", o.EmployeeHandler.DeleteEmployee)
}
