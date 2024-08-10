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

// DoctorsDoctorUUIDDelete - Deletes a doctor
func (s *Server) DoctorsDoctorIDDelete(ctx echo.Context) error {
	uuidStr := ctx.Param("doctorId")
	doctorID, err := uuid.Parse((uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := s.Controller.DeleteDoctorByID(doctorID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DoctorsDoctorUUIDGet - Returns a single doctor by UUID
func (s *Server) DoctorsDoctorIDGet(ctx echo.Context) error {
	uuidStr := ctx.Param("doctorId")
	doctorID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	doctor, err := s.Controller.GetDoctorByID(doctorID)
	if err != nil {
		return err
	}

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

	return ctx.JSON(http.StatusOK, resp)
}

// DoctorsDoctorUUIDPatientsGet - Returns a list of patients treated by doctor
func (s *Server) DoctorsDoctorIDPatientsGet(ctx echo.Context) error {
	uuidStr := ctx.Param("doctorId")
	doctorID, err := uuid.Parse((uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	patients, err := s.Controller.ListPatientsTreatedByDoctorID(doctorID)
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

// DoctorsGet - List ALL doctors
func (s *Server) DoctorsGet(ctx echo.Context) error {
	resps := []models.DoctorsResponse{}
	doctors, err := s.Controller.ListDoctors()
	if err != nil {
		return err
	}
	for _, doctor := range doctors {
		speciaties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
			Doctor:      doctor,
			Specialties: speciaties,
		})
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// DoctorsInstitutionUUIDGet - List ALL doctors in an institution
func (s *Server) DoctorsInstitutionIDGet(ctx echo.Context) error {
	resps := []models.DoctorsResponse{}
	uuidStr := ctx.Param("institutionId")
	institutionID, err := uuid.Parse((uuidStr))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	doctors, err := s.Controller.ListDoctorsByInstitutionID(institutionID)
	if err != nil {
		return err
	}

	for _, doctor := range doctors {
		speciaties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
			Doctor:      doctor,
			Specialties: speciaties,
		})
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// DoctorsLoginPost -
func (s *Server) DoctorsLoginPost(ctx echo.Context) error {
	request := models.Login{}
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	doctor, err := s.Controller.GetDoctorByLogin(request.Email, request.Password)
	if err != nil {
		return err
	}

	speciaties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
	if err != nil {
		return err
	}

	token, err := controller.NewClaim(doctor.ID.Bytes, doctor.InstitutionID.Bytes)
	if err != nil {
		return err
	}

	resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
		Doctor:      doctor,
		Specialties: speciaties,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token":  token,
		"doctor": resp,
	})
}

// DoctorsPost - Add a new doctor to the system
func (s *Server) DoctorsPost(ctx echo.Context) error {
	req := models.DoctorsPostRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	doctor, err := s.Controller.CreateDoctor(req)
	if err != nil {
		return err
	}

	specialties := []db.Specialty{}
	for _, specialty := range req.Specialties {
		specialty, err := s.Controller.GetSpecialtyByID(specialty)
		if err != nil {
			return err
		}

		specialties = append(specialties, specialty)
	}

	resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
		Doctor:      doctor,
		Specialties: specialties,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

// DoctorsPut - Update an existing doctor by ID
func (s *Server) DoctorsPut(ctx echo.Context) error {
	req := models.DoctorsPutRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	doctor, err := s.Controller.UpdateDoctorByID(req)
	if err != nil {
		return err
	}

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

	return ctx.JSON(http.StatusOK, resp)
}

// DoctorsSpecialtyIDGet - Returns a list of doctors by specialty
func (s *Server) DoctorsSpecialtyIDGet(ctx echo.Context) error {
	resps := []models.DoctorsResponse{}
	uuidStr := ctx.Param("specialtyId")
	specialtyID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	specialtyDoctorJunctions, err := s.Controller.ListDoctorsBySpecialtyID(specialtyID)
	if err != nil {
		return err
	}

	for _, specialtyDoctorJunction := range specialtyDoctorJunctions {
		doctor, err := s.Controller.GetDoctorByID(specialtyDoctorJunction.DoctorID.Bytes)
		if err != nil {
			return err
		}

		speciaties, err := s.Controller.ListSpecialtiesByDoctorID(doctor.ID.Bytes)
		if err != nil {
			return err
		}

		resp, err := models.NewDoctorResponse(models.NewDoctorResponseArgs{
			Doctor:      doctor,
			Specialties: speciaties,
		})
		if err != nil {
			return err
		}

		resps = append(resps, resp)
	}

	return ctx.JSON(http.StatusOK, resps)
}

// DoctorsSpecialtyIDPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
func (s *Server) DoctorsSpecialtyIDPatientsGet(ctx echo.Context) error {
	uuidStr := ctx.Param("specialtyId")
	specialtyID, err := uuid.Parse(uuidStr)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*controller.JWTCustomClaims)

	patients, err := s.Controller.ListPatientsTreatedByDoctorIDWithHealthRecordOfSpecialtyID(claims.UserID, specialtyID)
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
