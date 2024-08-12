package server

import (
	"fmt"
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// InstitutionsApprovedGet - List ALL approved institutions
func (s *Server) InstitutionsApprovedGet(ctx echo.Context) error {
	approvedInstitutions, err := s.Controller.ListInstitutions(true)
	if err != nil {
		return err
	}

	resp := []models.InstitutionWithUserResponse{}
	for _, institution := range approvedInstitutions {
		user, err := s.Controller.GetFirstInstitutionUser(institution.ID.Bytes)
		if err != nil {
			return err
		}

		instResp, err := models.NewInstitutionResponse(institution, user)
		if err != nil {
			return err
		}

		resp = append(resp, instResp)
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsEnrollmentDoctorIDRevokePost - Deny doctor into institution
func (s *Server) InstitutionsEnrollmentDoctorIDRevokePost(ctx echo.Context) error {
	uuidStr := ctx.Param("professionalId")
	profID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteInstitutionEnrollmentRequestsByProfID(profID); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	doctor, doctorErr := s.Controller.GetDoctorByID(profID)
	if doctorErr == nil {
		if err := s.Controller.DeleteDoctorByID(doctor.ID.Bytes); err != nil {
			return err
		}
	} else {
		nurse, err := s.Controller.GetNurseByID(profID)
		if err != nil {
			return fmt.Errorf("prof id cannot be fetch by neither doctor:%w  nor nurse: %w", doctorErr, err)
		}

		if err := s.Controller.DeleteNurseByID(nurse.ID.Bytes); err != nil {
			return err
		}
	}

	return ctx.NoContent(http.StatusNoContent)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestIDApprovePost - Approve doctor into institution
func (s *Server) InstitutionsEnrollmentRequestsEnrollmentRequestIDApprovePost(ctx echo.Context) error {
	uuidStr := ctx.Param("enrollmentRequestId")
	erID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	erByID, err := s.Controller.GetInstitutionEnrollmentRequestsByID(erID)
	if err != nil {
		return err
	}

	er, err := s.Controller.ApproveInstitutionEnrollmentRequestsByID(erByID)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionsEnrollmentRequestsResponse(er)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsEnrollmentRequestsEnrollmentRequestIDDenyPost - Deny doctor into institution
func (s *Server) InstitutionsEnrollmentRequestsEnrollmentRequestIDDenyPost(ctx echo.Context) error {
	uuidStr := ctx.Param("enrollmentRequestId")
	erID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	erByID, err := s.Controller.GetInstitutionEnrollmentRequestsByID(erID)
	if err != nil {
		return err
	}

	er, err := s.Controller.DenyInstitutionEnrollmentRequestsByID(erByID)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionsEnrollmentRequestsResponse(er)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsEnrollmentRequestsGet - List request to approve doctor into institution
func (s *Server) InstitutionsEnrollmentRequestsGet(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	iers, err := s.Controller.ListInstitutionsEnrollmentRequestsByInstitutionID(
		claims.InstitutionID,
	)
	if err != nil {
		return err
	}

	resps := []models.InstitutionsEnrollmentRequestsResponse{}
	for _, ier := range iers {
		resp, err := models.NewInstitutionsEnrollmentRequestsResponse(ier)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// InstitutionsEnrollmentRequestsPost - Send request to approve doctor into institution
// func (s *Server) InstitutionsEnrollmentRequestsPost(ctx echo.Context) error {
// 	// This will be implemented at controller level only.
// }

// InstitutionsGet - List ALL institutions
func (s *Server) InstitutionsGet(ctx echo.Context) error {
	institutions, err := s.Controller.ListInstitutions(false)
	if err != nil {
		return err
	}

	resp := []models.InstitutionWithUserResponse{}
	for _, institution := range institutions {
		user, err := s.Controller.GetFirstInstitutionUser(institution.ID.Bytes)
		if err != nil {
			return err
		}

		instResp, err := models.NewInstitutionResponse(institution, user)
		if err != nil {
			return err
		}

		resp = append(resp, instResp)
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsIDGet - Returns a single institution by _id
func (s *Server) InstitutionsIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	institution, err := s.Controller.GetInstitutionByID(institutionID)
	if err != nil {
		return err
	}

	user, err := s.Controller.GetFirstInstitutionUser(institutionID)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionResponse(institution, user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsInstitutionIDDelete - Delete an institution
func (s *Server) InstitutionsInstitutionIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteInstitutionByID(institutionID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// InstitutionsInstitutionIDUsersGet - list all institutions users on the system
func (s *Server) InstitutionsInstitutionIDUsersGet(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	institutionsUsers, err := s.Controller.ListInstitutionUsersByInstitutionID(institutionID)
	if err != nil {
		return err
	}

	resps := []models.InstitutionUserResponse{}
	for _, user := range institutionsUsers {
		resp, err := models.NewInstitutionUserResponse(user)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// InstitutionsPost - Add a new institutions to the system
func (s *Server) InstitutionsPost(ctx echo.Context) error {
	req := models.InstitutionWithUserRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	institution, user, err := s.Controller.CreateInstitution(
		req.Institution,
		req.InstitutionUser,
	)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionResponse(institution, user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsPut - Update an existing institutions by ID
func (s *Server) InstitutionsPut(ctx echo.Context) error {
	req := models.InstitutionWithID{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	institution, err := s.Controller.UpdateInstitution(req.Institution, req.ID)
	if err != nil {
		return err
	}

	user, err := s.Controller.GetFirstInstitutionUser(req.ID)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionResponse(institution, user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) InstitutionsPatientsGet(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	patients, err := s.Controller.ListPatientsByInstitutionID(claims.InstitutionID)
	if err != nil {
		return err
	}

	resps := []models.PatientResponse{}
	for _, patient := range patients {
		resp, err := models.NewPatientResponse(patient)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}
