package models

import (
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/google/uuid"
)

type Doctor struct {
	InstitutionUUID uuid.UUID `json:"institution_uuid,omitempty"`
	Firstname       string    `json:"firstname,omitempty"`
	Lastname        string    `json:"lastname,omitempty"`
	GovId           string    `json:"gov_id,omitempty"`
	Birthdate       string    `json:"birthdate,omitempty"`
	Email           string    `json:"email,omitempty"`
	PhoneNumber     string    `json:"phone_number,omitempty"`
	Credentials     string    `json:"credentials,omitempty"`
}

// DoctorEnrollmentRequest <= NOT USED ON API, ONLY DB
// type DoctorEnrollmentRequest struct {
// 	InstitutionUUID  uuid.UUID `json:"institution_uuid,omitempty"`
// 	DoctorUUID       uuid.UUID `json:"doctor_uuid,omitempty"`
// 	Pending          bool   `json:"pending,omitempty"`
// 	Approved         bool   `json:"approved,omitempty"`
// 	ProfessionalType string `json:"professional-type,omitempty"`
// }

type DoctorAccessRequest struct {
	PatientUUID uuid.UUID `json:"patient_uuid,omitempty"`
	DoctorUUID  uuid.UUID `json:"doctor_uuid,omitempty"`
}

type DoctorAccessResponse struct {
	UUID uuid.UUID `json:"uuid"`
	DoctorAccessRequest
}

func NewDoctorAccessResponse(dar db.DoctorAccessRequest) (DoctorAccessResponse, error) {
	darUUID, err := uuid.FromBytes(dar.ID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	doctorUUID, err := uuid.FromBytes(dar.DoctorID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	patientUUID, err := uuid.FromBytes(dar.DoctorID.Bytes[:])
	if err != nil {
		return DoctorAccessResponse{}, err
	}

	return DoctorAccessResponse{
		UUID: darUUID,
		DoctorAccessRequest: DoctorAccessRequest{
			PatientUUID: patientUUID,
			DoctorUUID:  doctorUUID,
		},
	}, nil
}

type DoctorsPostRequest struct {
	Doctor
	Password    string `json:"password,omitempty"`
	Specialties []int  `json:"specialties,omitempty"`
}

type DoctorsPutRequest struct {
	UUID string `json:"uuid,omitempty"`
	DoctorsPostRequest
	Approved bool `json:"approved,omitempty"`
}

type DoctorsResponse struct {
	UUID uuid.UUID `json:"uuid,omitempty"`
	Doctor
	Specialities   []Specialty `json:"specialities,omitempty"`
	Pending        bool        `json:"pending,omitempty"`
	PatientPending bool        `json:"patient_pending,omitempty"`
}

type NewDoctorResponseArgs struct {
	Doctor         db.Doctor
	Specialties    []db.Specialty
	PatientPending bool
	Pending        bool
}

func NewDoctorResponse(args NewDoctorResponseArgs) (DoctorsResponse, error) {
	doctorUUID, err := uuid.FromBytes(args.Doctor.ID.Bytes[:])
	if err != nil {
		return DoctorsResponse{}, err
	}

	institutionUUID, err := uuid.FromBytes(args.Doctor.InstitutionID.Bytes[:])
	if err != nil {
		return DoctorsResponse{}, err
	}

	// Fetch specialties
	specialties := []Specialty{}
	for _, s := range args.Specialties {
		specialtyUUID, err := uuid.FromBytes(s.ID.Bytes[:])
		if err != nil {
			return DoctorsResponse{}, err
		}
		specialties = append(specialties, Specialty{
			ID:          specialtyUUID,
			Description: s.Description,
			Name:        SpecialtyName(s.Name),
		})
	}

	return DoctorsResponse{
		UUID: doctorUUID,
		Doctor: Doctor{
			InstitutionUUID: institutionUUID,
			Firstname:       args.Doctor.Firstname,
			Lastname:        args.Doctor.Lastname,
			GovId:           args.Doctor.GovID,
			Birthdate:       args.Doctor.Birthdate.Time.Format("2006-01-02"),
			Email:           args.Doctor.Email,
			PhoneNumber:     args.Doctor.PhoneNumber,
			Credentials:     args.Doctor.Credentials,
		},
		Specialities:   specialties,
		Pending:        args.Doctor.Pending,
		PatientPending: args.Doctor.PatientPending,
	}, nil
}
