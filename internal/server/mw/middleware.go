// Package mw is user Middleware package
package mw

import (
	"net/http"

	"employee_management/er"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandlerX(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := c.Errors.Last()
			if err == nil {
				// no errors, abort with success
				return
			}

			e := er.From(err.Err)

			// if !e.NOP {
			// 	sentry.CaptureException(e)
			// }

			httpStatus := http.StatusInternalServerError
			if e.Status > 0 {
				httpStatus = e.Status
			}
			// Respond with the custom error message
			c.JSON(httpStatus, e)
		}()
		c.Next()
	}
}
