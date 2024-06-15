package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthRecordHealthReacordUUIDDelete - Deletes a health-record on the DB ONLY
func (c *Server) HealthRecordHealthReacordUUIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// HealthRecordHealthReacordUUIDGet - Find health-record by UUID
func (c *Server) HealthRecordHealthReacordUUIDGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// HealthRecordPost - Add a new health-record to the system
func (c *Server) HealthRecordPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
