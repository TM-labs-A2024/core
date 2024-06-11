package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PatientsAccessDoctorUuidRevokePost - Deny doctor access to patient records
func (c *Server) PatientsAccessDoctorUuidRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsAccessRequestsAccessRequestUuidApprovePost - Approve doctor access to patient records
func (c *Server) PatientsAccessRequestsAccessRequestUuidApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsAccessRequestsAccessRequestUuidDenyPost - Deny doctor access to patient records
func (c *Server) PatientsAccessRequestsAccessRequestUuidDenyPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsAccessRequestsGet - List requests from doctors to access patient records
func (c *Server) PatientsAccessRequestsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGet - List ALL patients
func (c *Server) PatientsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdDoctorsGet - Returns a list of doctors treating patients
func (c *Server) PatientsGovIdDoctorsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdGet - Find patient by govId
func (c *Server) PatientsGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdMedicalRecordsGet - List health records by patient
func (c *Server) PatientsGovIdMedicalRecordsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdOrdersGet - List health orders by patient
func (c *Server) PatientsGovIdOrdersGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsLoginPost -
func (c *Server) PatientsLoginPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsPatientUuidAccessRequestsPost - Make request for doctor to access patient records
func (c *Server) PatientsPatientUuidAccessRequestsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsPatientUuidDelete - Deletes a patient
func (c *Server) PatientsPatientUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsPost - Add a new patient to the system
func (c *Server) PatientsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsPut - Update an existing patient by uuid
func (c *Server) PatientsPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
