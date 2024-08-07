package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// InstitutionsInstitutionIDUsersGovIDGet - Returns a single institution user by gov id
func (s *Server) InstitutionsInstitutionIDUsersGovIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	govID := ctx.Param("govId")

	user, err := s.Controller.GetInstitutionUserByGovID(institutionID, govID)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionUserResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsInstitutionIDUsersLoginPost -
func (s *Server) InstitutionsInstitutionIDUsersLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	user, err := s.Controller.GetInstitutionUserByLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	token, err := controller.NewClaim(user.ID.Bytes, user.InstitutionID.Bytes)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionUserResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  resp,
	})
}

// InstitutionsInstitutionIDUsersPost - Add a new institutions user to the system
func (s *Server) InstitutionsInstitutionIDUsersPost(ctx echo.Context) error {
	req := models.InstitutionUserPostRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user, err := s.Controller.CreateInstitutionUser(req)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionUserResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsInstitutionIDUsersPut - Update an existing institutions user by ID
func (s *Server) InstitutionsInstitutionIDUsersPut(ctx echo.Context) error {
	req := models.InstitutionUserPutRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user, err := s.Controller.UpdateInstitutionUser(req)
	if err != nil {
		return err
	}

	resp, err := models.NewInstitutionUserResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// InstitutionsInstitutionIDUsersUserIDDelete - Deletes a institution user
func (s *Server) InstitutionsInstitutionIDUsersUserIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	uuidStr = ctx.Param("userId")
	userID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteInstitutionUserByInstitutionAndUserID(institutionID, userID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
