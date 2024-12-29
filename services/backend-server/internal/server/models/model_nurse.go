package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Nurse struct {
	InstitutionID uuid.UUID `json:"institutionId"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	GovID         string    `json:"govId"`
	Birthdate     string    `json:"birthdate"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	Credentials   string    `json:"credentials"`
	Sex           string    `json:"sex"`
}

type NursesPutRequest struct {
	ID uuid.UUID `json:"id"`
	Nurse
	Password string `json:"password"`
	Pending  bool   `json:"pending"`
}

type NursePostRequest struct {
	Nurse
	Password string `json:"password"`
}

type NursesResponse struct {
	UUID uuid.UUID `json:"id"`
	Nurse
	Pending bool `json:"pending"`
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
			GovID:         nurse.GovID,
			Birthdate:     nurse.Birthdate.Time.Format(constants.ISOLayout),
			Email:         nurse.Email,
			PhoneNumber:   nurse.PhoneNumber,
			Credentials:   nurse.Credentials,
			Sex:           nurse.Sex,
		},
		Pending: nurse.Pending,
	}, nil
}
