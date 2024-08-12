package server

import (
	"log/slog"
	"os"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server will hold all dependencies for your application.
type Server struct {
	Router          *echo.Echo
	Controller      *controller.Controller
	Logger          *slog.Logger
	ivEncryptionKey string
}

// NewServer returns an empty or an initialized container for your handlers.
func NewServer(conf config.Config) (Server, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	c, err := controller.NewController(conf, logger)
	if err != nil {
		return Server{}, err
	}

	s := Server{
		Router:          echo.New(),
		Logger:          logger,
		Controller:      c,
		ivEncryptionKey: conf.IVEncryptionKey,
	}

	return s, nil
}

func (s Server) Start(port string) error {
	s.AddRoutes()
	return s.Router.Start(port)
}

func (s Server) AddRoutes() {

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.CORS())

	// Groups
	unrestricted := s.Router.Group("/api")
	restricted := s.Router.Group("/api")
	restricted.Use(controller.NewJWTMiddlewareFunc())

	// DoctorsDoctorIDDelete - Deletes a doctor
	restricted.DELETE("/doctors/:doctorId", s.DoctorsDoctorIDDelete)

	// DoctorsDoctorIDGet - Returns a single doctor by ID
	restricted.GET("/doctors/:doctorId", s.DoctorsDoctorIDGet)

	// DoctorsDoctorIDPatientsGet - Returns a list of patients treated by doctor
	restricted.GET("/doctors/patients/:doctorId", s.DoctorsDoctorIDPatientsGet)

	// DoctorsGet - List ALL doctors
	restricted.GET("/doctors", s.DoctorsGet)

	// DoctorsInstitutionIDGet - List ALL doctors in an institution
	restricted.GET("/doctors/institutions/:institutionId", s.DoctorsInstitutionIDGet)

	// DoctorsLoginPost -
	unrestricted.POST("/doctors/login", s.DoctorsLoginPost)

	// DoctorsPost - Add a new doctor to the system
	unrestricted.POST("/doctors", s.DoctorsPost)

	// DoctorsPut - Update an existing doctor by ID
	restricted.PUT("/doctors", s.DoctorsPut)

	// DoctorsSpecialtyIDGet - Returns a list of doctors by specialty
	restricted.GET("/doctors/specialties/:specialtyId", s.DoctorsSpecialtyIDGet)

	// DoctorsSpecialtyIDPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
	restricted.GET("/doctors/patients/specialties/:specialtyId", s.DoctorsSpecialtyIDPatientsGet)

	// GovermentLoginPost -
	unrestricted.POST("/goverment/login", s.GovermentLoginPost)

	// GovernmentEnrollmentInstitutionIDRevokePost - Deny institution into the system
	restricted.POST("/government/enrollment/:institutionId/revoke", s.GovernmentEnrollmentInstitutionIDRevokePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestIDApprovePost - Approve instituz	tion into the system
	restricted.POST("/government/enrollment-requests/:enrollmentRequestId/approve", s.GovernmentEnrollmentRequestsEnrollmentRequestIDApprovePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestIDDenyPost - Deny institution into the system
	restricted.POST("/government/enrollment-requests/:enrollmentRequestId/deny", s.GovernmentEnrollmentRequestsEnrollmentRequestIDDenyPost)

	// GovernmentEnrollmentRequestsGet - List request to approve institution into government
	restricted.GET("/government/enrollment-requests", s.GovernmentEnrollmentRequestsGet)

	// GovernmentEnrollmentRequestsPost - Send request to approve institution into government
	// restricted.POST("/government/enrollment-requests", s.GovernmentEnrollmentRequestsPost)

	// HealthRecordHealthReacordIDDelete - Deletes a health-record on the db ONLY
	restricted.DELETE("/health-record/:healthRecordId", s.HealthRecordHealthReacordIDDelete)

	// HealthRecordHealthReacordIDGet - Find health-record by ID
	restricted.GET("/health-record/:healthRecordId", s.HealthRecordHealthReacordIDGet)

	// HealthRecordPost - Add a new health-record to the system
	restricted.POST("/health-record", s.HealthRecordPost)

	// HealthRecordEvolutionPost - Add a new health-record to the system with only JSON data
	restricted.POST("/health-record/evolution", s.HealthRecordEvolutionPost)

	// HealthRecordDetach - Add a new health-record to the system with only JSON data
	restricted.POST("/health-record/:healthRecordId/detach", s.HealthRecordHealthReacordIDDetatchPost)

	// InstitutionsApprovedGet - List ALL approved institutions
	unrestricted.GET("/institutions/approved", s.InstitutionsApprovedGet)

	// InstitutionsEnrollmentDoctorIDRevokePost - Deny doctor into institution
	restricted.POST("/institutions/enrollment/:professionalId/revoke", s.InstitutionsEnrollmentDoctorIDRevokePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestIDApprovePost - Approve doctor into institution
	restricted.POST("/institutions/enrollment-requests/:enrollmentRequestId/approve", s.InstitutionsEnrollmentRequestsEnrollmentRequestIDApprovePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestIDDenyPost - Deny doctor into institution
	restricted.POST("/institutions/enrollment-requests/:enrollmentRequestId/deny", s.InstitutionsEnrollmentRequestsEnrollmentRequestIDDenyPost)

	// InstitutionsEnrollmentRequestsGet - List request to approve doctor into institution
	restricted.GET("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsGet)

	// InstitutionsEnrollmentRequestsPost - Send request to approve doctor into institution
	// restricted.POST("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsPost)

	// InstitutionsGet - List ALL institutions
	unrestricted.GET("/institutions", s.InstitutionsGet)

	// InstitutionsGovIDGet - Returns a single institution by id
	restricted.GET("/institutions/:institutionId", s.InstitutionsIDGet)

	// InstitutionsInstitutionIDDelete - Delete an institution
	restricted.DELETE("/institutions/:institutionId", s.InstitutionsInstitutionIDDelete)

	// InstitutionsInstitutionIDUsersGet - list all institutions users on the system
	restricted.GET("/institutions/:institutionId/users", s.InstitutionsInstitutionIDUsersGet)

	// InstitutionsPost - Add a new institutions to the system
	unrestricted.POST("/institutions", s.InstitutionsPost)

	// InstitutionsPut - Update an existing institutions by ID
	restricted.PUT("/institutions", s.InstitutionsPut)

	// InstitutionsInstitutionIDUsersGovIDGet - Returns a single institution user by gov id
	restricted.GET("/institutions/:institutionId/users/:govId", s.InstitutionsInstitutionIDUsersGovIDGet)

	// InstitutionsInstitutionIDUsersLoginPost -
	unrestricted.POST("/institutions/users/login", s.InstitutionsInstitutionIDUsersLoginPost)

	// InstitutionsInstitutionIDUsersPost - Add a new institutions user to the system
	unrestricted.POST("/institutions/users", s.InstitutionsInstitutionIDUsersPost)

	// InstitutionsInstitutionIDUsersPut - Update an existing institutions user by ID
	restricted.PUT("/institutions/users", s.InstitutionsInstitutionIDUsersPut)

	// InstitutionsInstitutionIDUsersUserIDDelete - Deletes a institution user
	restricted.DELETE("/institutions/:institutionId/users/:userId", s.InstitutionsInstitutionIDUsersUserIDDelete)

	// InstitutionPatientsGet - List all patientes hospitalized on institution the nurse belongs to
	restricted.GET("/institutions/patients", s.InstitutionsPatientsGet)

	// NursesGet - List ALL nurses
	restricted.GET("/nurses", s.NursesGet)

	// NursesInstitutionIDGet - List ALL nurses in an institution
	restricted.GET("/nurses/institutions/:institutionId", s.NursesInstitutionIDGet)

	// NursesLoginPost -
	unrestricted.POST("/nurses/login", s.NursesLoginPost)

	// NursesNurseIDDelete - Deletes a nurse
	restricted.DELETE("/nurses/:nurseId", s.NursesNurseIDDelete)

	// NursesNurseIDGet - Find nurse by ID
	restricted.GET("/nurses/:nurseId", s.NursesNurseIDGet)

	// NursesPost - Add a new nurse to the system
	unrestricted.POST("/nurses", s.NursesPost)

	// NursesPut - Update an existing nurse by ID
	restricted.PUT("/nurses", s.NursesPut)

	// PatientsAccessDoctorIDRevokePost - Deny doctor access to patient records
	restricted.POST("/patients/access/:doctorId/revoke", s.PatientsAccessDoctorIDRevokePost)

	// PatientsAccessRequestsAccessRequestIDApprovePost - Approve doctor access to patient records
	restricted.POST("/patients/access-requests/:accessRequestId/approve", s.PatientsAccessRequestsAccessRequestIDApprovePost)

	// PatientsAccessRequestsAccessRequestIDDenyPost - Deny doctor access to patient records
	restricted.POST("/patients/access-requests/:accessRequestId/deny", s.PatientsAccessRequestsAccessRequestIDDenyPost)

	// PatientsAccessRequestsGet - List requests from doctors to access patient records
	restricted.GET("/patients/access-requests", s.PatientsAccessRequestsGet)

	// PatientsGet - List ALL patients
	restricted.GET("/patients", s.PatientsGet)

	// PatientsGovIDDoctorsGet - Returns a list of doctors treating patients
	restricted.GET("/patients/:govId/doctors", s.PatientsGovIDDoctorsGet)

	// PatientsGovIDGet - Find patient by govID
	restricted.GET("/patients/:govId", s.PatientsGovIDGet)

	// PatientsGovIDHealthRecordsGet - List health records by patient
	restricted.GET("/patients/:govId/health-records", s.PatientsGovIDHealthRecordsGet)

	// PatientsGovIDHealthRecordsSpecialtiesGet - List health records by patient and specialty ID
	restricted.GET("/patients/:govId/health-records/specialties", s.PatientsGovIDHealthRecordsSpecialtiesGet)

	// PatientsGovIDHealthRecordsSpecialtyIDGet - List health records by patient and specialty ID
	restricted.GET("/patients/:govId/health-records/:specialtyId", s.PatientsGovIDHealthRecordsSpecialtyIDGet)

	// PatientsGovIDOrdersGet - List health orders by patient
	restricted.GET("/patients/:govId/orders", s.PatientsGovIDOrdersGet)

	// PatientsLoginPost -
	unrestricted.POST("/patients/login", s.PatientsLoginPost)

	// PatientsPatientIDAccessRequestsPost - Make request for doctor to access patient records
	restricted.POST("/patients/:patientId/access-requests", s.PatientsPatientIDAccessRequestsPost)

	// PatientsPatientIDDelete - Deletes a patient
	restricted.DELETE("/patients/:patientId", s.PatientsPatientIDDelete)

	// PatientsPost - Add a new patient to the system
	unrestricted.POST("/patients", s.PatientsPost)

	// PatientsPut - Update an existing patient by uuid
	restricted.PUT("/patients", s.PatientsPut)

	// SpecialtiesGet - Returns a list of specialties
	unrestricted.GET("/specialties", s.SpecialtiesGet)
}
