package server

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
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
	healthReacordIDStr := ctx.Param("healthRecordId")
	healthReacordID, err := uuid.Parse(healthReacordIDStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if err := s.Controller.DeleteHealthRecordByID(healthReacordID); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

// HealthRecordHealthReacordIDGet - Find health-record by ID
func (s *Server) HealthRecordHealthReacordIDGet(ctx echo.Context) error {
	healthReacordIDStr := ctx.Param("healthRecordId")
	healthReacordID, err := uuid.Parse(healthReacordIDStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	hr, err := s.Controller.GetHealthRecordByID(healthReacordID)
	if err != nil {
		return err
	}

	specialty, err := s.Controller.GetSpecialtyByID(hr.SpecialtyID.Bytes)
	if err != nil {
		return err
	}

	patient, err := s.Controller.GetPatientByID(hr.PatientID.Bytes)
	if err != nil {
		return err
	}

	resp, err := models.NewHealthRecordResponse(db.CreateHealthRecordResult{
		HealthRecord: hr,
		Specialty:    specialty,
		Patient:      patient,
	}, s.ivEncryptionKey)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// HealthRecordPost - Add a new health-record to the system
func (s *Server) HealthRecordPost(ctx echo.Context) error {
	specialtyIDStr := ctx.FormValue("specialty")
	specialtyID, err := uuid.Parse(specialtyIDStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	recordType := ctx.FormValue("type")
	patientIDStr := ctx.FormValue("patientId")
	patientID, err := uuid.Parse(patientIDStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	title := ctx.FormValue("title")
	description := ctx.FormValue("description")

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	file, err := ctx.FormFile("payload")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	kind, err := filetype.Match(data)
	if err != nil {
		return err
	}

	res, err := s.Controller.CreateHealthRecord(controller.CreateHealthRecordArgs{
		Type:          recordType,
		SpecialtyID:   specialtyID,
		PatientID:     patientID,
		ContentFormat: kind.Extension,
		Title:         title,
		Description:   description,
		DoctorID:      claims.UserID,
		Payload:       bytes.NewReader(data),
	})
	if err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(storagePath, file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := dst.Write(data); err != nil {
		return err
	}

	resp, err := models.NewHealthRecordResponse(res, s.ivEncryptionKey)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// HealthRecordEvolutionPost - Add a new evolution to the system
func (s *Server) HealthRecordEvolutionPost(ctx echo.Context) error {
	req := models.EvolutionRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	res, err := s.Controller.CreateEvolution(controller.CreateEvolutionArgs{
		EvolutionRequest: req,
		InstitutionID:    claims.InstitutionID,
		DoctorID:         claims.UserID,
	})
	if err != nil {
		return err
	}

	resp, err := models.NewHealthRecordResponse(res, s.ivEncryptionKey)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}
