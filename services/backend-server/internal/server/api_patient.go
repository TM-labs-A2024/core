package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// PatientsAccessDoctorIDRevokePost - Deny doctor access to patient records
func (s *Server) PatientsAccessDoctorIDRevokePost(ctx echo.Context) error {
	uuidStr := ctx.Param("doctorId")
	doctorID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	ar, err := s.Controller.GetAccessRequestByPatientAndDoctorID(doctorID, claims.UserID)
	if err != nil {
		return err
	}

	if err := s.Controller.DeleteAccessRequestsByID(ar.ID.Bytes); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// PatientsAccessRequestsAccessRequestIDApprovePost - Approve doctor access to patient records
func (s *Server) PatientsAccessRequestsAccessRequestIDApprovePost(ctx echo.Context) error {
	uuidStr := ctx.Param("accessRequestId")
	accessRequestID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ar, err := s.Controller.GetAccessRequestByID(accessRequestID)
	if err != nil {
		return err
	}

	ar, err = s.Controller.ApproveAccessRequestsByID(ar)
	if err != nil {
		return err
	}

	resp, err := models.NewDoctorAccessResponse(ar)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsAccessRequestsAccessRequestIDDenyPost - Deny doctor access to patient records
func (s *Server) PatientsAccessRequestsAccessRequestIDDenyPost(ctx echo.Context) error {
	uuidStr := ctx.Param("accessRequestId")
	accessRequestID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ar, err := s.Controller.GetAccessRequestByID(accessRequestID)
	if err != nil {
		return err
	}

	ar, err = s.Controller.DenyAccessRequestsByID(ar)
	if err != nil {
		return err
	}

	resp, err := models.NewDoctorAccessResponse(ar)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsAccessRequestsGet - List requests from doctors to access patient records
func (s *Server) PatientsAccessRequestsGet(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	ars, err := s.Controller.ListAccessRequestsByPatientID(claims.UserID)
	if err != nil {
		return err
	}

	resps := []models.DoctorAccessResponse{}
	for _, ar := range ars {
		resp, err := models.NewDoctorAccessResponse(ar)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGet - List ALL patients
func (s *Server) PatientsGet(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)
	doctorID, err := uuid.Parse(claims.ID)
	if err != nil {
		return err
	}

	patients, err := s.Controller.ListPatients(doctorID)
	if err != nil {
		return err
	}

	resps := []models.PatientResponse{}
	for _, patient := range patients {
		resp, err := models.NewPatientResponse(patient)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIDDoctorsGet - Returns a list of doctors treating patients
func (s *Server) PatientsGovIDDoctorsGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	doctors, err := s.Controller.ListPatientApprovedDoctorsByGovID(govID)
	if err != nil {
		return err
	}

	resps := []models.DoctorsResponse{}
	for _, doctor := range doctors {
		specialties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
			Doctor:      doctor,
			Specialties: specialties,
		})
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIDGet - Find patient by gov_id
func (s *Server) PatientsGovIDGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	patient, err := s.Controller.GetPatientByGovID(govID)
	if err != nil {
		return err
	}

	resp, err := models.NewPatientResponse(patient)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsGovIDHealthRecordsGet - List health records by patient
func (s *Server) PatientsGovIDHealthRecordsGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	healthRecords, err := s.Controller.ListHealthRecordPatientsGovID(govID)
	if err != nil {
		return err
	}

	patient, err := s.Controller.GetPatientByGovID(govID)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewHealthRecordResponse(db.CreateHealthRecordResult{
			HealthRecord: healthRecord,
			Specialty:    specialty,
			Patient:      patient,
		}, s.ivEncryptionKey)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIDHealthRecordsSpecialtiesGet - List health records specialty ids by patient id
func (s *Server) PatientsGovIDHealthRecordsSpecialtiesGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	healthRecords, err := s.Controller.ListHealthRecordPatientsGovID(govID)
	if err != nil {
		return err
	}

	specialties := []db.Specialty{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		specialties = append(specialties, specialty)
	}

	resps := models.NewSpecialtiesResponse(specialties)

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIDHealthRecordsSpecialtyIDGet - List health records by patient and specialty ID
func (s *Server) PatientsGovIDHealthRecordsSpecialtyIDGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	uuidStr := ctx.Param("specialtyId")
	specialtyID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	healthRecords, err := s.Controller.ListHealthRecordByPatientsGovAndSpecialtyID(
		govID,
		specialtyID,
	)
	if err != nil {
		return err
	}

	patient, err := s.Controller.GetPatientByGovID(govID)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(specialtyID)
		if err != nil {
			return err
		}

		resp, err := models.NewHealthRecordResponse(db.CreateHealthRecordResult{
			HealthRecord: healthRecord,
			Specialty:    specialty,
			Patient:      patient,
		}, s.ivEncryptionKey)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIDOrdersGet - List health orders by patient
func (s *Server) PatientsGovIDOrdersGet(ctx echo.Context) error {
	govID := ctx.Param("govId")

	healthRecords, err := s.Controller.ListOrdersByPatientGovID(govID)
	if err != nil {
		return err
	}

	patient, err := s.Controller.GetPatientByGovID(govID)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewHealthRecordResponse(db.CreateHealthRecordResult{
			HealthRecord: healthRecord,
			Specialty:    specialty,
			Patient:      patient,
		}, s.ivEncryptionKey)
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsLoginPost -
func (s *Server) PatientsLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	user, err := s.Controller.GetPatientByLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	token, err := controller.NewClaim(user.ID.Bytes, user.InstitutionID.Bytes)
	if err != nil {
		return err
	}

	resp, err := models.NewPatientResponse(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token":   token,
		"patient": resp,
	})
}

// PatientsPatientIDAccessRequestsPost - Make request for doctor to access patient records
func (s *Server) PatientsPatientIDAccessRequestsPost(ctx echo.Context) error {
	uuidStr := ctx.Param("patientId")
	patientID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	ar, err := s.Controller.CreateAccessRequestWithDoctorAndPatientID(claims.UserID, patientID)
	if err != nil {
		return err
	}

	resp, err := models.NewDoctorAccessResponse(ar)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsPatientIDDelete - Deletes a patient
func (s *Server) PatientsPatientIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("patientId")
	patientID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeletePatientByID(patientID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// PatientsPost - Add a new patient to the system
func (s *Server) PatientsPost(ctx echo.Context) error {
	req := models.PatientPostRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	patient, err := s.Controller.CreatePatient(req)
	if err != nil {
		return err
	}

	resp, err := models.NewPatientResponse(patient)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsPut - Update an existing patient by uuid
func (s *Server) PatientsPut(ctx echo.Context) error {
	req := models.PatientPutRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	patient, err := s.Controller.UpdatePatientByID(req)
	if err != nil {
		return err
	}

	resp, err := models.NewPatientResponse(patient)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}
