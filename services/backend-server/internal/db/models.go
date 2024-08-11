// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type HealthRecordType string

const (
	HealthRecordTypeEvolucin HealthRecordType = "evolución"
	HealthRecordTypeAnlisis  HealthRecordType = "análisis"
	HealthRecordTypeOrden    HealthRecordType = "orden"
)

func (e *HealthRecordType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = HealthRecordType(s)
	case string:
		*e = HealthRecordType(s)
	default:
		return fmt.Errorf("unsupported scan type for HealthRecordType: %T", src)
	}
	return nil
}

type NullHealthRecordType struct {
	HealthRecordType HealthRecordType
	Valid            bool // Valid is true if HealthRecordType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHealthRecordType) Scan(value interface{}) error {
	if value == nil {
		ns.HealthRecordType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.HealthRecordType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHealthRecordType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.HealthRecordType), nil
}

type InstitutionType string

const (
	InstitutionTypeHospital InstitutionType = "hospital"
	InstitutionTypeClnica   InstitutionType = "clínica"
)

func (e *InstitutionType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = InstitutionType(s)
	case string:
		*e = InstitutionType(s)
	default:
		return fmt.Errorf("unsupported scan type for InstitutionType: %T", src)
	}
	return nil
}

type NullInstitutionType struct {
	InstitutionType InstitutionType
	Valid           bool // Valid is true if InstitutionType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullInstitutionType) Scan(value interface{}) error {
	if value == nil {
		ns.InstitutionType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.InstitutionType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullInstitutionType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.InstitutionType), nil
}

type InstitutionUserRole string

const (
	InstitutionUserRoleAdministrador InstitutionUserRole = "administrador"
	InstitutionUserRoleObservador    InstitutionUserRole = "observador"
)

func (e *InstitutionUserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = InstitutionUserRole(s)
	case string:
		*e = InstitutionUserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for InstitutionUserRole: %T", src)
	}
	return nil
}

type NullInstitutionUserRole struct {
	InstitutionUserRole InstitutionUserRole
	Valid               bool // Valid is true if InstitutionUserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullInstitutionUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.InstitutionUserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.InstitutionUserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullInstitutionUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.InstitutionUserRole), nil
}

type PatientStatus string

const (
	PatientStatusHospitalizado PatientStatus = "hospitalizado"
	PatientStatusRegular       PatientStatus = "regular"
)

func (e *PatientStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PatientStatus(s)
	case string:
		*e = PatientStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PatientStatus: %T", src)
	}
	return nil
}

type NullPatientStatus struct {
	PatientStatus PatientStatus
	Valid         bool // Valid is true if PatientStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPatientStatus) Scan(value interface{}) error {
	if value == nil {
		ns.PatientStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PatientStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPatientStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PatientStatus), nil
}

type Doctor struct {
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	ID             pgtype.UUID
	InstitutionID  pgtype.UUID
	Firstname      string
	Lastname       string
	GovID          string
	Birthdate      pgtype.Timestamp
	Email          string
	Sex            string
	Password       string
	PhoneNumber    string
	Credentials    string
	Pending        bool
	PatientPending bool
}

type DoctorAccessRequest struct {
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	ID        pgtype.UUID
	PatientID pgtype.UUID
	DoctorID  pgtype.UUID
	Pending   bool
	Approved  bool
}

type DoctorSpecialty struct {
	DoctorID    pgtype.UUID
	SpecialtyID pgtype.UUID
}

type Government struct {
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	ID        pgtype.UUID
	Email     string
	Password  string
}

type GovernmentEnrollmentRequest struct {
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	ID            pgtype.UUID
	InstitutionID pgtype.UUID
	GovernmentID  pgtype.UUID
	Pending       bool
	Approved      bool
}

type HealthRecord struct {
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	ID            pgtype.UUID
	PatientID     pgtype.UUID
	Author        string
	Title         string
	Description   string
	PublicKey     pgtype.Text
	Type          HealthRecordType
	SpecialtyID   pgtype.UUID
	ContentFormat string
}

type Institution struct {
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
	ID           pgtype.UUID
	GovernmentID pgtype.UUID
	Name         string
	Address      string
	Credentials  string
	Type         InstitutionType
	GovID        string
	Pending      bool
}

type InstitutionEnrollmentRequest struct {
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	ID            pgtype.UUID
	InstitutionID pgtype.UUID
	DoctorID      pgtype.UUID
	NurseID       pgtype.UUID
	Pending       bool
	Approved      bool
}

type InstitutionUser struct {
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	ID            pgtype.UUID
	InstitutionID pgtype.UUID
	Firstname     string
	Lastname      string
	GovID         string
	Birthdate     pgtype.Timestamp
	Email         string
	Password      string
	PhoneNumber   string
	Role          InstitutionUserRole
}

type Nurse struct {
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	ID            pgtype.UUID
	InstitutionID pgtype.UUID
	Firstname     string
	Lastname      string
	GovID         string
	Birthdate     pgtype.Timestamp
	Sex           string
	Email         string
	Password      string
	PhoneNumber   string
	Credentials   string
	Pending       bool
}

type Patient struct {
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
	ID                pgtype.UUID
	InstitutionID     pgtype.UUID
	Firstname         string
	Lastname          string
	GovID             string
	Birthdate         pgtype.Timestamp
	Email             string
	Password          string
	PhoneNumber       string
	Sex               string
	Pending           bool
	Status            PatientStatus
	Bed               string
	PrivateKey        string
	BlockchainAddress string
}

type Specialty struct {
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	ID          pgtype.UUID
	Description string
	Name        string
}
