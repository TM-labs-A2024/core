package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/utils"
	"github.com/google/uuid"
)

type HealthRecord struct {
	Content       string `json:"content"`
	Type          string `json:"type"`
	ContentFormat string `json:"content-format"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
}

type EvolutionType string

const (
	EvolutionTypeHospitalizacion EvolutionType = "hospitalización"
	EvolutionTypeAlta            EvolutionType = "alta"
)

type EvolutionRequest struct {
	Specialty   uuid.UUID `json:"specialty"`
	PatientID   uuid.UUID `json:"patientId"`
	Bed         *string   `json:"bed"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Payload     struct {
		Type struct {
			Value string `json:"value"`
			Label string `json:"label"`
		} `json:"type"`
		Reason      string `json:"reason"`
		Description string `json:"description"`
		History     string `json:"history"`
		Examination string `json:"examination"`
		Comments    string `json:"comments"`
	} `json:"payload"`
}

type HealthRecordResponse struct {
	ID uuid.UUID `json:"id"`
	HealthRecord
	Specialty Specialty `json:"specialty"`
}

func NewHealthRecordResponse(res db.CreateHealthRecordResult, iv string) (HealthRecordResponse, error) {
	specialty := NewSpecialtyResponse(res.Specialty)
	url, err := utils.GetAESDecrypted(
		res.HealthRecord.PublicKey,
		res.Patient.PrivateKey,
		iv,
	)
	if err != nil {
		return HealthRecordResponse{}, err
	}
	return HealthRecordResponse{
		ID: res.HealthRecord.ID.Bytes,
		HealthRecord: HealthRecord{
			Content:       url,
			Type:          string(res.HealthRecord.Type),
			ContentFormat: res.HealthRecord.ContentFormat,
			Title:         res.HealthRecord.Title,
			Description:   res.HealthRecord.Description,
			Author:        res.HealthRecord.Author,
		},
		Specialty: specialty,
	}, nil
}
