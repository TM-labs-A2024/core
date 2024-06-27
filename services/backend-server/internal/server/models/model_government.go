package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Government struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GovernmentEnrollmentRequest struct {
	InstitutionId uuid.UUID `json:"institutionId,omitempty"`
	ID            uuid.UUID `json:"id,omitempty"`
	Pending       bool      `json:"pending,omitempty"`
	Approved      bool      `json:"approved,omitempty"`
}

func NewGovernmentEnrollmentRequest(er db.GovernmentEnrollmentRequest) GovernmentEnrollmentRequest {
	return GovernmentEnrollmentRequest{
		InstitutionId: er.InstitutionID.Bytes,
		ID:            er.ID.Bytes,
		Pending:       er.Pending,
		Approved:      er.Approved,
	}
}
