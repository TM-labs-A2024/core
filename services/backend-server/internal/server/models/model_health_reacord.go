package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type HealthRecord struct {
	Content       string    `json:"content"`
	Type          string    `json:"type"`
	Specialty     Specialty `json:"specialty"`
	ContentFormat string    `json:"content-format"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Author        string    `json:"author"`
}

type HealthRecordResponse struct {
	ID uuid.UUID `json:"id"`
	HealthRecord
}

func NewHealthRecordResponse(hr db.HealthRecord, spec db.Specialty) HealthRecordResponse {
	specialty := NewSpecialtyResponse(spec)

	return HealthRecordResponse{
		ID: hr.ID.Bytes,
		HealthRecord: HealthRecord{
			Content:       "https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/001_Tacos_de_carnitas%2C_carne_asada_y_al_pastor.jpg/1920px-001_Tacos_de_carnitas%2C_carne_asada_y_al_pastor.jpg",
			Type:          string(hr.Type),
			Specialty:     specialty,
			ContentFormat: hr.ContentFormat,
			Title:         hr.Title,
			Description:   hr.Description,
			Author:        hr.Author,
		},
	}
}
