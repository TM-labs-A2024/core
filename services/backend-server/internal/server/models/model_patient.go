package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Patient struct {
	InstitutionID uuid.UUID `json:"institution_id"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	GovID         string    `json:"govId"`
	Birthdate     string    `json:"birthdate"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	Sex           string    `json:"sex"`
	Status        string    `json:"status"`
	Bed           string    `json:"bed"`
	Pending       bool      `json:"pending"`
}

type PatientPostRequest struct {
	Patient
	Password string `json:"password"`
}

type PatientPutRequest struct {
	ID uuid.UUID `json:"id"`
	PatientPostRequest
}

type PatientResponse struct {
	ID uuid.UUID `json:"id"`
	Patient
}

func NewPatientResponse(patient db.Patient) (PatientResponse, error) {
	return PatientResponse{
		ID: patient.ID.Bytes,
		Patient: Patient{
			Firstname:     patient.Firstname,
			Lastname:      patient.Lastname,
			GovID:         patient.GovID,
			Birthdate:     patient.Birthdate.Time.Format(constants.ISOLayout),
			Email:         patient.Email,
			PhoneNumber:   patient.PhoneNumber,
			Sex:           patient.Sex,
			Pending:       patient.Pending,
			Status:        string(patient.Status),
			Bed:           patient.Bed,
			InstitutionID: patient.InstitutionID.Bytes,
		},
	}, nil
}

type PatientsHealthRecordsResponse struct {
	UUID          uuid.UUID `json:"id"`
	Content       string    `json:"content"`
	Type          string    `json:"type"`
	Specialty     Specialty `json:"specialty"`
	ContentFormat string    `json:"contentFormat"`
}

func NewPatientsHealthRecordsResponse(hr db.HealthRecord, content string, specialty db.Specialty) (PatientsHealthRecordsResponse, error) {
	hrUUID, err := uuid.FromBytes(hr.ID.Bytes[:])
	if err != nil {
		return PatientsHealthRecordsResponse{}, err
	}

	specialtyUUID, err := uuid.FromBytes(specialty.ID.Bytes[:])
	if err != nil {
		return PatientsHealthRecordsResponse{}, err
	}

	return PatientsHealthRecordsResponse{
		UUID:    hrUUID,
		Content: content,
		Type:    string(hr.Type),
		Specialty: Specialty{
			ID:          specialtyUUID,
			Description: specialty.Description,
			Name:        SpecialtyName(specialty.Name),
		},
		ContentFormat: hr.ContentFormat,
	}, nil
}
