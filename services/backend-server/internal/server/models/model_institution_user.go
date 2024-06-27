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
	Firstname     string              `json:"firstname,omitempty"`
	Lastname      string              `json:"lastname,omitempty"`
	GovId         string              `json:"govId,omitempty"`
	Birthdate     string              `json:"birthdate,omitempty"`
	Email         string              `json:"email,omitempty"`
	PhoneNumber   string              `json:"phoneNumber,omitempty"`
	Role          InstitutionUserRole `json:"role,omitempty"`
	InstitutionID uuid.UUID           `json:"institutionId,omitempty"`
}

type InstitutionUserPostRequest struct {
	InstitutionUser
	Password string `json:"password,omitempty"`
}

type InstitutionUserPutRequest struct {
	ID uuid.UUID `json:"id"`
	InstitutionUser
	Password string `json:"password,omitempty"`
}

type InstitutionUserResponse struct {
	ID uuid.UUID `json:"id,omitempty"`
	InstitutionUser
}

func NewInstitutionUserResponse(user db.InstitutionUser) (InstitutionUserResponse, error) {
	resp := InstitutionUserResponse{
		ID: user.ID.Bytes,
		InstitutionUser: InstitutionUser{
			Firstname:     user.Firstname,
			Lastname:      user.Lastname,
			GovId:         user.GovID,
			Birthdate:     user.Birthdate.Time.Format(constants.ISOLayout),
			Email:         user.Email,
			PhoneNumber:   user.PhoneNumber,
			Role:          InstitutionUserRole(user.Role),
			InstitutionID: user.InstitutionID.Bytes,
		},
	}

	return resp, nil
}
