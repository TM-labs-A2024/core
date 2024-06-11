package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InstitutionsInstitutionUuidUsersGovIdGet - Returns a single institution user by gov id
func (c *Server) InstitutionsInstitutionUuidUsersGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidUsersLoginPost -
func (c *Server) InstitutionsInstitutionUuidUsersLoginPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidUsersPost - Add a new institutions user to the system
func (c *Server) InstitutionsInstitutionUuidUsersPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidUsersPut - Update an existing institutions user by Id
func (c *Server) InstitutionsInstitutionUuidUsersPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidUsersUserUuidDelete - Deletes a institution user
func (c *Server) InstitutionsInstitutionUuidUsersUserUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
