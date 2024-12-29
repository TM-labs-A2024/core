package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c Controller) CreatePatient(req models.PatientPostRequest) (db.Patient, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.Patient{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Patient{}, err
	}

	patient, err := txQuery.CreatePatient(context.Background(), db.CreatePatientParams{
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		GovID:       req.GovID,
		Birthdate:   pgtype.Timestamp{Valid: true, Time: birthdate},
		Email:       req.Email,
		Crypt:       req.Password,
		PhoneNumber: req.PhoneNumber,
		Sex:         req.Sex,
		Pending:     req.Pending,
		Status:      db.PatientStatus(req.Status),
		Bed:         req.Bed,
	})
	if err != nil {
		return db.Patient{}, err
	}

	address, err := utils.GenerateKey(patient.ID.Bytes, c.ivEncryptionKey)
	if err != nil {
		return db.Patient{}, err
	}

	privKey, err := utils.GenerateKey(patient.ID.Bytes, c.ivEncryptionKey)
	if err != nil {
		return db.Patient{}, err
	}

	patient, err = txQuery.SetPatientAddressAndPrivateKey(
		context.Background(),
		db.SetPatientAddressAndPrivateKeyParams{
			BlockchainAddress: address,
			PrivateKey:        privKey,
			ID:                patient.ID,
		},
	)
	if err != nil {
		return db.Patient{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) UpdatePatientByID(req models.PatientPutRequest) (db.Patient, error) {
	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Patient{}, err
	}

	institutionId := pgtype.UUID{}
	if req.InstitutionID != uuid.Nil {
		institutionId = pgtype.UUID{
			Bytes: req.InstitutionID,
			Valid: true,
		}
	}

	patient, err := c.queries.UpdatePatientByID(context.Background(), db.UpdatePatientByIDParams{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		GovID:     req.GovID,
		Birthdate: pgtype.Timestamp{
			Time:  birthdate,
			Valid: true,
		},
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Sex:         req.Sex,
		Pending:     req.Pending,
		Status:      db.PatientStatus(req.Status),
		Bed:         req.Bed,
		ID: pgtype.UUID{
			Bytes: req.ID,
			Valid: true,
		},
		InstitutionID: institutionId,
	})
	if err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) DeletePatientByID(id uuid.UUID) error {
	if err := c.queries.DeletePatientByID(
		context.Background(),
		pgtype.UUID{Valid: true, Bytes: id},
	); err != nil {
		return err
	}

	return nil
}

func (c Controller) GetPatientByLogin(email, crypt string) (db.Patient, error) {
	patient, err := c.queries.GetPatientByLogin(
		context.Background(),
		db.GetPatientByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) GetPatientByGovID(govID string) (db.Patient, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govID)
	if err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) GetPatientByID(id uuid.UUID) (db.Patient, error) {
	patientId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	patient, err := c.queries.GetPatientByID(context.Background(), patientId)
	if err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) ListPatients(doctorID uuid.UUID) ([]db.Patient, error) {
	patients, err := c.queries.ListPatients(context.Background())
	if err != nil {
		return nil, err
	}

	ars, err := c.queries.ListAccessRequestsByDoctorID(context.Background(), pgtype.UUID{
		Bytes: doctorID,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}

	approvedPatientLedger := map[pgtype.UUID]bool{}
	for _, ar := range ars {
		approvedPatientLedger[ar.PatientID] = ar.Approved
	}

	result := []db.Patient{}
	for _, patient := range patients {
		patient.Pending = false
		if approved, ok := approvedPatientLedger[patient.ID]; ok && approved {
			continue
		} else if ok && !approved {
			patient.Pending = true
		}

		result = append(result, patient)
	}

	return result, nil
}

func (c Controller) ListHealthRecordPatientsGovID(govID string) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govID)
	if err != nil {
		return nil, err
	}

	healthRecords, err := c.queries.ListHealthRecordsByPatientID(
		context.Background(),
		patient.ID,
	)
	if err != nil {
		return nil, err
	}

	return healthRecords, nil
}

