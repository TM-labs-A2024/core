package models

import "github.com/google/uuid"

type Specialty struct {
	ID          uuid.UUID     `json:"id,omitempty"`
	Description string        `json:"description,omitempty"`
	Name        SpecialtyName `json:"name,omitempty"`
}
