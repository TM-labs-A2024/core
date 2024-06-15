package models

type Specialty struct {
	ID          int           `json:"id,omitempty"`
	Description string        `json:"description,omitempty"`
	Name        SpecialtyName `json:"name,omitempty"`
}
