package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InstitutionsEnrollmentDoctorUuidRevokePost - Deny doctor into institution
func (c *Server) InstitutionsEnrollmentDoctorUuidRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestUuidApprovePost - Approve doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsEnrollmentRequestUuidApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestUuidDenyPost - Deny doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsEnrollmentRequestUuidDenyPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsGet - List request to approve doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsPost - Send request to approve doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsGet - List ALL institutions
func (c *Server) InstitutionsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsGovIdGet - Returns a single institution by govId
func (c *Server) InstitutionsGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidDelete - Delete an institution
func (c *Server) InstitutionsInstitutionUuidDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUuidUsersGet - list all institutions users on the system
func (c *Server) InstitutionsInstitutionUuidUsersGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsPost - Add a new institutions to the system
func (c *Server) InstitutionsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsPut - Update an existing institutions by Id
func (c *Server) InstitutionsPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
