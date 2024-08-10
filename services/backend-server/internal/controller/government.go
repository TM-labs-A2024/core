package controller

import (
	"context"
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (c Controller) CreateGovernmentEnrollmentRequest(institutionID, governmentID uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	er, err := c.queries.CreateGovernmentEnrollmentRequests(
		context.Background(),
		db.CreateGovernmentEnrollmentRequestsParams{
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: institutionID,
			},
			GovernmentID: pgtype.UUID{
				Valid: true,
				Bytes: governmentID,
			},
		},
	)	
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) ApproveGovernmentEnrollmentRequest(enrollmentRequestID uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.GovernmentEnrollmentRequest{}, err
	}
	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	er, err := txQuery.GetGovernmentEnrollmentRequestByID(context.Background(), pgtype.UUID{
		Valid: true,
		Bytes: enrollmentRequestID,
	})
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	er, err = txQuery.UpdatePendingGovernmentEnrollmentRequestsByID(
		context.Background(),
		db.UpdatePendingGovernmentEnrollmentRequestsByIDParams{
			InstitutionID: er.InstitutionID,
			Pending:       false,
			Approved:      true,
			ID:            er.ID,
		},
	)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	institution, err := txQuery.GetInstitutionByID(context.Background(), er.InstitutionID)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	count, err := txQuery.CountPendingGovernmentEnrollmentRequestsByInstitutionID(context.Background(), institution.ID)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	institution.Pending = count > 0
	if _, err := txQuery.UpdateInstitutionByID(context.Background(), db.UpdateInstitutionByIDParams{
		Address:      institution.Address,
		Name:         institution.Name,
		Credentials:  institution.Credentials,
		Type:         institution.Type,
		GovID:        institution.GovID,
		Pending:      institution.Pending,
		GovernmentID: institution.GovernmentID,
		ID:           institution.ID,
	}); err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	tx.Commit(context.Background())

	return er, nil
}

func (c Controller) DenyGovernmentEnrollmentRequest(enrollmentRequestID uuid.UUID) (db.GovernmentEnrollmentRequest, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.GovernmentEnrollmentRequest{}, err
	}
	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	er, err := txQuery.GetGovernmentEnrollmentRequestByID(context.Background(), pgtype.UUID{
		Valid: true,
		Bytes: enrollmentRequestID,
	})
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	er, err = txQuery.UpdatePendingGovernmentEnrollmentRequestsByID(
		context.Background(),
		db.UpdatePendingGovernmentEnrollmentRequestsByIDParams{
			InstitutionID: er.InstitutionID,
			Pending:       false,
			Approved:      false,
			ID:            er.ID,
		},
	)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	institution, err := txQuery.GetInstitutionByID(context.Background(), er.InstitutionID)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	count, err := txQuery.CountPendingGovernmentEnrollmentRequestsByInstitutionID(context.Background(), institution.ID)
	if err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	institution.Pending = count > 0
	if _, err := txQuery.UpdateInstitutionByID(context.Background(), db.UpdateInstitutionByIDParams{
		Address:      institution.Address,
		Name:         institution.Name,
		Credentials:  institution.Credentials,
		Type:         institution.Type,
		GovID:        institution.GovID,
		Pending:      institution.Pending,
		GovernmentID: institution.GovernmentID,
		ID:           institution.ID,
	}); err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.GovernmentEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) ListGovernmentEnrollmentRequest() ([]db.GovernmentEnrollmentRequest, error) {
	return c.queries.ListGovernmentEnrollmentRequests(context.Background())
}

func (c Controller) DeleteGovernmentEnrollmentRequestByInsitutionID(institutionId uuid.UUID) error {
	return c.queries.DeleteGovernmentEnrollmentRequestByInsitutionID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: institutionId,
		},
	)
}
