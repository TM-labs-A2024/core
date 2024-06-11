package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DoctorsDoctorUuidDelete - Deletes a doctor
func (c *Server) DoctorsDoctorUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsDoctorUuidGet - Returns a single doctor by UUID
func (c *Server) DoctorsDoctorUuidGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsDoctorUuidPatientsGet - Returns a list of patients treated by doctor
func (c *Server) DoctorsDoctorUuidPatientsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsGet - List ALL doctors
func (c *Server) DoctorsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsInstitutionUuidGet - List ALL doctors in an institution
func (c *Server) DoctorsInstitutionUuidGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsLoginPost -
func (c *Server) DoctorsLoginPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsPost - Add a new doctor to the system
func (c *Server) DoctorsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsPut - Update an existing doctor by Id
func (c *Server) DoctorsPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsSpecialtyIdGet - Returns a list of doctors by specialty
func (c *Server) DoctorsSpecialtyIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsSpecialtyIdPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
func (c *Server) DoctorsSpecialtyIdPatientsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
