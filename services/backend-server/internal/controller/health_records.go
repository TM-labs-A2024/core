package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	charset            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomStringLength = 25
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenRandomString(length int) string {
	return StringWithCharset(length, charset)
}

type CreateHealthRecordArgs struct {
	Type          string
	SpecialtyID   uuid.UUID
	PatientID     uuid.UUID
	ContentFormat string
	Title         string
	Description   string
	DoctorID      uuid.UUID
	Payload       io.Reader
}

type CreateEvolutionArgs struct {
	models.EvolutionRequest
	InstitutionID uuid.UUID
	DoctorID      uuid.UUID
}

func (c Controller) CreateHealthRecord(args CreateHealthRecordArgs) (db.CreateHealthRecordResult, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.CreateHealthRecordResult{}, err
	}
	defer tx.Rollback(context.Background())
	txQuery := c.queries.WithTx(tx)

	res, err := c.createHealthRecord(txQuery, args)
	if err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	if tx.Commit(context.Background()); err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	return res, nil
}

func (c Controller) CreateEvolution(args CreateEvolutionArgs) (db.CreateHealthRecordResult, error) {
	tx, err := c.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		c.logger.Debug("error message1:", slog.String("err", err.Error()))
		return db.CreateHealthRecordResult{}, err
	}
	defer func() {
		tx.Rollback(context.Background())

	}()
	txQuery := c.queries.WithTx(tx)

	data, err := json.Marshal(&args.EvolutionRequest.Payload)
	if err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	res, err := c.createHealthRecord(txQuery, CreateHealthRecordArgs{
		Type:          string(db.HealthRecordTypeEvolucin),
		SpecialtyID:   args.EvolutionRequest.Specialty,
		PatientID:     args.EvolutionRequest.PatientID,
		ContentFormat: "json",
		Title:         args.EvolutionRequest.Title,
		Description:   args.EvolutionRequest.Description,
		DoctorID:      args.DoctorID,
		Payload:       bytes.NewReader(data),
	})
	if err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	patient, err := txQuery.GetPatientByID(context.Background(), res.Patient.ID)
	if err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	if args.Bed != nil {
		patient.Bed = *args.Bed
		patient.InstitutionID = pgtype.UUID{
			Bytes: args.InstitutionID,
			Valid: true,
		}

		if _, err := txQuery.UpdatePatientByID(
			context.Background(),
			db.UpdatePatientByIDParams{
				Firstname:     patient.Firstname,
				Lastname:      patient.Lastname,
				GovID:         patient.GovID,
				Birthdate:     patient.Birthdate,
				Email:         patient.Email,
				PhoneNumber:   patient.PhoneNumber,
				Sex:           patient.Sex,
				Pending:       patient.Pending,
				Status:        db.PatientStatusHospitalizado,
				Bed:           patient.Bed,
				ID:            patient.ID,
				InstitutionID: patient.InstitutionID,
			},
		); err != nil {
			return db.CreateHealthRecordResult{}, err
		}
	}

	if tx.Commit(context.Background()); err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	return res, nil
}

func (c Controller) DeleteHealthRecordByID(id uuid.UUID) error {
	return c.queries.DeleteHealthRecordByID(context.Background(), pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
}

func (c Controller) DeleteHealthRecordDataByID(id uuid.UUID) error {
	return c.queries.DeleteHealthRecordDataByID(context.Background(), pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
}

func (c Controller) GetHealthRecordByID(id uuid.UUID) (db.HealthRecord, error) {
	return c.queries.GetHealthRecordByID(context.Background(), pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
}

func (c Controller) createHealthRecord(txQuery *db.Queries, args CreateHealthRecordArgs) (db.CreateHealthRecordResult, error) {
	doctor, err := txQuery.GetDoctorByID(
		context.TODO(),
		pgtype.UUID{
			Bytes: args.DoctorID,
			Valid: true,
		},
	)
	if err != nil {
		return db.CreateHealthRecordResult{}, fmt.Errorf("could not find doctor: %w", err)
	}

	specialty, err := txQuery.GetSpecialtyByID(
		context.TODO(),
		pgtype.UUID{
			Bytes: args.SpecialtyID,
			Valid: true,
		})
	if err != nil {
		return db.CreateHealthRecordResult{}, fmt.Errorf("could not find specialty: %w", err)
	}

	patient, err := txQuery.GetPatientByID(
		context.TODO(),
		pgtype.UUID{
			Bytes: args.PatientID,
			Valid: true,
		},
	)
	if err != nil {
		return db.CreateHealthRecordResult{}, fmt.Errorf("could not find patient: %w", err)
	}

	ar, err := txQuery.GetAccessRequestsByPatientAndDoctorID(
		context.TODO(),
		db.GetAccessRequestsByPatientAndDoctorIDParams{
			PatientID: patient.ID,
			DoctorID:  doctor.ID,
		},
	)
	if err != nil {
		return db.CreateHealthRecordResult{}, fmt.Errorf("could not find access request: %w", err)
	}

	if !ar.Approved {
		return db.CreateHealthRecordResult{}, fmt.Errorf("only approved doctors can upload health-rcords")
	}

	key, err := c.storage.UploadFile(args.PatientID, args.Payload)
	if err != nil {
		return db.CreateHealthRecordResult{}, err
	}

	encryptedURL, err := utils.GetAESEncrypted(
		key,
		patient.PrivateKey,
		c.ivEncryptionKey,
	)
	if err != nil {
		c.DeleteFile(key)
		return db.CreateHealthRecordResult{}, err
	}

	hr, err := c.queries.CreateHealthRecord(
		context.Background(),
		db.CreateHealthRecordParams{
			PatientID:     patient.ID,
			Type:          db.HealthRecordType(args.Type),
			SpecialtyID:   specialty.ID,
			ContentFormat: args.ContentFormat,
			Title:         args.Title,
			Description:   args.Description,
			Author:        doctor.Firstname + " " + doctor.Lastname,
			PublicKey: pgtype.Text{
				String: encryptedURL,
				Valid:  encryptedURL != "",
			},
		},
	)
	if err != nil {
		c.DeleteFile(key)
		return db.CreateHealthRecordResult{}, err
	}

	if err := c.blockchain.CreateHealthRecord(
		patient.BlockchainAddress,
		encryptedURL,
	); err != nil {
		c.DeleteFile(key)
		return db.CreateHealthRecordResult{}, fmt.Errorf("could not insert to blockchain: %w", err)
	}

	return db.CreateHealthRecordResult{
		HealthRecord: hr,
		Specialty:    specialty,
		Patient:      patient,
	}, nil
}
