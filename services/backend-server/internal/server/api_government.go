package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GovermentLoginPost -
func (c *Server) GovermentLoginPost(ctx echo.Context) error {
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

// GovernmentEnrollmentInstitutionUUIDRevokePost - Deny institution into the system
func (c *Server) GovernmentEnrollmentInstitutionUUIDRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsEnrollmentRequestUUIDApprovePost - Approve institution into the system
func (c *Server) GovernmentEnrollmentRequestsEnrollmentRequestUUIDApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsEnrollmentRequestUUIDDenyPost - Deny institution into the system
func (c *Server) GovernmentEnrollmentRequestsEnrollmentRequestUUIDDenyPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsGet - List request to approve institution into government
func (c *Server) GovernmentEnrollmentRequestsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsPost - Send request to approve institution into government
func (c *Server) GovernmentEnrollmentRequestsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
