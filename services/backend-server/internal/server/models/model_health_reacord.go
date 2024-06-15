package models

type HealthReacord struct {
	Content       string    `json:"content,omitempty"`
	Type          string    `json:"type,omitempty"`
	Specialty     Specialty `json:"specialty,omitempty"`
	ContentFormat string    `json:"content-format,omitempty"`
}
