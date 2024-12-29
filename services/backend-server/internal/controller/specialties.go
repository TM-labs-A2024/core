package controller

import (
	"context"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) GetSpecialtyByID(id uuid.UUID) (db.Specialty, error) {
	specialty, err := c.queries.GetSpecialtyByID(
		context.Background(),
		pgtype.UUID{
			Bytes: id,
			Valid: true,
		},
	)
	if err != nil {
		return db.Specialty{}, err
	}

	return specialty, nil
}

func (c Controller) ListSpecialtiesByDoctorID(id uuid.UUID) ([]db.Specialty, error) {
	specialtiesByDoctor, err := c.queries.ListSpecialtyDoctorJunctionsByDoctorID(
		context.Background(),
		pgtype.UUID{
			Bytes: id,
			Valid: true,
		},
	)
	if err != nil {
		return nil, err
	}

	specialties := []db.Specialty{}
	for _, specialityByDoctor := range specialtiesByDoctor {
		specialty, err := c.queries.GetSpecialtyByID(
			context.Background(),
			specialityByDoctor.SpecialtyID,
		)
		if err != nil {
			return nil, err
		}

		specialties = append(specialties, specialty)
	}

	return specialties, nil
}

func (c Controller) ListSpecialties() ([]db.Specialty, error) {
	specialties, err := c.queries.ListSpecialties(context.Background())
	if err != nil {
		return nil, err
	}

	return specialties, nil
}

func (c Controller) LinkDoctorToSpecialty(docID, specID uuid.UUID) error {
	if _, err := c.queries.CreateSpecialtyDoctorJunction(context.Background(),
		db.CreateSpecialtyDoctorJunctionParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: docID,
			},
			SpecialtyID: pgtype.UUID{
				Valid: true,
				Bytes: specID,
			},
		},
	); err != nil {
		return err
	}

	return nil
}

func (c Controller) UnlinkDoctorToSpecialty(docID, specID uuid.UUID) error {
	if err := c.queries.DeleteSpecialtyDoctorJunction(context.Background(),
		db.DeleteSpecialtyDoctorJunctionParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: docID,
			},
			SpecialtyID: pgtype.UUID{
				Valid: true,
				Bytes: specID,
			},
		},
	); err != nil {
		return err
	}

	return nil
}
