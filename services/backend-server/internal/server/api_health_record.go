package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

const (
	storagePath = "./records"
)

// HealthRecordHealthReacordIDDelete - Deletes a health-record on the DB ONLY
func (s *Server) HealthRecordHealthReacordIDDelete(ctx echo.Context) error {
	healthReacordIdStr := ctx.Param("healthRecordId")
	healthReacordId, err := uuid.Parse(healthReacordIdStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if err := s.Controller.DeleteHealthRecordByID(healthReacordId); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

// HealthRecordHealthReacordIDGet - Find health-record by ID
func (s *Server) HealthRecordHealthReacordIDGet(ctx echo.Context) error {
	healthReacordIdStr := ctx.Param("healthRecordId")
	healthReacordId, err := uuid.Parse(healthReacordIdStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	hr, err := s.Controller.GetHealthRecordByID(healthReacordId)
	if err != nil {
		return err
	}

	spacialty, err := s.Controller.GetSpecialtyByID(hr.SpecialtyID.Bytes)
	if err != nil {
		return err
	}

	resp := models.NewHealthRecordResponse(hr, spacialty)

	return ctx.JSON(http.StatusOK, resp)
}

// HealthRecordPost - Add a new health-record to the system
func (s *Server) HealthRecordPost(ctx echo.Context) error {
	specialtyIdStr := ctx.FormValue("specialty")
	specialtyId, err := uuid.Parse(specialtyIdStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	recordType := ctx.FormValue("type")
	patientIdStr := ctx.FormValue("patientId")
	patientId, err := uuid.Parse(patientIdStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	title := ctx.FormValue("title")
	description := ctx.FormValue("description")

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	doctor, err := s.Controller.GetDoctorByID(claims.UserID)
	if err != nil {
		return err
	}

	file, err := ctx.FormFile("payload")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buf, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		return err
	}

	hr, err := s.Controller.CreateHealthRecord(controller.CreateHealthRecordArgs{
		Type:          recordType,
		SpecialtyId:   specialtyId,
		PatientId:     patientId,
		ContentFormat: kind.Extension,
		Title:         title,
		Description:   description,
		Author:        doctor.Firstname + " " + doctor.Lastname,
	})
	if err != nil {
		return err
	}

	spacialty, err := s.Controller.GetSpecialtyByID(specialtyId)
	if err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(storagePath, file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := dst.Write(buf); err != nil {
		return err
	}

	resp := models.NewHealthRecordResponse(hr, spacialty)

	return ctx.JSON(http.StatusOK, resp)
}
