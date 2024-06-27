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

func (c Controller) CreatePatient(req models.PatientPostRequest) (db.Patient, error) {
	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Patient{}, err
	}

	patient, err := c.queries.CreatePatient(context.Background(), db.CreatePatientParams{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		GovID:     req.GovId,
		Birthdate: pgtype.Timestamp{
			Valid: true,
			Time:  birthdate,
		},
		Email:       req.Email,
		Crypt:       req.Password,
		PhoneNumber: req.PhoneNumber,
		Sex:         req.Sex,
		Pending:     req.Pending,
	})
	if err != nil {
		return db.Patient{}, err
	}

	return patient, nil
}

func (c Controller) UpdatePatientByID(req models.PatientPutRequest) (db.Patient, error) {
	birthdate, err := time.Parse(constants.ISOLayout, req.Birthdate)
	if err != nil {
		return db.Patient{}, err
	}

	patient, err := c.queries.UpdatePatientByID(context.Background(), db.UpdatePatientByIDParams{
		ID: pgtype.UUID{
			Valid: true,
			Bytes: req.ID,
		},
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		GovID:     req.GovId,
		Birthdate: pgtype.Timestamp{
			Valid: true,
			Time:  birthdate,
		},
		Email:       req.Email,
		Crypt:       req.Password,
		PhoneNumber: req.PhoneNumber,
		Sex:         req.Sex,
		Pending:     req.Pending,
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
	nurse, err := c.queries.GetPatientByLogin(
		context.Background(),
		db.GetPatientByLoginParams{
			Email: email,
			Crypt: crypt,
		},
	)
	if err != nil {
		return db.Patient{}, err
	}

	return nurse, nil
}

func (c Controller) GetPatientByGovID(govId string) (db.Patient, error) {
	nurse, err := c.queries.GetPatientByGovID(context.Background(), govId)
	if err != nil {
		return db.Patient{}, err
	}

	return nurse, nil
}

func (c Controller) ListPatients() ([]db.Patient, error) {
	patients, err := c.queries.ListPatients(context.Background())
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (c Controller) ListHealthRecordPatientsGovID(govId string) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govId)
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

func (c Controller) ListPatientApprovedDoctors(govId string) ([]db.Doctor, error) {
	patient, err := c.queries.GetPatientByGovID(
		context.Background(),
		govId,
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

func (c Controller) CreateAccessRequestWithDoctorAndPatientID(doctorId, patientId uuid.UUID) (db.DoctorAccessRequest, error) {
	ar, err := c.queries.CreateAccessRequest(
		context.Background(),
		db.CreateAccessRequestParams{
			PatientID: pgtype.UUID{
				Valid: true,
				Bytes: patientId,
			},
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctorId,
			},
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) ListAccessRequestsByPatientId(patientId uuid.UUID) ([]db.DoctorAccessRequest, error) {
	ars, err := c.queries.ListAccessRequestsByPatientID(
		context.Background(),
		pgtype.UUID{
			Valid: true,
			Bytes: patientId,
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

func (c Controller) GetAccessRequestByPatientAndDoctorID(doctorId, patientId uuid.UUID) (db.DoctorAccessRequest, error) {
	ar, err := c.queries.GetAccessRequestsByPatientAndDoctorID(
		context.Background(),
		db.GetAccessRequestsByPatientAndDoctorIDParams{
			PatientID: pgtype.UUID{
				Valid: true,
				Bytes: patientId,
			},
			DoctorID: pgtype.UUID{
				Valid: true,
				Bytes: doctorId,
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
	ar, err := c.queries.UpdateAccessRequestByID(
		context.Background(),
		db.UpdateAccessRequestByIDParams{
			PatientID: ar.PatientID,
			DoctorID:  ar.DoctorID,
			Pending:   false,
			Approved:  false,
			ID:        ar.ID,
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) ApproveAccessRequestsByID(ar db.DoctorAccessRequest) (db.DoctorAccessRequest, error) {
	var err error
	ar, err = c.queries.UpdateAccessRequestByID(
		context.Background(),
		db.UpdateAccessRequestByIDParams{
			PatientID: ar.PatientID,
			DoctorID:  ar.DoctorID,
			Pending:   false,
			Approved:  true,
			ID:        ar.ID,
		},
	)
	if err != nil {
		return db.DoctorAccessRequest{}, err
	}

	return ar, nil
}

func (c Controller) ListHealthRecordByPatientsGovAndSpecialtyID(govId string, specialtyId uuid.UUID) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govId)
	if err != nil {
		return nil, err
	}

	healthRecords, err := c.queries.ListHealthRecordsBySpecialtyAndPatientID(
		context.Background(),
		db.ListHealthRecordsBySpecialtyAndPatientIDParams{
			PatientID: patient.ID,
			SpecialtyID: pgtype.UUID{
				Valid: true,
				Bytes: specialtyId,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return healthRecords, nil
}

func (c Controller) ListOrdersByPatientGovID(govId string) ([]db.HealthRecord, error) {
	patient, err := c.queries.GetPatientByGovID(context.Background(), govId)
	if err != nil {
		return nil, err
	}

	healthRecords, err := c.queries.ListHealthRecordsByTypeAndPatientID(
		context.Background(),
		db.ListHealthRecordsByTypeAndPatientIDParams{
			PatientID: patient.ID,
			Type:      "order",
		},
	)
	if err != nil {
		return nil, err
	}

	return healthRecords, nil
}
