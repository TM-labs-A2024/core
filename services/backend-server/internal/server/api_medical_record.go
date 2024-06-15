package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/labstack/echo/v4"
)

// MedicalRecordHealthRecordUUIDDelete - Deletes a medical-record on the DB ONLY
func (c *Server) MedicalRecordHealthRecordUUIDDelete(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

// MedicalRecordHealthRecordUUIDGet - Find medical-record by UUID
func (c *Server) MedicalRecordHealthRecordUUIDGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, db.HealthRecord{})
}

// MedicalRecordPost - Add a new medical-record to the system
func (c *Server) MedicalRecordPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, db.HealthRecord{})
}
