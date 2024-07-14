package controller

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
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
	SpecialtyId   uuid.UUID
	PatientId     uuid.UUID
	ContentFormat string
	Title         string
	Description   string
	Author        string
}

func (c Controller) CreateHealthRecord(args CreateHealthRecordArgs) (db.HealthRecord, error) {
	h := sha256.New()
	h.Write([]byte(GenRandomString(25)))
	publicKey := fmt.Sprintf("%x", h.Sum(nil))

	h.Write([]byte(GenRandomString(25)))
	privateKey := fmt.Sprintf("%x", h.Sum(nil))

	return c.queries.CreateHealthRecord(context.Background(), db.CreateHealthRecordParams{
		PatientID: pgtype.UUID{
			Valid: true,
			Bytes: args.PatientId,
		},
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Type:       args.Type,
		SpecialtyID: pgtype.UUID{
			Valid: true,
			Bytes: args.SpecialtyId,
		},
		ContentFormat: args.ContentFormat,
		Title:         args.Title,
		Description:   args.Description,
		Author:        args.Author,
	})
}

func (c Controller) DeleteHealthRecordByID(id uuid.UUID) error {
	return c.queries.DeleteHealthRecordByID(context.Background(), pgtype.UUID{
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
