package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) ListInstitutionUsersByInstitutionID(institutiondID uuid.UUID) ([]db.InstitutionUser, error) {
	dbID := pgtype.UUID{
		Bytes: institutiondID,
		Valid: true,
	}

	institutionUsers, err := c.queries.ListInstitutionUserByInstitutionID(
		context.Background(),
		dbID,
	)
	if err != nil {
		return nil, err
	}

	return institutionUsers, nil
}

func (c Controller) CreateInstitutionUser(user models.InstitutionUserPostRequest) (db.InstitutionUser, error) {
	return c.createInstitutionUser(c.queries, user)
}

func (c Controller) createInstitutionUser(queries *db.Queries, user models.InstitutionUserPostRequest) (db.InstitutionUser, error) {
	switch models.InstitutionUserRole(user.Role) {
	case models.InstitutionUserRoleAdministrador, models.InstitutionUserRoleObservador:
		break
	default:
		return db.InstitutionUser{}, fmt.Errorf("invalid institution user role")
	}

	birthdate, err := time.Parse(constants.ISOLayout, user.Birthdate)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	institutionUser, err := queries.CreateInstitutionUser(
		context.Background(),
		db.CreateInstitutionUserParams{
			InstitutionID: pgtype.UUID{
				Bytes: user.InstitutionID,
				Valid: true,
			},
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			GovID:     user.GovID,
			Birthdate: pgtype.Timestamp{
				Time:  birthdate,
				Valid: true,
			},
			Email:       user.Email,
			Crypt:       user.Password,
			PhoneNumber: user.PhoneNumber,
			Role:        db.InstitutionUserRole(user.Role),
		},
	)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	return institutionUser, nil
}

func (c Controller) UpdateInstitutionUser(user models.InstitutionUserPutRequest) (db.InstitutionUser, error) {
	switch models.InstitutionUserRole(user.Role) {
	case models.InstitutionUserRoleAdministrador, models.InstitutionUserRoleObservador:
		break
	default:
		return db.InstitutionUser{}, fmt.Errorf("invalid institution user role")
	}

	birthdate, err := time.Parse("2006-02-04", user.Birthdate)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	institutionUser, err := c.queries.UpdateInstitutionUserByGovID(
		context.Background(),
		db.UpdateInstitutionUserByGovIDParams{
			InstitutionID: pgtype.UUID{
				Bytes: user.InstitutionID,
				Valid: true,
			},
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			GovID:     user.GovID,
			Birthdate: pgtype.Timestamp{
				Time:  birthdate,
				Valid: true,
			},
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Role:        db.InstitutionUserRole(user.Role),
		},
	)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	return institutionUser, nil
}

func (c Controller) GetInstitutionUserByLogin(email, crypt string) (db.InstitutionUser, error) {
	institutionUser, err := c.queries.GetInstitutionUserByLogin(
		context.Background(),
		db.GetInstitutionUserByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	return institutionUser, nil
}

func (c Controller) GetInstitutionUserByGovID(insitutionID uuid.UUID, govID string) (db.InstitutionUser, error) {
	institutionUser, err := c.queries.GetInstitutionUserByGovAndInstitutionID(
		context.Background(),
		db.GetInstitutionUserByGovAndInstitutionIDParams{
			GovID: govID,
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: insitutionID,
			},
		},
	)
	if err != nil {
		return db.InstitutionUser{}, err
	}

	return institutionUser, nil
}

func (c Controller) DeleteInstitutionUserByInstitutionAndUserID(insitutionID, userID uuid.UUID) error {
	err := c.queries.DeleteInstitutionUserByInsitutionAndUserID(
		context.Background(),
		db.DeleteInstitutionUserByInsitutionAndUserIDParams{
			ID: pgtype.UUID{
				Valid: true,
				Bytes: userID,
			},
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: insitutionID,
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
