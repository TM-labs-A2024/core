package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type InstitutionUserRole string

// List of InstitutionUserRole
const (
	VIEWER InstitutionUserRole = "viewer"
	ADMIN  InstitutionUserRole = "admin"
)

type InstitutionUser struct {
	Firstname       string              `json:"firstname,omitempty"`
	Lastname        string              `json:"lastname,omitempty"`
	GovId           string              `json:"gov_id,omitempty"`
	Birthdate       string              `json:"birthdate,omitempty"`
	Email           string              `json:"email,omitempty"`
	Password        string              `json:"password,omitempty"`
	PhoneNumber     string              `json:"phone_number,omitempty"`
	Role            InstitutionUserRole `json:"role,omitempty"`
	InstitutionUUID uuid.UUID           `json:"institution_uuid,omitempty"`
}

type InstitutionsUsersResponse struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	InstitutionUser
}

func NewInstitutionUserResponse(user db.InstitutionUser) (InstitutionsUsersResponse, error) {
	userUUID, err := uuid.FromBytes(user.Uuid.Bytes[:])
	if err != nil {
		return InstitutionsUsersResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(user.InstitutionUuid.Bytes[:])
	if err != nil {
		return InstitutionsUsersResponse{}, err
	}

	return InstitutionsUsersResponse{
		UUID: userUUID,
		InstitutionUser: InstitutionUser{
			Firstname:       user.Firstname,
			Lastname:        user.Lastname,
			GovId:           user.GovID,
			Birthdate:       user.Birthdate.Time.Format(constants.ISOLayout),
			Email:           user.Email,
			Password:        user.Password,
			PhoneNumber:     user.PhoneNumber,
			Role:            InstitutionUserRole(user.Role),
			InstitutionUUID: institutionUUID,
		},
	}, nil
}

type InstitutionsInstitutionUUIDUsersPostRequest struct {
	InstitutionUser
	Password string `json:"password,omitempty"`
}

type InstitutionsInstitutionUUIDUsersPutRequest struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	InstitutionsInstitutionUUIDUsersPostRequest
}
