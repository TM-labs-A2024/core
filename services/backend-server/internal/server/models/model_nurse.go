package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Nurse struct {
	InstitutionUUID uuid.UUID `json:"institution_uuid,omitempty"`
	Firstname       string    `json:"firstname,omitempty"`
	Lastname        string    `json:"lastname,omitempty"`
	GovId           string    `json:"gov_id,omitempty"`
	Birthdate       string    `json:"birthdate,omitempty"`
	Email           string    `json:"email,omitempty"`
	PhoneNumber     string    `json:"phone_number,omitempty"`
	Credentials     string    `json:"credentials,omitempty"`
}

type NursesPutRequest struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Nurse
	Password string `json:"password,omitempty"`
}

type NursesPostRequest struct {
	Nurse
	Password string `json:"password,omitempty"`
}

type NursesResponse struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Nurse
	Pending bool `json:"pending,omitempty"`
}

func NewNurseResponse(nurse db.Nurse) (NursesResponse, error) {
	nurseUUID, err := uuid.FromBytes(nurse.Uuid.Bytes[:])
	if err != nil {
		return NursesResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(nurse.InstitutionUuid.Bytes[:])
	if err != nil {
		return NursesResponse{}, err
	}

	return NursesResponse{
		UUID: nurseUUID,
		Nurse: Nurse{
			InstitutionUUID: institutionUUID,
			Firstname:       nurse.Firstname,
			Lastname:        nurse.Lastname,
			GovId:           nurse.GovID,
			Birthdate:       nurse.Birthdate.Time.Format(constants.ISOLayout),
			Email:           nurse.Email,
			PhoneNumber:     nurse.PhoneNumber,
			Credentials:     nurse.Credentials,
		},
		Pending: nurse.Pending,
	}, nil
}
