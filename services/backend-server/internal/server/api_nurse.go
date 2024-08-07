package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// NursesGet - List ALL nurses
func (s *Server) NursesGet(ctx echo.Context) error {
	nurses, err := s.Controller.ListNurses()
	if err != nil {
		return err
	}

	resps := []models.NursesResponse{}
	for _, nurse := range nurses {
		resp, err := models.NewNurseResponse(nurse)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}
	return ctx.JSON(http.StatusOK, resps)
}

// NursesInstitutionIDGet - List ALL nurses in an institution
func (s *Server) NursesInstitutionIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	nurses, err := s.Controller.ListNursesByInstitutionID(institutionID)
	if err != nil {
		return err
	}

	resps := []models.NursesResponse{}
	for _, nurse := range nurses {
		resp, err := models.NewNurseResponse(nurse)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}
	return ctx.JSON(http.StatusOK, resps)
}

// NursesLoginPost -
func (s *Server) NursesLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	user, err := s.Controller.GetNurseByLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	token, err := controller.NewClaim(user.ID.Bytes, user.InstitutionID.Bytes)
	if err != nil {
		return err
	}

	resp, err := models.NewNurseResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
		"nurse": resp,
	})
}

// NursesNurseIDDelete - Deletes a nurse
func (s *Server) NursesNurseIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("nurseId")
	nurseID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteNurseByID(nurseID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// NursesNurseIDGet - Find nurse by ID
func (s *Server) NursesNurseIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("nurseId")
	nurseID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	nurse, err := s.Controller.GetNurseByID(nurseID)
	if err != nil {
		return err
	}

	resp, err := models.NewNurseResponse(nurse)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// NursesPost - Add a new nurse to the system
func (s *Server) NursesPost(ctx echo.Context) error {
	req := models.NursePostRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user, err := s.Controller.CreateNurse(req)
	if err != nil {
		return err
	}

	resp, err := models.NewNurseResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// NursesPut - Update an existing nurse by ID
func (s *Server) NursesPut(ctx echo.Context) error {
	req := models.NursesPutRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user, err := s.Controller.UpdateNurse(req)
	if err != nil {
		return err
	}

	resp, err := models.NewNurseResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}
