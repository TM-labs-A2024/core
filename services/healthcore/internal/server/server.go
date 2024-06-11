package server

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server will hold all dependencies for your application.
type Server struct {
	Router *echo.Echo
	Logger *slog.Logger
}

// NewServer returns an empty or an initialized container for your handlers.
func NewServer() (Server, error) {
	s := Server{
		Router: echo.New(),
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	return s, nil
}

func (s *Server) Start(port string) error {
	return s.Router.Start(port)
}

func (s *Server) AddRoutes() {

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	// DoctorsDoctorUuidDelete - Deletes a doctor
	s.Router.DELETE("/doctors/:doctorUuid", s.DoctorsDoctorUuidDelete)

	// DoctorsDoctorUuidGet - Returns a single doctor by UUID
	s.Router.GET("/doctors/:doctorUuid", s.DoctorsDoctorUuidGet)

	// DoctorsDoctorUuidPatientsGet - Returns a list of patients treated by doctor
	s.Router.GET("/doctors/:doctorUuid/patients", s.DoctorsDoctorUuidPatientsGet)

	// DoctorsGet - List ALL doctors
	s.Router.GET("/doctors", s.DoctorsGet)

	// DoctorsInstitutionUuidGet - List ALL doctors in an institution
	s.Router.GET("/doctors/:institutionUuid", s.DoctorsInstitutionUuidGet)

	// DoctorsLoginPost -
	s.Router.POST("/doctors/login", s.DoctorsLoginPost)

	// DoctorsPost - Add a new doctor to the system
	s.Router.POST("/doctors", s.DoctorsPost)

	// DoctorsPut - Update an existing doctor by Id
	s.Router.PUT("/doctors", s.DoctorsPut)

	// DoctorsSpecialtyIdGet - Returns a list of doctors by specialty
	s.Router.GET("/doctors/:specialtyId", s.DoctorsSpecialtyIdGet)

	// DoctorsSpecialtyIdPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
	s.Router.GET("/doctors/:specialtyId/patients", s.DoctorsSpecialtyIdPatientsGet)

	// GovermentLoginPost -
	s.Router.POST("/goverment/login", s.GovermentLoginPost)

	// GovernmentEnrollmentInstitutionUuidRevokePost - Deny institution into the system
	s.Router.POST("/government/enrollment/:institutionUuid/revoke", s.GovernmentEnrollmentInstitutionUuidRevokePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestUuidApprovePost - Approve institution into the system
	s.Router.POST("/government/enrollment-requests/:enrollmentRequestUuid/approve", s.GovernmentEnrollmentRequestsEnrollmentRequestUuidApprovePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestUuidDenyPost - Deny institution into the system
	s.Router.POST("/government/enrollment-requests/:enrollmentRequestUuid/deny", s.GovernmentEnrollmentRequestsEnrollmentRequestUuidDenyPost)

	// GovernmentEnrollmentRequestsGet - List request to approve institution into government
	s.Router.GET("/government/enrollment-requests", s.GovernmentEnrollmentRequestsGet)

	// GovernmentEnrollmentRequestsPost - Send request to approve institution into government
	s.Router.POST("/government/enrollment-requests", s.GovernmentEnrollmentRequestsPost)

	// InstitutionsEnrollmentDoctorUuidRevokePost - Deny doctor into institution
	s.Router.POST("/institutions/enrollment/:doctorUuid/revoke", s.InstitutionsEnrollmentDoctorUuidRevokePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestUuidApprovePost - Approve doctor into institution
	s.Router.POST("/institutions/enrollment-requests/:enrollmentRequestUuid/approve", s.InstitutionsEnrollmentRequestsEnrollmentRequestUuidApprovePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestUuidDenyPost - Deny doctor into institution
	s.Router.POST("/institutions/enrollment-requests/:enrollmentRequestUuid/deny", s.InstitutionsEnrollmentRequestsEnrollmentRequestUuidDenyPost)

	// InstitutionsEnrollmentRequestsGet - List request to approve doctor into institution
	s.Router.GET("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsGet)

	// InstitutionsEnrollmentRequestsPost - Send request to approve doctor into institution
	s.Router.POST("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsPost)

	// InstitutionsGet - List ALL institutions
	s.Router.GET("/institutions", s.InstitutionsGet)

	// InstitutionsGovIdGet - Returns a single institution by govId
	s.Router.GET("/institutions/:govId", s.InstitutionsGovIdGet)

	// InstitutionsInstitutionUuidDelete - Delete an institution
	s.Router.DELETE("/institutions/:institutionUuid", s.InstitutionsInstitutionUuidDelete)

	// InstitutionsInstitutionUuidUsersGet - list all institutions users on the system
	s.Router.GET("/institutions/:institutionUuid/users", s.InstitutionsInstitutionUuidUsersGet)

	// InstitutionsPost - Add a new institutions to the system
	s.Router.POST("/institutions", s.InstitutionsPost)

	// InstitutionsPut - Update an existing institutions by Id
	s.Router.PUT("/institutions", s.InstitutionsPut)

	// InstitutionsInstitutionUuidUsersGovIdGet - Returns a single institution user by gov id
	s.Router.GET("/institutions/:institutionUuid/users/:govId", s.InstitutionsInstitutionUuidUsersGovIdGet)

	// InstitutionsInstitutionUuidUsersLoginPost -
	s.Router.POST("/institutions/:institutionUuid/users/login", s.InstitutionsInstitutionUuidUsersLoginPost)

	// InstitutionsInstitutionUuidUsersPost - Add a new institutions user to the system
	s.Router.POST("/institutions/:institutionUuid/users", s.InstitutionsInstitutionUuidUsersPost)

	// InstitutionsInstitutionUuidUsersPut - Update an existing institutions user by Id
	s.Router.PUT("/institutions/:institutionUuid/users", s.InstitutionsInstitutionUuidUsersPut)

	// InstitutionsInstitutionUuidUsersUserUuidDelete - Deletes a institution user
	s.Router.DELETE("/institutions/:institutionUuid/users/:userUuid", s.InstitutionsInstitutionUuidUsersUserUuidDelete)

	// MedicalRecordHealthRecordUuidDelete - Deletes a medical-record on the DB ONLY
	s.Router.DELETE("/medical-record/:healthRecordUuid", s.MedicalRecordHealthRecordUuidDelete)

	// MedicalRecordHealthRecordUuidGet - Find medical-record by UUID
	s.Router.GET("/medical-record/:healthRecordUuid", s.MedicalRecordHealthRecordUuidGet)

	// MedicalRecordPost - Add a new medical-record to the system
	s.Router.POST("/medical-record", s.MedicalRecordPost)

	// NursesGet - List ALL nurses
	s.Router.GET("/nurses", s.NursesGet)

	// NursesInstitutionUuidGet - List ALL nurses in an institution
	s.Router.GET("/nurses/:institutionUuid", s.NursesInstitutionUuidGet)

	// NursesLoginPost -
	s.Router.POST("/nurses/login", s.NursesLoginPost)

	// NursesNurseUuidDelete - Deletes a nurse
	s.Router.DELETE("/nurses/:nurseUuid", s.NursesNurseUuidDelete)

	// NursesNurseUuidGet - Find nurse by UUID
	s.Router.GET("/nurses/:nurseUuid", s.NursesNurseUuidGet)

	// NursesPost - Add a new nurse to the system
	s.Router.POST("/nurses", s.NursesPost)

	// NursesPut - Update an existing nurse by UUID
	s.Router.PUT("/nurses", s.NursesPut)

	// PatientsAccessDoctorUuidRevokePost - Deny doctor access to patient records
	s.Router.POST("/patients/access/:doctorUuid/revoke", s.PatientsAccessDoctorUuidRevokePost)

	// PatientsAccessRequestsAccessRequestUuidApprovePost - Approve doctor access to patient records
	s.Router.POST("/patients/access-requests/:accessRequestUuid/approve", s.PatientsAccessRequestsAccessRequestUuidApprovePost)

	// PatientsAccessRequestsAccessRequestUuidDenyPost - Deny doctor access to patient records
	s.Router.POST("/patients/access-requests/:accessRequestUuid/deny", s.PatientsAccessRequestsAccessRequestUuidDenyPost)

	// PatientsAccessRequestsGet - List requests from doctors to access patient records
	s.Router.GET("/patients/access-requests", s.PatientsAccessRequestsGet)

	// PatientsGet - List ALL patients
	s.Router.GET("/patients", s.PatientsGet)

	// PatientsGovIdDoctorsGet - Returns a list of doctors treating patients
	s.Router.GET("/patients/:govId/doctors", s.PatientsGovIdDoctorsGet)

	// PatientsGovIdGet - Find patient by govId
	s.Router.GET("/patients/:govId", s.PatientsGovIdGet)

	// PatientsGovIdMedicalRecordsGet - List health records by patient
	s.Router.GET("/patients/:govId/medical-records", s.PatientsGovIdMedicalRecordsGet)

	// PatientsGovIdOrdersGet - List health orders by patient
	s.Router.GET("/patients/:govId/orders", s.PatientsGovIdOrdersGet)

	// PatientsLoginPost -
	s.Router.POST("/patients/login", s.PatientsLoginPost)

	// PatientsPatientUuidAccessRequestsPost - Make request for doctor to access patient records
	s.Router.POST("/patients/:patientUuid/access-requests", s.PatientsPatientUuidAccessRequestsPost)

	// PatientsPatientUuidDelete - Deletes a patient
	s.Router.DELETE("/patients/:patientUuid", s.PatientsPatientUuidDelete)

	// PatientsPost - Add a new patient to the system
	s.Router.POST("/patients", s.PatientsPost)

	// PatientsPut - Update an existing patient by uuid
	s.Router.PUT("/patients", s.PatientsPut)

	// SpecialtiesGet - Returns a list of specialties
	s.Router.GET("/specialties", s.SpecialtiesGet)

}
