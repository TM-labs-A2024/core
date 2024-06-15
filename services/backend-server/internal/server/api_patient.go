package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// PatientsAccessDoctorUUIDRevokePost - Deny doctor access to patient records
func (c *Server) PatientsAccessDoctorUUIDRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsAccessRequestsAccessRequestUUIDApprovePost - Approve doctor access to patient records
func (c *Server) PatientsAccessRequestsAccessRequestUUIDApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsAccessRequestsAccessRequestUUIDDenyPost - Deny doctor access to patient records
func (c *Server) PatientsAccessRequestsAccessRequestUUIDDenyPost(ctx echo.Context) error {
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

// PatientsGovIdGet - Find patient by gov_id
func (c *Server) PatientsGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdHealthRecordsGet - List health records by patient
func (c *Server) PatientsGovIdHealthRecordsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdHealthRecordsSpecialtiesGet - List health records by patient and specialty Id
func (c *Server) PatientsGovIdHealthRecordsSpecialtiesGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdHealthRecordsSpecialtyIdGet - List health records by patient and specialty Id
func (c *Server) PatientsGovIdHealthRecordsSpecialtyIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsGovIdOrdersGet - List health orders by patient
func (c *Server) PatientsGovIdOrdersGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsLoginPost -
func (c *Server) PatientsLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	// TODO: login and obtain doctor UUID

	token, err := controller.NewClaim(uuid.New())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// PatientsPatientUUIDAccessRequestsPost - Make request for doctor to access patient records
func (c *Server) PatientsPatientUUIDAccessRequestsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// PatientsPatientUUIDDelete - Deletes a patient
func (c *Server) PatientsPatientUUIDDelete(ctx echo.Context) error {
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
