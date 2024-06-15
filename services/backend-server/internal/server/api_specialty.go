package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SpecialtiesGet - Returns a list of specialties
func (c *Server) SpecialtiesGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
