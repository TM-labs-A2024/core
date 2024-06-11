package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NursesGet - List ALL nurses
func (c *Server) NursesGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesInstitutionUuidGet - List ALL nurses in an institution
func (c *Server) NursesInstitutionUuidGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesLoginPost -
func (c *Server) NursesLoginPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesNurseUuidDelete - Deletes a nurse
func (c *Server) NursesNurseUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesNurseUuidGet - Find nurse by UUID
func (c *Server) NursesNurseUuidGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesPost - Add a new nurse to the system
func (c *Server) NursesPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesPut - Update an existing nurse by UUID
func (c *Server) NursesPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