func (c Controller) ListPatientApprovedDoctorsByGovID(govID string) ([]db.Doctor, error) {
	patient, err := c.queries.GetPatientByGovID(
		context.Background(),
		govID,
	)
	if err != nil {
		return nil, err
	}

	accessRequests, err := c.queries.ListApprovedAccessRequestsByPatientID(
		context.Background(),
		patient.ID,
	)
	if err != nil {
		return nil, err
	}

	doctors := []db.Doctor{}
	for _, request := range accessRequests {
		doctor, err := c.queries.GetDoctorByID(
			context.Background(),
			request.DoctorID,
		)
		if err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (c Controller) CreateAccessRequestWithDoctorAndPatientID(doctorID, patientID uuid.UUID) (db.DoctorAccessRequest, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.DoctorAccessRequest{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	ar, err := txQuery.CreateAccessRequest(
		context.Background(),
		db.CreateAccessRequestParams{
			PatientID: pgtype.UUID{
				Valid: true,
				Bytes: patientID,
			},
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctorID,
			},
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	doctor, err := txQuery.GetDoctorByID(context.Background(), ar.DoctorID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	if _, err = txQuery.UpdateDoctorByID(
		context.Background(),
		db.UpdateDoctorByIDParams{
			InstitutionID:  doctor.InstitutionID,
			Firstname:      doctor.Firstname,
			Lastname:       doctor.Lastname,
			GovID:          doctor.GovID,
			Birthdate:      doctor.Birthdate,
			Email:          doctor.Email,
			PhoneNumber:    doctor.PhoneNumber,
			Credentials:    doctor.Credentials,
			Pending:        doctor.Pending,
			PatientPending: true,
			Sex:            doctor.Sex,
			ID:             doctor.ID,
		},
	); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	patient, err := txQuery.GetPatientByID(context.Background(), ar.PatientID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	if txQuery.UpdatePatientByID(context.Background(), db.UpdatePatientByIDParams{
		Firstname:     patient.Firstname,
		Lastname:      patient.Lastname,
		GovID:         patient.GovID,
		Birthdate:     patient.Birthdate,
		Email:         patient.Email,
		PhoneNumber:   patient.PhoneNumber,
		Sex:           patient.Sex,
		Pending:       true,
		Status:        patient.Status,
		Bed:           patient.Bed,
		InstitutionID: patient.InstitutionID,
		ID:            patient.ID,
	}); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) ListAccessRequestsByPatientID(patientID uuid.UUID) ([]db.DoctorAccessRequest, error) {
	ars, err := c.queries.ListAccessRequestsByPatientID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: patientID,
		},
	)
	if err != nil {
		return nil, err
	}

	return ars, nil
}

func (c Controller) GetAccessRequestByID(id uuid.UUID) (db.DoctorAccessRequest, error) {
	ar, err := c.queries.GetAccessRequestsByID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: id,
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) GetAccessRequestByPatientAndDoctorID(doctorID, patientID uuid.UUID) (db.DoctorAccessRequest, error) {
	ar, err := c.queries.GetAccessRequestsByPatientAndDoctorID(
		context.Background(),
		db.GetAccessRequestsByPatientAndDoctorIDParams{
			PatientID: pgtype.UUID{
				Valid: true,
				Bytes: patientID,
			},
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctorID,
			},
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) DeleteAccessRequestsByID(id uuid.UUID) error {
	if err := c.queries.DeleteAccessRequestByID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: id,
		},
	); err != nil {
		return err
	}

	return nil
}

func (c Controller) DenyAccessRequestsByID(ar db.DoctorAccessRequest) (db.DoctorAccessRequest, error) {
	return c.updateAccessRequestByID(ar, false)
}

func (c Controller) ApproveAccessRequestsByID(ar db.DoctorAccessRequest) (db.DoctorAccessRequest, error) {
	return c.updateAccessRequestByID(ar, true)
}

func (c Controller) updateAccessRequestByID(ar db.DoctorAccessRequest, approve bool) (db.DoctorAccessRequest, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.DoctorAccessRequest{}, err
	}

	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	ar, err = txQuery.UpdateAccessRequestByID(
		context.Background(),
		db.UpdateAccessRequestByIDParams{
			PatientID: ar.PatientID,
			DoctorID:  ar.DoctorID,
			Pending:   false,
			Approved:  approve,
			ID:        ar.ID,
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	doctor, err := txQuery.GetDoctorByID(context.Background(), ar.DoctorID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	count, err := txQuery.CountPendingAccessRequestsByDoctorID(context.Background(), doctor.ID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	doctor.PatientPending = count > 0
	if _, err = txQuery.UpdateDoctorByID(
		context.Background(),
		db.UpdateDoctorByIDParams{
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
		},
	); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	patient, err := txQuery.GetPatientByID(context.Background(), ar.PatientID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	count, err = txQuery.CountPendingAccessRequestsByPatientID(context.Background(), patient.ID)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	patient.Pending = count > 0
	if txQuery.UpdatePatientByID(context.Background(), db.UpdatePatientByIDParams{
		Firstname:     patient.Firstname,
		Lastname:      patient.Lastname,
		GovID:         patient.GovID,
		Birthdate:     patient.Birthdate,
		Email:         patient.Email,
		PhoneNumber:   patient.PhoneNumber,
		Sex:           patient.Sex,
		Pending:       patient.Pending,
		Status:        patient.Status,
		Bed:           patient.Bed,
		InstitutionID: patient.InstitutionID,
		ID:            patient.ID,
	}); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) ListHealthRecordByPatientsGovAndSpecialtyID(govID string, specialtyID uuid.UUID) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govID)
	if err != nil {
		return nil, err
	}

	healthRecords, err := c.queries.ListHealthRecordsBySpecialtyAndPatientID(
		context.Background(),
		db.ListHealthRecordsBySpecialtyAndPatientIDParams{
			PatientID: patient.ID,
			SpecialtyID: pgtype.UUID{
				Valid: true,
				Bytes: specialtyID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return healthRecords, nil
}

func (c Controller) ListOrdersByPatientGovID(govID string) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govID)
	if err != nil {
		return nil, err
	}

	healthRecords, err := c.queries.ListHealthRecordsByTypeAndPatientID(
		context.Background(),
		db.ListHealthRecordsByTypeAndPatientIDParams{
			PatientID: patient.ID,
			Type:      db.HealthRecordTypeOrden,
		},
	)
	if err != nil {
		return nil, err
	}

	return healthRecords, nil
}

func (c *Controller) ListPatientsByInstitutionID(institutionID uuid.UUID) ([]db.Patient, error) {
	patients, err := c.queries.ListPatientsByInstitutionID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: institutionID,
		},
	)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
