package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Government struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GovernmentEnrollmentRequest struct {
	InstitutionID uuid.UUID `json:"institutionId"`
	ID            uuid.UUID `json:"id"`
	Pending       bool      `json:"pending"`
	Approved      bool      `json:"approved"`
}

func NewGovernmentEnrollmentRequest(er db.GovernmentEnrollmentRequest) GovernmentEnrollmentRequest {
	return GovernmentEnrollmentRequest{
		InstitutionID: er.InstitutionID.Bytes,
		ID:            er.ID.Bytes,
		Pending:       er.Pending,
		Approved:      er.Approved,
	}
}
