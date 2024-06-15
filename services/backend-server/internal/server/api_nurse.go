package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// NursesGet - List ALL nurses
func (c *Server) NursesGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesInstitutionUUIDGet - List ALL nurses in an institution
func (c *Server) NursesInstitutionUUIDGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesLoginPost -
func (c *Server) NursesLoginPost(ctx echo.Context) error {
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

// NursesNurseUUIDDelete - Deletes a nurse
func (c *Server) NursesNurseUUIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// NursesNurseUUIDGet - Find nurse by UUID
func (c *Server) NursesNurseUUIDGet(ctx echo.Context) error {
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
