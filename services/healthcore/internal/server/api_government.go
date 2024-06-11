package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GovermentLoginPost -
func (c *Server) GovermentLoginPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentInstitutionUuidRevokePost - Deny institution into the system
func (c *Server) GovernmentEnrollmentInstitutionUuidRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsEnrollmentRequestUuidApprovePost - Approve institution into the system
func (c *Server) GovernmentEnrollmentRequestsEnrollmentRequestUuidApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// GovernmentEnrollmentRequestsEnrollmentRequestUuidDenyPost - Deny institution into the system
func (c *Server) GovernmentEnrollmentRequestsEnrollmentRequestUuidDenyPost(ctx echo.Context) error {
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
