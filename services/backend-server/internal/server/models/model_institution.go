package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Institution struct {
	Name        string `json:"name,omitempty"`
	GovId       string `json:"gov_id,omitempty"`
	Credentials string `json:"credentials,omitempty"`
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
}

type InstitutionsResponse struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Institution
	Pending bool `json:"pending,omitempty"`
}

func NewInstitutionResponse(institution db.Institution) (InstitutionsResponse, error) {
	institutionUUID, err := uuid.FromBytes(institution.Uuid.Bytes[:])
	if err != nil {
		return InstitutionsResponse{}, err
	}

	return InstitutionsResponse{
		UUID: institutionUUID,
		Institution: Institution{
			Name:        institution.Name,
			GovId:       institution.GovID,
			Credentials: institution.Credentials,
			Type:        institution.Type,
			Address:     institution.Address,
		},
		Pending: institution.Pending,
	}, nil
}

type InstitutionsEnrollmentRequestsResponse struct {
	UUID             uuid.UUID `json:"uuid,omitempty"`
	InstitutionUUID  uuid.UUID `json:"institution_uuid,omitempty"`
	DoctorUUID       uuid.UUID `json:"doctor_uuid,omitempty"`
	Pending          bool      `json:"pending,omitempty"`
	Approved         bool      `json:"approved,omitempty"`
	ProfessionalType string    `json:"professional-type,omitempty"`
}

func NewInstitutionsEnrollmentRequestsResponse(uer db.InstitutionEnrollmentRequest) (InstitutionsEnrollmentRequestsResponse, error) {
	uerUUID, err := uuid.FromBytes(uer.Uuid.Bytes[:])
	if err != nil {
		return InstitutionsEnrollmentRequestsResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(uer.InstitutionUuid.Bytes[:])
	if err != nil {
		return InstitutionsEnrollmentRequestsResponse{}, err
	}

	doctorUUID, err := uuid.FromBytes(uer.DoctorUuid.Bytes[:])
	if err != nil {
		return InstitutionsEnrollmentRequestsResponse{}, err
	}

	return InstitutionsEnrollmentRequestsResponse{
		UUID:             uerUUID,
		InstitutionUUID:  institutionUUID,
		DoctorUUID:       doctorUUID,
		Pending:          uer.Pending,
		Approved:         uer.Approved,
		ProfessionalType: string(uer.ProfessionalType),
	}, nil
}

func NewGovernmentEnrollmentRequestsResponse(ger db.GovernmentEnrollmentRequest) (GovernmentEnrollmentRequestsResponse, error) {
	gerUUID, err := uuid.FromBytes(ger.Uuid.Bytes[:])
	if err != nil {
		return GovernmentEnrollmentRequestsResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(ger.InstitutionUuid.Bytes[:])
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
	UUID            uuid.UUID `json:"uuid,omitempty"`
	InstitutionUUID uuid.UUID `json:"institution_uuid,omitempty"`
	Pending         bool      `json:"pending,omitempty"`
	Approved        bool      `json:"approved,omitempty"`
}

type InstitutionEnrollmentRequest struct {
	InstitutionUUID uuid.UUID `json:"institution_uuid,omitempty"`
	Pending         bool      `json:"pending,omitempty"`
	Approved        bool      `json:"approved,omitempty"`
}

type InstitutionWithUser struct {
	Institution
	InstitutionUser InstitutionUser `json:"institution_user,omitempty"`
}
