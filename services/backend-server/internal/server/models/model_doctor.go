package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/constants"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Doctor struct {
	InstitutionID uuid.UUID `json:"institutionId"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	GovId         string    `json:"govId"`
	Birthdate     string    `json:"birthdate"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	Credentials   string    `json:"credentials"`
	Sex           string    `json:"sex"`
}

// DoctorEnrollmentRequest <= NOT USED ON API, ONLY db
// type DoctorEnrollmentRequest struct {
// 	InstitutionUUID  uuid.UUID `json:"institutionId"`
// 	DoctorUUID       uuid.UUID `json:"doctorId"`
// 	Pending          bool   `json:"pending"`
// 	Approved         bool   `json:"approved"`
// 	ProfessionalType string `json:"professional-type"`
// }

type DoctorAccessRequest struct {
	PatientID uuid.UUID `json:"patientId"`
	DoctorID  uuid.UUID `json:"doctorId"`
	Pending   bool      `json:"pending"`
	Approved  bool      `json:"approved"`
}

type DoctorPutAccessRequest struct {
	ID uuid.UUID `json:"id"`
	DoctorAccessRequest
}

type DoctorAccessResponse DoctorPutAccessRequest

func NewDoctorAccessResponse(dar db.DoctorAccessRequest) (DoctorAccessResponse, error) {
	darID, err := uuid.FromBytes(dar.ID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	doctorID, err := uuid.FromBytes(dar.DoctorID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	patientID, err := uuid.FromBytes(dar.PatientID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	return DoctorAccessResponse{
		ID: darID,
		DoctorAccessRequest: DoctorAccessRequest{
			PatientID: patientID,
			DoctorID:  doctorID,
			Pending:   dar.Pending,
			Approved:  dar.Approved,
		},
	}, nil
}

type DoctorsPostRequest struct {
	Doctor
	Password    string      `json:"password"`
	Specialties []uuid.UUID `json:"specialties"`
}

type DoctorsPutRequest struct {
	ID uuid.UUID `json:"id"`
	DoctorsPostRequest
	Approved bool `json:"approved"`
}

type DoctorsResponse struct {
	ID uuid.UUID `json:"id"`
	Doctor
	Specialities   []Specialty `json:"specialities"`
	Pending        bool        `json:"pending"`
	PatientPending bool        `json:"patientPending"`
}

type NewDoctorResponseArgs struct {
	Doctor         db.Doctor
	Specialties    []db.Specialty
	PatientPending bool
	Pending        bool
}

func NewDoctorResponse(args NewDoctorResponseArgs) (DoctorsResponse, error) {
	doctorID, err := uuid.FromBytes(args.Doctor.ID.Bytes[:])
	if err != nil {
		return DoctorsResponse{}, err
	}

	institutionID, err := uuid.FromBytes(args.Doctor.InstitutionID.Bytes[:])
	if err != nil {
		return DoctorsResponse{}, err
	}

	// Fetch specialties
	specialties := []Specialty{}
	for _, s := range args.Specialties {
		specialties = append(specialties, Specialty{
			ID:          s.ID.Bytes,
			Description: s.Description,
			Name:        SpecialtyName(s.Name),
		})
	}

	return DoctorsResponse{
		ID: doctorID,
		Doctor: Doctor{
			InstitutionID: institutionID,
			Firstname:     args.Doctor.Firstname,
			Lastname:      args.Doctor.Lastname,
			GovId:         args.Doctor.GovID,
			Birthdate:     args.Doctor.Birthdate.Time.Format(constants.ISOLayout),
			Email:         args.Doctor.Email,
			PhoneNumber:   args.Doctor.PhoneNumber,
			Credentials:   args.Doctor.Credentials,
			Sex:           args.Doctor.Sex,
		},
		Specialities:   specialties,
		Pending:        args.Doctor.Pending,
		PatientPending: args.Doctor.PatientPending,
	}, nil
}
