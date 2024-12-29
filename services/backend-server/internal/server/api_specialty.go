package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/labstack/echo/v4"
)

// SpecialtiesGet - Returns a list of specialties
func (c *Server) SpecialtiesGet(ctx echo.Context) error {
	specialties, err := c.Controller.ListSpecialties()
	if err != nil {
		return err
	}

	resp := models.NewSpecialtiesResponse(specialties)

	return ctx.JSON(http.StatusOK, resp)
}
