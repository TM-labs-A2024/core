package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Nurse struct {
	InstitutionID uuid.UUID `json:"institutionId,omitempty"`
	Firstname     string    `json:"firstname,omitempty"`
	Lastname      string    `json:"lastname,omitempty"`
	GovId         string    `json:"govId,omitempty"`
	Birthdate     string    `json:"birthdate,omitempty"`
	Email         string    `json:"email,omitempty"`
	PhoneNumber   string    `json:"phoneNumber,omitempty"`
	Credentials   string    `json:"credentials,omitempty"`
	Sex           string    `json:"sex,omitempty"`
}

type NursesPutRequest struct {
	ID uuid.UUID `json:"id,omitempty"`
	Nurse
	Password string `json:"password,omitempty"`
	Pending  bool   `json:"pending,omitempty"`
}

type NursePostRequest struct {
	Nurse
	Password string `json:"password,omitempty"`
}

type NursesResponse struct {
	UUID uuid.UUID `json:"id,omitempty"`
	Nurse
	Pending bool `json:"pending,omitempty"`
}

func NewNurseResponse(nurse db.Nurse) (NursesResponse, error) {
	nurseUUID, err := uuid.FromBytes(nurse.ID.Bytes[:])
	if err != nil {
		return NursesResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(nurse.InstitutionID.Bytes[:])
	if err != nil {
		return NursesResponse{}, err
	}

	return NursesResponse{
		UUID: nurseUUID,
		Nurse: Nurse{
			InstitutionID: institutionUUID,
			Firstname:     nurse.Firstname,
			Lastname:      nurse.Lastname,
			GovId:         nurse.GovID,
			Birthdate:     nurse.Birthdate.Time.Format(constants.ISOLayout),
			Email:         nurse.Email,
			PhoneNumber:   nurse.PhoneNumber,
			Credentials:   nurse.Credentials,
			Sex:           nurse.Sex,
		},
		Pending: nurse.Pending,
	}, nil
}