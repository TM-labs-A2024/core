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
	Firstname     string              `json:"firstname"`
	Lastname      string              `json:"lastname"`
	GovId         string              `json:"govId"`
	Birthdate     string              `json:"birthdate"`
	Email         string              `json:"email"`
	PhoneNumber   string              `json:"phoneNumber"`
	Role          InstitutionUserRole `json:"role"`
	InstitutionID uuid.UUID           `json:"institutionId"`
}

type InstitutionUserPostRequest struct {
	InstitutionUser
	Password string `json:"password"`
}

type InstitutionUserPutRequest struct {
	ID uuid.UUID `json:"id"`
	InstitutionUser
	Password string `json:"password"`
}

type InstitutionUserResponse struct {
	ID uuid.UUID `json:"id"`
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
