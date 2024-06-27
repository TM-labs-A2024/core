package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Institution struct {
	Name        string `json:"name,omitempty"`
	GovId       string `json:"govId,omitempty"`
	Credentials string `json:"credentials,omitempty"`
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
}

type CreateInstitutionRequest struct {
	Name            string          `json:"name,omitempty"`
	GovId           string          `json:"govId,omitempty"`
	Credentials     string          `json:"credentials,omitempty"`
	Type            string          `json:"type,omitempty"`
	Address         string          `json:"address,omitempty"`
	InstitutionUser InstitutionUser `json:"institutionUser,omitempty"`
}

type InstitutionWithID struct {
	ID uuid.UUID `json:"id,omitempty"`
	Institution
	Pending bool `json:"pending,omitempty"`
}

type InstitutionWithUserRequest struct {
	InstitutionWithID
	InstitutionUser InstitutionUserPostRequest `json:"institutionUser,omitempty"`
}

type InstitutionWithUserResponse struct {
	InstitutionWithID
	InstitutionUser InstitutionUserResponse `json:"institutionUser,omitempty"`
}

func NewInstitutionResponse(institution db.Institution, user db.InstitutionUser) (InstitutionWithUserResponse, error) {
	resp := InstitutionWithUserResponse{
		InstitutionWithID: InstitutionWithID{
			ID: institution.ID.Bytes,
			Institution: Institution{
				Name:        institution.Name,
				GovId:       institution.GovID,
				Credentials: institution.Credentials,
				Type:        string(institution.Type),
				Address:     institution.Address,
			},
			Pending: institution.Pending,
		},
		InstitutionUser: InstitutionUserResponse{
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
		},
	}

	return resp, nil
}

type InstitutionsEnrollmentRequestsResponse struct {
	UUID          uuid.UUID  `json:"id,omitempty"`
	InstitutionID uuid.UUID  `json:"institutionId,omitempty"`
	DoctorID      *uuid.UUID `json:"doctorId,omitempty"`
	NurseID       *uuid.UUID `json:"nurseId,omitempty"`
	Pending       bool       `json:"pending,omitempty"`
	Approved      bool       `json:"approved,omitempty"`
}

func NewInstitutionsEnrollmentRequestsResponse(ier db.InstitutionEnrollmentRequest) (InstitutionsEnrollmentRequestsResponse, error) {
	userUUID, err := uuid.FromBytes(ier.ID.Bytes[:])
	if err != nil {
		return InstitutionsEnrollmentRequestsResponse{}, err
	}

	institutionID, err := uuid.FromBytes(ier.InstitutionID.Bytes[:])
	if err != nil {
		return InstitutionsEnrollmentRequestsResponse{}, err
	}

	doctorID, _ := uuid.FromBytes(ier.DoctorID.Bytes[:])
	nurseID, _ := uuid.FromBytes(ier.NurseID.Bytes[:])

	resp := InstitutionsEnrollmentRequestsResponse{
		UUID:          userUUID,
		Pending:       ier.Pending,
		Approved:      ier.Approved,
		InstitutionID: institutionID,
	}

	if ier.DoctorID.Valid {
		resp.DoctorID = &doctorID
	} else {
		resp.NurseID = &nurseID
	}

	return resp, nil
}

func NewGovernmentEnrollmentRequestsResponse(ger db.GovernmentEnrollmentRequest) (GovernmentEnrollmentRequestsResponse, error) {
	gerUUID, err := uuid.FromBytes(ger.ID.Bytes[:])
	if err != nil {
		return GovernmentEnrollmentRequestsResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(ger.InstitutionID.Bytes[:])
	if err != nil {
		return GovernmentEnrollmentRequestsResponse{}, err
	}

	return GovernmentEnrollmentRequestsResponse{
		UUID:            gerUUID,
		InstitutionUUID: institutionUUID,
		Pending:         ger.Pending,
		Approved:        ger.Approved,
	}, nil
}

type GovernmentEnrollmentRequestsResponse struct {
	UUID            uuid.UUID `json:"id,omitempty"`
	InstitutionUUID uuid.UUID `json:"institutionId,omitempty"`
	Pending         bool      `json:"pending,omitempty"`
	Approved        bool      `json:"approved,omitempty"`
}

type InstitutionEnrollmentRequest struct {
	InstitutionUUID uuid.UUID `json:"institutionId,omitempty"`
	Pending         bool      `json:"pending,omitempty"`
	Approved        bool      `json:"approved,omitempty"`
}
