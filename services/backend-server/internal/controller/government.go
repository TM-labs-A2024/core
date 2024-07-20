package controller

import (
	"context"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) GetGovernmentByLogin(email, crypt string) (db.Government, error) {
	goverment, err := c.queries.GetGovernmentByLogin(
		context.Background(),
		db.GetGovernmentByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.Government{}, err
	}

	return goverment, nil
}

func (c Controller) CreateGovernmentEnrollmentRequest(institutionId uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	er, err := c.queries.CreateGovernmentEnrollmentRequests(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: institutionId,
		},
	)	
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) ApproveGovernmentEnrollmentRequest(enrollmentRequestId uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	er, err := c.queries.GetGovernmentEnrollmentRequestByID(context.Background(), pgtype.UUID{
		Valid: true,
		Bytes: enrollmentRequestId,
	})
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	er, err = c.queries.UpdateGovernmentEnrollmentRequestsByID(
		context.Background(),
		db.UpdateGovernmentEnrollmentRequestsByIDParams{
			InstitutionID: er.InstitutionID,
			Pending:       false,
			Approved:      true,
			ID:            er.ID,
		},
	)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) DenyGovernmentEnrollmentRequest(enrollmentRequestId uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	er, err := c.queries.GetGovernmentEnrollmentRequestByID(context.Background(), pgtype.UUID{
		Valid: true,
		Bytes: enrollmentRequestId,
	})
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	er, err = c.queries.UpdateGovernmentEnrollmentRequestsByID(
		context.Background(),
		db.UpdateGovernmentEnrollmentRequestsByIDParams{
			InstitutionID: er.InstitutionID,
			Pending:       false,
			Approved:      false,
			ID:            er.ID,
		},
	)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) ListGovernmentEnrollmentRequest() ([]db.GovernmentEnrollmentRequest, error) {
	return c.queries.ListGovernmentEnrollmentRequests(context.Background())
}
