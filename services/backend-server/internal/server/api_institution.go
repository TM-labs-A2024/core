package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

// InstitutionsApprovedGet - List ALL approved institutions
func (s *Server) InstitutionsApprovedGet(ctx echo.Context) error {
	approvedInstitutions, err := s.Controller.ListApprovedInstitutions()
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

	idToDelete := pgtype.UUID{}
	doctor, err := s.Controller.GetDoctorByID(profID)
	if err == pgx.ErrNoRows {
		nurse, err := s.Controller.GetNurseByID(profID)
		if err != nil {
			s.Logger.Debug(
				"cannot find doctor or nurse by ID",
				"err", err,
				"id", idToDelete,
			)
			return err
		}
		idToDelete = nurse.ID
	} else if err != nil {
		return err
	}

	if doctor.ID.Valid {
		if err := s.Controller.DeleteDoctorByID(idToDelete.Bytes); err != nil {
			return err
		}
	} else {
		if err := s.Controller.DeleteNurseByID(idToDelete.Bytes); err != nil {
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
	institutions, err := s.Controller.ListInstitutions()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, institutions)
}

// InstitutionsGovIdGet - Returns a single institution by gov_id
func (s *Server) InstitutionsGovIdGet(ctx echo.Context) error {
	govID := ctx.Param("govId")
	s.Logger.Debug(govID)

	institution, err := s.Controller.GetInstitutionByGovID(govID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, institution)
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

	institution, err := s.Controller.CreateInstitution(req.Institution)
	if err != nil {
		return err
	}

	req.InstitutionUser.InstitutionID = institution.ID.Bytes

	user, err := s.Controller.CreateInstitutionUser(req.InstitutionUser)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionResponse(institution, user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsPut - Update an existing institutions by Id
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
