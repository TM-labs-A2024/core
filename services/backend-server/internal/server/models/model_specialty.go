package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Specialty struct {
	ID          uuid.UUID     `json:"id"`
	Description string        `json:"description"`
	Name        SpecialtyName `json:"name"`
}

func NewSpecialtyResponse(specialty db.Specialty) Specialty {
	return Specialty{
		ID:          specialty.ID.Bytes,
		Description: specialty.Description,
		Name:        SpecialtyName(specialty.Name),
	}
}

func NewSpecialtiesResponse(specialties []db.Specialty) []Specialty {
	resp := []Specialty{}
	for _, specialty := range specialties {
		resp = append(resp, NewSpecialtyResponse(specialty))
	}

	return resp
}
