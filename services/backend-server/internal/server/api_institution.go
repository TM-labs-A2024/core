package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InstitutionsApprovedGet - List ALL approved institutions
func (c *Server) InstitutionsApprovedGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentDoctorUUIDRevokePost - Deny doctor into institution
func (c *Server) InstitutionsEnrollmentDoctorUUIDRevokePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestUUIDApprovePost - Approve doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsEnrollmentRequestUUIDApprovePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestUUIDDenyPost - Deny doctor into institution
func (c *Server) InstitutionsEnrollmentRequestsEnrollmentRequestUUIDDenyPost(ctx echo.Context) error {
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

// InstitutionsGovIdGet - Returns a single institution by gov_id
func (c *Server) InstitutionsGovIdGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUUIDDelete - Delete an institution
func (c *Server) InstitutionsInstitutionUUIDDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// InstitutionsInstitutionUUIDUsersGet - list all institutions users on the system
func (c *Server) InstitutionsInstitutionUUIDUsersGet(ctx echo.Context) error {
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
