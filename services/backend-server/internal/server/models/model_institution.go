package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Institution struct {
	Name        string `json:"name"`
	GovId       string `json:"govId"`
	Credentials string `json:"credentials"`
	Type        string `json:"type"`
	Address     string `json:"address"`
}

type CreateInstitutionRequest struct {
	Name            string          `json:"name"`
	GovId           string          `json:"govId"`
	Credentials     string          `json:"credentials"`
	Type            string          `json:"type"`
	Address         string          `json:"address"`
	InstitutionUser InstitutionUser `json:"institutionUser"`
}

type InstitutionWithID struct {
	ID uuid.UUID `json:"id"`
	Institution
	Pending bool `json:"pending"`
}

type InstitutionWithUserRequest struct {
	InstitutionWithID
	InstitutionUser InstitutionUserPostRequest `json:"institutionUser"`
}

type InstitutionWithUserResponse struct {
	InstitutionWithID
	InstitutionUser InstitutionUserResponse `json:"institutionUser"`
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
	UUID          uuid.UUID  `json:"id"`
	InstitutionID uuid.UUID  `json:"institutionId"`
	DoctorID      *uuid.UUID `json:"doctorId"`
	NurseID       *uuid.UUID `json:"nurseId"`
	Pending       bool       `json:"pending"`
	Approved      bool       `json:"approved"`
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
	UUID            uuid.UUID `json:"id"`
	InstitutionUUID uuid.UUID `json:"institutionId"`
	Pending         bool      `json:"pending"`
	Approved        bool      `json:"approved"`
}

type InstitutionEnrollmentRequest struct {
	InstitutionUUID uuid.UUID `json:"institutionId"`
	Pending         bool      `json:"pending"`
	Approved        bool      `json:"approved"`
}
