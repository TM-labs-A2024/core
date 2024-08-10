package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) DeleteDoctorByID(doctorID uuid.UUID) error {
	return c.queries.DeleteDoctorByID(context.Background(), pgtype.UUID{
		Bytes: doctorID,
		Valid: true,
	})
}

func (c Controller) GetDoctorByID(doctorID uuid.UUID) (db.Doctor, error) {
	return c.queries.GetDoctorByID(context.Background(), pgtype.UUID{
		Bytes: doctorID,
		Valid: true,
	})
}

func (c Controller) GetDoctorByLogin(email, crypt string) (db.Doctor, error) {
	institutionUser, err := c.queries.GetDoctorByLogin(
		context.Background(),
		db.GetDoctorByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.Doctor{}, err
	}

	return institutionUser, nil
}

func (c Controller) ListDoctors() ([]db.Doctor, error) {
	institutionUser, err := c.queries.ListDoctors(context.Background())
	if err != nil {
		return nil, err
	}

	return institutionUser, nil
}

func (c Controller) ListDoctorsByInstitutionID(id uuid.UUID) ([]db.Doctor, error) {
	institutionUser, err := c.queries.ListDoctorsByInstitutionID(
		context.Background(),
		pgtype.UUID{
			Bytes: id,
			Valid: true,
		})
	if err != nil {
		return nil, err
	}

	return institutionUser, nil
}

func (c Controller) ListDoctorsBySpecialtyID(id uuid.UUID) ([]db.DoctorSpecialty, error) {
	specialtyDoctorJunctions, err := c.queries.ListSpecialtyDoctorJunctionsBySpecialtyID(
		context.Background(),
		pgtype.UUID{
			Bytes: id,
			Valid: true,
		})
	if err != nil {
		return nil, err
	}

	return specialtyDoctorJunctions, nil
}

func (c Controller) ListAccessRequestsByDoctorID(doctorID uuid.UUID) ([]db.DoctorAccessRequest, error) {
	accessRequests, err := c.queries.ListAccessRequestsByDoctorID(
		context.Background(),
		pgtype.UUID{
			Bytes: doctorID,
			Valid: true,
		},
	)
	if err != nil {
		return nil, err
	}

	return accessRequests, nil
}

func (c Controller) CreateDoctor(req models.DoctorsPostRequest) (db.Doctor, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.Doctor{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Doctor{}, err
	}

	doctor, err := txQuery.CreateDoctor(
		context.Background(),
		db.CreateDoctorParams{
			InstitutionID: pgtype.UUID{
				Bytes: req.InstitutionID,
				Valid: true,
			},
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			GovID:     req.GovID,
			Birthdate: pgtype.Timestamp{
				Time:  birthdate,
				Valid: true,
			},
			Crypt:       req.Password,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Credentials: req.Credentials,
			Sex:         req.Sex,
		})
	if err != nil {
		return db.Doctor{}, err
	}

	_, err = txQuery.CreateInstitutionEnrollmentRequest(
		context.Background(),
		db.CreateInstitutionEnrollmentRequestParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctor.ID.Bytes,
			},
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: doctor.InstitutionID.Bytes,
			},
		},
	)
	if err != nil {
		return db.Doctor{}, err
	}

	for _, specialty := range req.Specialties {
		if _, err := txQuery.CreateSpecialtyDoctorJunction(context.Background(),
			db.CreateSpecialtyDoctorJunctionParams{
				DoctorID: doctor.ID,
				SpecialtyID: pgtype.UUID{
					Valid: true,
					Bytes: specialty,
				},
			},
		); err != nil {
			return db.Doctor{}, err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.Doctor{}, err
	}

	return doctor, nil
}

func (c Controller) UpdateDoctorByID(req models.DoctorsPutRequest) (db.Doctor, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.Doctor{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		c.logger.Debug("error messagw2:", slog.String("err", err.Error()))
		return db.Doctor{}, err
	}

	doctor, err := txQuery.UpdateDoctorByID(
		context.Background(),
		db.UpdateDoctorByIDParams{
			ID: pgtype.UUID{
				Bytes: req.ID,
				Valid: true,
			},
			InstitutionID: pgtype.UUID{
				Bytes: req.InstitutionID,
				Valid: true,
			},
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			GovID:     req.GovID,
			Birthdate: pgtype.Timestamp{
				Time:  birthdate,
				Valid: true,
			},
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Credentials: req.Credentials,
			Sex:         req.Sex,
		})
	if err != nil {
		return db.Doctor{}, err
	}

	specialties, err := txQuery.ListSpecialtyDoctorJunctionsByDoctorID(
		context.Background(),
		pgtype.UUID{
			Bytes: req.ID,
			Valid: true,
		},
	)
	if err != nil {
		return db.Doctor{}, err
	}

	specialtiesToAdd := map[uuid.UUID]bool{}
	for _, specialty := range req.Specialties {
		specialtiesToAdd[specialty] = true
	}

	toDelete := []db.DoctorSpecialty{}
	for _, specialty := range specialties {
		if _, ok := specialtiesToAdd[specialty.SpecialtyID.Bytes]; !ok {
			toDelete = append(toDelete, specialty)
		} else {
			specialtiesToAdd[specialty.SpecialtyID.Bytes] = false
		}
	}

	for _, specialty := range toDelete {
		if err := txQuery.DeleteSpecialtyDoctorJunction(
			context.Background(),
			db.DeleteSpecialtyDoctorJunctionParams(specialty),
		); err != nil {
			return db.Doctor{}, err
		}
	}

	for specialty, toAdd := range specialtiesToAdd {
		if toAdd {
			txQuery.CreateSpecialtyDoctorJunction(
				context.Background(),
				db.CreateSpecialtyDoctorJunctionParams{
					DoctorID: doctor.ID,
					SpecialtyID: pgtype.UUID{
						Valid: true,
						Bytes: specialty,
					},
				},
			)
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.Doctor{}, err
	}

	return doctor, nil
}

func (c Controller) ListPatientsTreatedByDoctorID(id uuid.UUID) ([]db.Patient, error) {
	patients, err := c.queries.ListPatientsTreatedByDoctorID(
		context.Background(),
		pgtype.UUID{Bytes: id, Valid: true},
	)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (c Controller) ListPatientsTreatedByDoctorIDWithHealthRecordOfSpecialtyID(doctorID, specialtyID uuid.UUID) ([]db.Patient, error) {
	patients, err := c.queries.ListPatientsTreatedByDoctorIDWithHealthRecordOfSpecialtyID(
		context.Background(),
		db.ListPatientsTreatedByDoctorIDWithHealthRecordOfSpecialtyIDParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctorID,
			},
			SpecialtyID: pgtype.UUID{
				Valid: true,
				Bytes: specialtyID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
