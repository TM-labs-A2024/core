package server

import (
	"net/http"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

// PatientsAccessDoctorIDRevokePost - Deny doctor access to patient records
func (s *Server) PatientsAccessDoctorIDRevokePost(ctx echo.Context) error {
	uuidStr := ctx.Param("doctorId")
	doctorId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	ar, err := s.Controller.GetAccessRequestByPatientAndDoctorID(doctorId, claims.UserID)
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
	accessRequestId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ar, err := s.Controller.GetAccessRequestByID(accessRequestId)
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
	accessRequestId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	ar, err := s.Controller.GetAccessRequestByID(accessRequestId)
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

	ars, err := s.Controller.ListAccessRequestsByPatientId(claims.UserID)
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
	patients, err := s.Controller.ListPatients()
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

// PatientsGovIdDoctorsGet - Returns a list of doctors treating patients
func (s *Server) PatientsGovIdDoctorsGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	doctors, err := s.Controller.ListPatientApprovedDoctors(govId)
	if err != nil {
		return err
	}

	resps := []models.DoctorsResponse{}
	for _, doctor := range doctors {
		specialties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
		if err != nil {
			return err
		}

		accessRequests, err := s.Controller.ListAccessRequestsByDoctorID(doctor.ID.Bytes)
		if err != nil && err != pgx.ErrNoRows {
			return err
		}

		enrollmentRequests, err := s.Controller.GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID(
			doctor.ID.Bytes,
			doctor.InstitutionID.Bytes,
		)
		if err != nil && err != pgx.ErrNoRows {
			return err
		}

		resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
			Doctor:         doctor,
			Specialties:    specialties,
			PatientPending: len(accessRequests) != 0,
			Pending:        enrollmentRequests.ID.Valid,
		})
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIdGet - Find patient by gov_id
func (s *Server) PatientsGovIdGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	patient, err := s.Controller.GetPatientByGovID(govId)
	if err != nil {
		return err
	}

	resp, err := models.NewPatientResponse(patient)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// PatientsGovIdHealthRecordsGet - List health records by patient
func (s *Server) PatientsGovIdHealthRecordsGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	healthRecords, err := s.Controller.ListHealthRecordPatientsGovID(govId)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		resp := models.NewHealthRecordResponse(healthRecord, specialty)

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIdHealthRecordsSpecialtiesGet - List health records by patient and specialty Id
func (s *Server) PatientsGovIdHealthRecordsSpecialtiesGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	healthRecords, err := s.Controller.ListHealthRecordPatientsGovID(govId)
	if err != nil {
		return err
	}

	resps := map[uuid.UUID][]models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		specialty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		resp := models.NewHealthRecordResponse(healthRecord, specialty)

		resps[specialty.ID.Bytes] = append(resps[specialty.ID.Bytes], resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIdHealthRecordsSpecialtyIdGet - List health records by patient and specialty Id
func (s *Server) PatientsGovIdHealthRecordsSpecialtyIdGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	uuidStr := ctx.Param("specialtyId")
	specialtyId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	healthRecords, err := s.Controller.ListHealthRecordByPatientsGovAndSpecialtyID(
		govId,
		specialtyId,
	)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		speciaty, err := s.Controller.GetSpecialtyByID(specialtyId)
		if err != nil {
			return err
		}

		resp := models.NewHealthRecordResponse(healthRecord, speciaty)

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// PatientsGovIdOrdersGet - List health orders by patient
func (s *Server) PatientsGovIdOrdersGet(ctx echo.Context) error {
	govId := ctx.Param("govId")

	healthRecords, err := s.Controller.ListOrdersByPatientGovID(govId)
	if err != nil {
		return err
	}

	resps := []models.HealthRecordResponse{}
	for _, healthRecord := range healthRecords {
		speciaty, err := s.Controller.GetSpecialtyByID(healthRecord.SpecialtyID.Bytes)
		if err != nil {
			return err
		}

		resp := models.NewHealthRecordResponse(healthRecord, speciaty)

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

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// PatientsPatientIDAccessRequestsPost - Make request for doctor to access patient records
func (s *Server) PatientsPatientIDAccessRequestsPost(ctx echo.Context) error {
	uuidStr := ctx.Param("patientId")
	patientId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	ar, err := s.Controller.CreateAccessRequestWithDoctorAndPatientID(claims.UserID, patientId)
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
	patientId, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeletePatientByID(patientId); err != nil {
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
