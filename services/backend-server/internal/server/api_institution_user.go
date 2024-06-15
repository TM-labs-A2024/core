package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// InstitutionsInstitutionUUIDUsersGovIdGet - Returns a single institution user by gov id
func (c *Server) InstitutionsInstitutionUUIDUsersGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUUIDUsersLoginPost -
func (c *Server) InstitutionsInstitutionUUIDUsersLoginPost(ctx echo.Context) error {
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

// InstitutionsInstitutionUUIDUsersPost - Add a new institutions user to the system
func (c *Server) InstitutionsInstitutionUUIDUsersPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUUIDUsersPut - Update an existing institutions user by Id
func (c *Server) InstitutionsInstitutionUUIDUsersPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUUIDUsersUserUUIDDelete - Deletes a institution user
func (c *Server) InstitutionsInstitutionUUIDUsersUserUUIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
