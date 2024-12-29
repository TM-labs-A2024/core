package controller

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) GetFirstInstitutionUser(institutionID uuid.UUID) (db.InstitutionUser, error) {
	return c.queries.GetFirstInstitutionUserByInstitutionID(context.Background(), pgtype.UUID{
		Bytes: institutionID,
		Valid: true,
	})
}

func (c Controller) CreateInstitutionEnrollmentRequestDoctor(docID, instID uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	er, err := c.queries.CreateInstitutionEnrollmentRequest(
		context.Background(),
		db.CreateInstitutionEnrollmentRequestParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: docID,
			},
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: instID,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) CreateInstitutionEnrollmentRequestNurse(nurseID, instID uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	institutionId := pgtype.UUID{
		Valid: true,
		Bytes: instID,
	}
	er, err := c.queries.CreateInstitutionEnrollmentRequest(
		context.Background(),
		db.CreateInstitutionEnrollmentRequestParams{
			InstitutionID: institutionId,
			NurseID: pgtype.UUID{
				Valid: true,
				Bytes: nurseID,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	institution, err := c.queries.GetInstitutionByID(context.Background(), institutionId)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	institution.Pending = true
	if _, err := c.queries.UpdateInstitutionByID(context.Background(),
		db.UpdateInstitutionByIDParams{
			Name:         institution.Name,
			Address:      institution.Address,
			Credentials:  institution.Credentials,
			Type:         institution.Type,
			GovID:        institution.GovID,
			Pending:      institution.Pending,
			GovernmentID: institution.GovernmentID,
			ID:           institution.ID,
		}); err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) GetInstitutionEnrollmentRequestsByID(erID uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	er, err := c.queries.GetInstitutionEnrollmentRequestsByID(
		context.Background(),
		pgtype.UUID{
			Bytes: erID,
			Valid: true,
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) DeleteInstitutionEnrollmentRequestsByProfID(profID uuid.UUID) error {
	return c.queries.DeleteInstitutionEnrollmentRequestByProfID(
		context.Background(),
		pgtype.UUID{
			Bytes: profID,
			Valid: true,
		},
	)
}

func (c Controller) ApproveInstitutionEnrollmentRequestsByID(erByID db.InstitutionEnrollmentRequest) (db.InstitutionEnrollmentRequest, error) {
	return c.updateInstitutionEnrollmentRequestByID(erByID, true)
}

func (c Controller) DenyInstitutionEnrollmentRequestsByID(erByID db.InstitutionEnrollmentRequest) (db.InstitutionEnrollmentRequest, error) {
	return c.updateInstitutionEnrollmentRequestByID(erByID, false)
}

func (c Controller) updateInstitutionEnrollmentRequestByID(erByID db.InstitutionEnrollmentRequest, approve bool) (db.InstitutionEnrollmentRequest, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.InstitutionEnrollmentRequest{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	er, err := txQuery.UpdateInstitutionEnrollmentRequestByID(
		context.Background(),
		db.UpdateInstitutionEnrollmentRequestByIDParams{
			InstitutionID: erByID.InstitutionID,
			DoctorID:      erByID.DoctorID,
			NurseID:       erByID.NurseID,
			Pending:       false,
			Approved:      approve,
			ID:            erByID.ID,
		})
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	if erByID.DoctorID.Valid {
		doctor, err := txQuery.GetDoctorByID(context.Background(), erByID.DoctorID)
		if err != nil {
			return db.InstitutionEnrollmentRequest{}, err
		}

		count, err := txQuery.CountPendingInstitutionEnrollmentRequestByDoctorID(context.Background(), er.DoctorID)
		if err != nil {
			return db.InstitutionEnrollmentRequest{}, err
		}

		doctor.Pending = count > 0
		if _, err := txQuery.UpdateDoctorByID(context.Background(), db.UpdateDoctorByIDParams{
			InstitutionID:  doctor.InstitutionID,
			Firstname:      doctor.Firstname,
			Lastname:       doctor.Lastname,
			GovID:          doctor.GovID,
			Birthdate:      doctor.Birthdate,
			Email:          doctor.Email,
			PhoneNumber:    doctor.PhoneNumber,
			Credentials:    doctor.Credentials,
			Pending:        doctor.Pending,
			PatientPending: doctor.PatientPending,
			Sex:            doctor.Sex,
			ID:             doctor.ID,
		}); err != nil {
			return db.InstitutionEnrollmentRequest{}, err
		}

	}

	if erByID.NurseID.Valid {
		nurse, err := txQuery.GetNurseByID(context.Background(), erByID.NurseID)
		if err != nil {
			return db.InstitutionEnrollmentRequest{}, err
		}

		count, err := txQuery.CountPendingInstitutionEnrollmentRequestByNurseID(context.Background(), er.NurseID)
		if err != nil {
			return db.InstitutionEnrollmentRequest{}, err

		}

		nurse.Pending = count > 0
		log.Println(er.Pending)
		if _, err := txQuery.UpdateNurseByID(context.Background(), db.UpdateNurseByIDParams{
			InstitutionID: nurse.InstitutionID,
			Firstname:     nurse.Firstname,
			Lastname:      nurse.Lastname,
			GovID:         nurse.GovID,
			Birthdate:     nurse.Birthdate,
			Email:         nurse.Email,
			PhoneNumber:   nurse.PhoneNumber,
			Credentials:   nurse.Credentials,
			Pending:       nurse.Pending,
			Sex:           nurse.Sex,
			ID:            nurse.ID,
		}); err != nil {
			return db.InstitutionEnrollmentRequest{}, err
		}

	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) ListInstitutionsEnrollmentRequestsByInstitutionID(institutionID uuid.UUID) ([]db.InstitutionEnrollmentRequest, error) {
	ier, err := c.queries.ListInstitutionEnrollmentRequestsByInstitutionID(
		context.Background(),
		pgtype.UUID{
			Bytes: institutionID,
			Valid: true,
		},
	)
	if err != nil {
		return nil, err
	}

	return ier, nil
}

func (c Controller) ListInstitutions(onlyApproved bool) ([]db.Institution, error) {
	var (
		err          error
		institutions []db.Institution
	)

	if onlyApproved {
		institutions, err = c.queries.ListApprovedInstitutions(context.Background())
		if err != nil {
			return nil, err
		}
	} else {
		institutions, err = c.queries.ListInstitutions(context.Background())
		if err != nil {
			return nil, err
		}
	}
	return institutions, nil
}

func (c Controller) GetInstitutionByID(id uuid.UUID) (db.Institution, error) {
	dbID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	institutions, err := c.queries.GetInstitutionByID(context.Background(), dbID)
	if err != nil {
		return db.Institution{}, err
	}

	count, err := c.queries.CountPendingGovernmentEnrollmentRequestsByInstitutionID(context.Background(), dbID)
	if err != nil {
		return db.Institution{}, err
	}

	institutions.Pending = count > 0

	return institutions, nil
}

func (c Controller) DeleteInstitutionByID(id uuid.UUID) error {
	dbID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	if err := c.queries.DeleteInstitutionByID(context.Background(), dbID); err != nil {
		return err
	}

	return nil
}

func (c Controller) CreateInstitution(institution models.Institution, user models.InstitutionUserPostRequest) (db.Institution, db.InstitutionUser, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.Institution{}, db.InstitutionUser{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	gov, err := txQuery.GetGovernment(context.Background())
	if err != nil {
		return db.Institution{}, db.InstitutionUser{}, err
	}

	inst, err := txQuery.CreateInstitution(context.Background(), db.CreateInstitutionParams{
		GovernmentID: gov.ID,
		Name:         institution.Name,
		Address:      institution.Address,
		Credentials:  institution.Credentials,
		Type:         db.InstitutionType(institution.Type),
		GovID:        institution.GovID,
	})
	if err != nil {
		return db.Institution{}, db.InstitutionUser{}, err
	}

	user.InstitutionID = inst.ID.Bytes
	u, err := c.createInstitutionUser(txQuery, user)
	if err != nil {
		return db.Institution{}, db.InstitutionUser{}, err
	}

	_, err = txQuery.CreateGovernmentEnrollmentRequests(context.Background(),
		db.CreateGovernmentEnrollmentRequestsParams{
			InstitutionID: inst.ID,
			GovernmentID:  gov.ID,
		})
	if err != nil {
		return db.Institution{}, db.InstitutionUser{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.Institution{}, db.InstitutionUser{}, err
	}

	return inst, u, nil
}

func (c Controller) UpdateInstitution(institution models.Institution, id uuid.UUID) (db.Institution, error) {
	switch db.InstitutionType(institution.Type) {
	case db.InstitutionTypeClnica, db.InstitutionTypeHospital:
		break
	default:
		return db.Institution{}, fmt.Errorf("invalid intitution type: %s", db.InstitutionType(institution.Type))
	}

	return c.queries.UpdateInstitutionByID(context.Background(), db.UpdateInstitutionByIDParams{
		Name:        institution.Name,
		Address:     institution.Address,
		Credentials: institution.Credentials,
		Type:        db.InstitutionType(institution.Type),
		GovID:       institution.GovID,
		Pending:     false,
		ID: pgtype.UUID{
			Bytes: id,
			Valid: true,
		},
	})
}

func (c Controller) GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID(doctorID, institutionID uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	enrollmentRequest, err := c.queries.GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID(
		context.Background(),
		db.GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionIDParams{
			DoctorID: pgtype.UUID{
				Bytes: doctorID,
				Valid: true,
			},
			InstitutionID: pgtype.UUID{
				Bytes: institutionID,
				Valid: true,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return enrollmentRequest, nil
}
