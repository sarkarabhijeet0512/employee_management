package handler_test

import (
	"bytes"
	"context"
	"employee_management/config"
	"employee_management/internal/server"
	"employee_management/internal/server/handler"
	"employee_management/pkg/employee"
	"employee_management/utils/initialize"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
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
func TestHealthz(t *testing.T) {
	// router = setupMockServer()
	assert.NotNil(t, router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"ok":"ok"}`, w.Body.String())
}

// func TestEmployeeHandler_GetEmployeeByIDNotFound(t *testing.T) {
// 	h := params.EmployeeHandler
// 	// h := params.EmployeeHandler
// 	router.GET("/employee/:id", func(c *gin.Context) {
// 		h.FetchEmployee(c)
// 	})

// 	// Create a mock request to pass to the handler
// 	req, err := http.NewRequest("GET", "/employee/123", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a ResponseRecorder to record the response
// 	w := httptest.NewRecorder()

// 	// Use the Gin router to serve the request to the ResponseRecorder
// 	router.ServeHTTP(w, req)

// 	// Check the status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// }
func TestEmployeeHandler_GetEmployeeByIDFound(t *testing.T) {
	h := params.EmployeeHandler
	// h := params.EmployeeHandler
	router.GET("/employee/:id", func(c *gin.Context) {
		h.FetchEmployee(c)
	})

	// Create a mock request to pass to the handler
	req, err := http.NewRequest("GET", "/employee/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Use the Gin router to serve the request to the ResponseRecorder
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

}

// func TestEmployeeHandler_FetchALLEmployeeNOResultfound(t *testing.T) {
// 	// h := params.EmployeeHandler
// 	h := params.EmployeeHandler
// 	router.GET("/employee", func(c *gin.Context) {
// 		h.FetchALLEmployee(c)
// 	})

// 	payload := &bytes.Buffer{}
// 	writer := multipart.NewWriter(payload)
// 	_ = writer.WriteField("limit", "-1")
// 	_ = writer.WriteField("id", "3")
// 	err := writer.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// Create a mock request to pass to the handler
// 	req, err := http.NewRequest("GET", "/employee", payload)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a ResponseRecorder to record the response
// 	w := httptest.NewRecorder()

// 	// Use the Gin router to serve the request to the ResponseRecorder
// 	router.ServeHTTP(w, req)

// 	// Check the status code
// 	assert.Equal(t, `{"success":true,"message":"No data found","data":[]}`, w.Body.String())
// }

func TestEmployeeHandler_FetchALLEmployeefound(t *testing.T) {
	// h := params.EmployeeHandler
	h := params.EmployeeHandler
	router.GET("/employee", func(c *gin.Context) {
		h.FetchALLEmployee(c)
	})

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("limit", "-1")
	_ = writer.WriteField("id", "20")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create a mock request to pass to the handler
	req, err := http.NewRequest("GET", "/employee", payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Use the Gin router to serve the request to the ResponseRecorder
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, `{"success":true,"message":"Fetched Sucessfully","data":[{"id":20,"name":"John Doe","mobile":"1234567890","position":"Developer","salary":50000,"is_active":true,"created_at":"2024-06-09T03:30:43.380492+05:30","updated_at":"2024-06-09T03:30:43.380493+05:30","Mutex":null}],"meta":{"current_page":1,"total_pages":1,"total_data_count":1}}`, w.Body.String())
}

// func TestEmployeeHandler_UpdateEmployeeNotFound(t *testing.T) {
// 	// h := params.EmployeeHandler
// 	h := params.EmployeeHandler
// 	router.PUT("/employee/:id", func(c *gin.Context) {
// 		h.UpdateEmployee(c)
// 	})

// 	payload := strings.NewReader(`{
// 		"salary":6800000.00
// 	}`)
// 	// Create a mock request to pass to the handler
// 	req, err := http.NewRequest("PUT", "/employee/1", payload)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Create a ResponseRecorder to record the response
// 	w := httptest.NewRecorder()

// 	// Use the Gin router to serve the request to the ResponseRecorder
// 	router.ServeHTTP(w, req)

//		// Check the status code
//		assert.Equal(t, `{"success":true,"message":"Employee Not found"}`, w.Body.String())
//	}
func TestEmployeeHandler_UpdateEmployeeSuccess(t *testing.T) {
	// h := params.EmployeeHandler
	h := params.EmployeeHandler
	router.PUT("/employee/:id", func(c *gin.Context) {
		h.UpdateEmployee(c)
	})

	payload := strings.NewReader(`{
		"salary":500000
	}`)
	// Create a mock request to pass to the handler
	req, err := http.NewRequest("PUT", "/employee/20", payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Use the Gin router to serve the request to the ResponseRecorder
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, `{"success":true,"message":"Updated Sucessfully","data":{"id":0,"name":"","mobile":"","position":"","salary":0,"is_active":false,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}}`, w.Body.String())
}

// func TestEmployeeHandler_DeleteEmployee(t *testing.T) {
// 	type fields struct {
// 		log             *logrus.Logger
// 		employeeService *employee.Service
// 	}
// 	type args struct {
// 		c *gin.Context
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &EmployeeHandler{
// 				log:             tt.fields.log,
// 				employeeService: tt.fields.employeeService,
// 			}
// 			h.DeleteEmployee(tt.args.c)
// 		})
// 	}
// }
