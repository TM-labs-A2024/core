package server

import (
	"log"
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GovermentLoginPost -
func (s *Server) GovermentLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	user, err := s.Controller.GetGovernmentByLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	token, err := controller.NewClaim(user.ID.Bytes, uuid.UUID{})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
		"id":    user.ID,
	})
}

// GovernmentEnrollmentInstitutionIDRevokePost - Deny institution into the system
func (s *Server) GovernmentEnrollmentInstitutionIDRevokePost(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	instID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteInstitutionByID(instID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

// GovernmentEnrollmentRequestsEnrollmentRequestIDApprovePost - Approve institution into the system
func (s *Server) GovernmentEnrollmentRequestsEnrollmentRequestIDApprovePost(ctx echo.Context) error {
	uuidStr := ctx.Param("enrollmentRequestId")
	log.Println(uuidStr)
	erID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	er, err := s.Controller.ApproveGovernmentEnrollmentRequest(erID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.NewGovernmentEnrollmentRequest(er))
}

// GovernmentEnrollmentRequestsEnrollmentRequestIDDenyPost - Deny institution into the system
func (s *Server) GovernmentEnrollmentRequestsEnrollmentRequestIDDenyPost(ctx echo.Context) error {
	uuidStr := ctx.Param("enrollmentRequestId")
	erID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	er, err := s.Controller.DenyGovernmentEnrollmentRequest(erID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.NewGovernmentEnrollmentRequest(er))
}

// GovernmentEnrollmentRequestsGet - List request to approve institution into government
func (s *Server) GovernmentEnrollmentRequestsGet(ctx echo.Context) error {
	ers, err := s.Controller.ListGovernmentEnrollmentRequest()
	if err != nil {
		return err
	}

	resps := []models.GovernmentEnrollmentRequest{}
	for _, er := range ers {
		resp := models.NewGovernmentEnrollmentRequest(er)
		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// GovernmentEnrollmentRequestsPost - Send request to approve institution into government
// func (s *Server) GovernmentEnrollmentRequestsPost(ctx echo.Context) error {
// 	return ctx.JSON(http.StatusOK, nil)
// }
