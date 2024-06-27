package controller

import (
	"context"
	"fmt"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) ListApprovedInstitutions() ([]db.Institution, error) {
	c.logger.Debug("ListApprovedInstitutions")
	return c.queries.ListApprovedInstitutions(context.Background())
}

func (c Controller) GetFirstInstitutionUser(institutionID uuid.UUID) (db.InstitutionUser, error) {
	c.logger.Debug("GetFirstInstitutionUser")
	return c.queries.GetFirstInstitutionUserByInstitutionID(context.Background(), pgtype.UUID{
		Bytes: institutionID,
		Valid: true,
	})
}

func (c Controller) CreateInstitutionEnrollmentRequestDoctor(docId, instId uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	er, err := c.queries.CreateInstitutionEnrollmentRequest(
		context.Background(),
		db.CreateInstitutionEnrollmentRequestParams{
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: docId,
			},
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: instId,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) CreateInstitutionEnrollmentRequestNurse(nurseId, instId uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	er, err := c.queries.CreateInstitutionEnrollmentRequest(
		context.Background(),
		db.CreateInstitutionEnrollmentRequestParams{
			InstitutionID: pgtype.UUID{
				Valid: true,
				Bytes: instId,
			},
			NurseID:  pgtype.UUID{
				Valid: true,
				Bytes: nurseId,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) GetInstitutionEnrollmentRequestsByID(erID uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	c.logger.Debug("GetInstitutionEnrollmentRequestsByID")
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

func (c Controller) ApproveInstitutionEnrollmentRequestsByID(erByID db.InstitutionEnrollmentRequest) (db.InstitutionEnrollmentRequest, error) {
	c.logger.Debug("UpdateInstitutionEnrollmentRequestsByID")

	er, err := c.queries.UpdateInstitutionEnrollmentRequestByID(
		context.Background(),
		db.UpdateInstitutionEnrollmentRequestByIDParams{
			InstitutionID: erByID.InstitutionID,
			DoctorID:      erByID.DoctorID,
			NurseID:       erByID.NurseID,
			Pending:       false,
			Approved:      true,
			ID:            erByID.ID,
		})
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return er, nil
}

func (c Controller) DenyInstitutionEnrollmentRequestsByID(erByID db.InstitutionEnrollmentRequest) (db.InstitutionEnrollmentRequest, error) {
	c.logger.Debug("UpdateInstitutionEnrollmentRequestsByID")

	er, err := c.queries.UpdateInstitutionEnrollmentRequestByID(
		context.Background(),
		db.UpdateInstitutionEnrollmentRequestByIDParams{
			InstitutionID: erByID.InstitutionID,
			DoctorID:      erByID.DoctorID,
			NurseID:       erByID.NurseID,
			Pending:       false,
			Approved:      false,
			ID:            erByID.ID,
		})
	if err != nil {
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

func (c Controller) ListInstitutions() ([]db.Institution, error) {
	c.logger.Debug("ListInstitutions")
	institutions, err := c.queries.ListInstitutions(context.Background())
	if err != nil {
		return nil, err
	}

	return institutions, nil
}

func (c Controller) GetInstitutionByGovID(govID string) (db.Institution, error) {
	c.logger.Debug("GetInstitutionByGovID")
	c.logger.Debug(govID)
	institutions, err := c.queries.GetInstitutionByGovID(context.Background(), govID)
	if err != nil {
		return db.Institution{}, err
	}

	return institutions, nil
}

func (c Controller) DeleteInstitutionByID(id uuid.UUID) error {
	c.logger.Debug("DeleteInstitutionByID")
	dbID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	if err := c.queries.DeleteInstitutionByID(context.Background(), dbID); err != nil {
		return err
	}

	return nil
}

func (c Controller) CreateInstitution(institution models.Institution) (db.Institution, error) {
	c.logger.Debug("CreateInstitution")
	switch db.InstitutionType(institution.Type) {
	case db.InstitutionTypeClinic, db.InstitutionTypeHospital:
		break
	default:
		return db.Institution{}, fmt.Errorf("invalid intitution type: %s", institution.Type)
	}

	inst, err := c.queries.CreateInstitution(context.Background(), db.CreateInstitutionParams{
		Name:        institution.Name,
		Address:     institution.Address,
		Credentials: institution.Credentials,
		Type:        db.InstitutionType(institution.Type),
		GovID:       institution.GovId,
	})
	if err != nil {
		return db.Institution{}, err
	}

	_, err = c.queries.CreateGovernmentEnrollmentRequests(context.Background(), inst.ID)
	if err != nil {
		return db.Institution{}, err
	}

	return inst, nil
}

func (c Controller) UpdateInstitution(institution models.Institution, id uuid.UUID) (db.Institution, error) {
	c.logger.Debug("UpdateInstitution")
	switch db.InstitutionType(institution.Type) {
	case db.InstitutionTypeClinic, db.InstitutionTypeHospital:
		break
	default:
		return db.Institution{}, fmt.Errorf("invalid intitution type: %s", institution.Type)
	}

	return c.queries.UpdateInstitutionByID(context.Background(), db.UpdateInstitutionByIDParams{
		Name:        institution.Name,
		Address:     institution.Address,
		Credentials: institution.Credentials,
		Type:        db.InstitutionType(institution.Type),
		GovID:       institution.GovId,
		Pending:     false,
		ID: pgtype.UUID{
			Bytes: id,
			Valid: true,
		},
	})
}

func (c Controller) GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID(doctorId, institutionId uuid.UUID) (db.InstitutionEnrollmentRequest, error) {
	enrollmentRequest, err := c.queries.GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID(
		context.Background(),
		db.GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionIDParams{
			DoctorID: pgtype.UUID{
				Bytes: doctorId,
				Valid: true,
			},
			InstitutionID: pgtype.UUID{
				Bytes: institutionId,
				Valid: true,
			},
		},
	)
	if err != nil {
		return db.InstitutionEnrollmentRequest{}, err
	}

	return enrollmentRequest, nil
}
