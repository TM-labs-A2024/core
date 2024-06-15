package server

import (
	"fmt"
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// DoctorsDoctorUUIDDelete - Deletes a doctor
func (c *Server) DoctorsDoctorUUIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("doctor_uuid")
	doctorUUID, err := uuid.FromBytes([]byte(uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	_ = doctorUUID

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	c.Logger.Debug(fmt.Sprintf("%s deleted by %s", doctorUUID, claims.UserUUID))

	return ctx.NoContent(http.StatusNoContent)
}

// DoctorsDoctorUUIDGet - Returns a single doctor by UUID
func (c *Server) DoctorsDoctorUUIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("doctor_uuid")
	doctorUUID, err := uuid.FromBytes([]byte(uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	_ = doctorUUID

	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsDoctorUUIDPatientsGet - Returns a list of patients treated by doctor
func (c *Server) DoctorsDoctorUUIDPatientsGet(ctx echo.Context) error {
	uuidStr := ctx.Param("doctor_uuid")
	doctorUUID, err := uuid.FromBytes([]byte(uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	_ = doctorUUID

	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsGet - List ALL doctors
func (c *Server) DoctorsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsInstitutionUUIDGet - List ALL doctors in an institution
func (c *Server) DoctorsInstitutionUUIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("institution_uuid")
	doctorUUID, err := uuid.FromBytes([]byte(uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	_ = doctorUUID

	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsLoginPost -
func (c *Server) DoctorsLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// TODO: login and obtain doctor UUID

	token, err := controller.NewClaim(uuid.New())
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// DoctorsPost - Add a new doctor to the system
func (c *Server) DoctorsPost(ctx echo.Context) error {
	request := models.DoctorsPostRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsPut - Update an existing doctor by Id
func (c *Server) DoctorsPut(ctx echo.Context) error {
	request := models.DoctorsPutRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsSpecialtyIdGet - Returns a list of doctors by specialty
func (c *Server) DoctorsSpecialtyIdGet(ctx echo.Context) error {
	request := models.DoctorsPostRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}

// DoctorsSpecialtyIdPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
func (c *Server) DoctorsSpecialtyIdPatientsGet(ctx echo.Context) error {
	request := models.DoctorsPostRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}
