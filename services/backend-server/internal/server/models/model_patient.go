package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Patient struct {
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	GovId       string `json:"gov_id,omitempty"`
	Birthdate   string `json:"birthdate,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Pending     bool   `json:"pending,omitempty"`
}

type PatientsPostRequest struct {
	Patient
	Password string `json:"password,omitempty"`
}

type PatientsPutRequest struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	PatientsPostRequest
}

type PatientsResponse struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Patient
}

func NewPatientsResponse(patient db.Patient) (PatientsResponse, error) {
	patientUUID, err := uuid.FromBytes(patient.Uuid.Bytes[:])
	if err != nil {
		return PatientsResponse{}, err
	}

	return PatientsResponse{
		UUID: patientUUID,
		Patient: Patient{
			Firstname:   patient.Firstname,
			Lastname:    patient.Lastname,
			GovId:       patient.GovID,
			Birthdate:   patient.Birthdate.Time.Format(constants.ISOLayout),
			Email:       patient.Email,
			PhoneNumber: patient.PhoneNumber,
			Sex:         patient.Sex,
			Pending:     patient.Pending,
		},
	}, nil
}

type PatientsHealthRecordsResponse struct {
	UUID          uuid.UUID `json:"uuid,omitempty"`
	Content       string    `json:"content,omitempty"`
	Type          string    `json:"type,omitempty"`
	Specialty     Specialty `json:"specialty,omitempty"`
	ContentFormat string    `json:"content-format,omitempty"`
}

func NewPatientsHealthRecordsResponse(hr db.HealthRecord, content string, specialty db.Specialty) (PatientsHealthRecordsResponse, error) {
	hrUUID, err := uuid.FromBytes(hr.Uuid.Bytes[:])
	if err != nil {
		return PatientsHealthRecordsResponse{}, err
	}

	return PatientsHealthRecordsResponse{
		UUID:    hrUUID,
		Content: content,
		Type:    hr.Type,
		Specialty: Specialty{
			ID:          int(specialty.ID),
			Description: specialty.Description,
			Name:        SpecialtyName(specialty.Name),
		},
		ContentFormat: hr.ContentFormat,
	}, nil
}
