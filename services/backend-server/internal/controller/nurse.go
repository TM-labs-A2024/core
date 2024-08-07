package controller

import (
	"context"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) DeleteNurseByID(nurseID uuid.UUID) error {
	return c.queries.DeleteNurseByID(context.Background(), pgtype.UUID{
		Bytes: nurseID,
		Valid: true,
	})
}

func (c Controller) GetNurseByID(nurseID uuid.UUID) (db.Nurse, error) {
	return c.queries.GetNurseByID(context.Background(), pgtype.UUID{
		Bytes: nurseID,
		Valid: true,
	})
}

func (c Controller) GetNurseByLogin(email, crypt string) (db.Nurse, error) {
	nurse, err := c.queries.GetNurseByLogin(
		context.Background(),
		db.GetNurseByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.Nurse{}, err
	}

	return nurse, nil
}

func (c Controller) CreateNurse(req models.NursePostRequest) (db.Nurse, error) {
	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Nurse{}, err
	}

	nurse, err := c.queries.CreateNurse(context.Background(), db.CreateNurseParams{
		InstitutionID: pgtype.UUID{
			Valid: true,
			Bytes: req.InstitutionID,
		},
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		GovID:     req.GovID,
		Birthdate: pgtype.Timestamp{
			Valid: true,
			Time:  birthdate,
		},
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Credentials: req.Credentials,
		Crypt:       req.Password,
		Sex:         req.Sex,
	})
	if err != nil {
		return db.Nurse{}, err
	}

	return nurse, nil
}
func (c Controller) UpdateNurse(req models.NursesPutRequest) (db.Nurse, error) {
	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Nurse{}, err
	}

	nurse, err := c.queries.UpdateNurseByID(context.Background(), db.UpdateNurseByIDParams{
		InstitutionID: pgtype.UUID{
			Valid: true,
			Bytes: req.InstitutionID,
		},
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		GovID:     req.GovID,
		Birthdate: pgtype.Timestamp{
			Valid: true,
			Time:  birthdate,
		},
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Credentials: req.Credentials,
		ID: pgtype.UUID{
			Valid: true,
			Bytes: req.ID,
		},
		Sex: req.Sex,
	})
	if err != nil {
		return db.Nurse{}, err
	}

	return nurse, nil
}

func (c Controller) ListNurses() ([]db.Nurse, error) {
	return c.queries.ListNurses(context.Background())
}

func (c Controller) ListNursesByInstitutionID(institutionID uuid.UUID) ([]db.Nurse, error) {
	return c.queries.ListNursesByInstitutionID(context.Background(), pgtype.UUID{
		Valid: true,
		Bytes: institutionID,
	})
}
