package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MedicalRecordHealthRecordUuidDelete - Deletes a medical-record on the DB ONLY
func (c *Server) MedicalRecordHealthRecordUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// MedicalRecordHealthRecordUuidGet - Find medical-record by UUID
func (c *Server) MedicalRecordHealthRecordUuidGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// MedicalRecordPost - Add a new medical-record to the system
func (c *Server) MedicalRecordPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
