package server

import (
	"context"
	"log/slog"
	"os"

	"github.com/TM-labs-A2024/core/services/backend-server/internal/controller"
	"github.com/TM-labs-A2024/core/services/backend-server/internal/db"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server will hold all dependencies for your application.
type Server struct {
	Router *echo.Echo
	Logger *slog.Logger
	DB     *db.Queries
}

// NewServer returns an empty or an initialized container for your handlers.
func NewServer() (Server, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return Server{}, err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	s := Server{
		Router: echo.New(),
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		DB:     queries,
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

	// Groups
	unrestricted := s.Router.Group("/api")
	restricted := s.Router.Group("/api")
	restricted.Use(controller.NewJWTMiddlewareFunc())

	// DoctorsDoctorUUIDDelete - Deletes a doctor
	restricted.DELETE("/doctors/:doctor_uuid", s.DoctorsDoctorUUIDDelete)

	// DoctorsDoctorUUIDGet - Returns a single doctor by UUID
	restricted.GET("/doctors/:doctor_uuid", s.DoctorsDoctorUUIDGet)

	// DoctorsDoctorUUIDPatientsGet - Returns a list of patients treated by doctor
	restricted.GET("/doctors/:doctor_uuid/patients", s.DoctorsDoctorUUIDPatientsGet)

	// DoctorsGet - List ALL doctors
	restricted.GET("/doctors", s.DoctorsGet)

	// DoctorsInstitutionUUIDGet - List ALL doctors in an institution
	restricted.GET("/doctors/:institution_uuid", s.DoctorsInstitutionUUIDGet)

	// DoctorsLoginPost -
	unrestricted.POST("/doctors/login", s.DoctorsLoginPost)

	// DoctorsPost - Add a new doctor to the system
	unrestricted.POST("/doctors", s.DoctorsPost)

	// DoctorsPut - Update an existing doctor by Id
	restricted.PUT("/doctors", s.DoctorsPut)

	// DoctorsSpecialtyIdGet - Returns a list of doctors by specialty
	restricted.GET("/doctors/:specialtyId", s.DoctorsSpecialtyIdGet)

	// DoctorsSpecialtyIdPatientsGet - Returns a list of patients that have at least one record for a given  specialty that are treated by a doctor
	restricted.GET("/doctors/:specialtyId/patients", s.DoctorsSpecialtyIdPatientsGet)

	// GovermentLoginPost -
	unrestricted.POST("/goverment/login", s.GovermentLoginPost)

	// GovernmentEnrollmentInstitutionUUIDRevokePost - Deny institution into the system
	restricted.POST("/government/enrollment/:institution_uuid/revoke", s.GovernmentEnrollmentInstitutionUUIDRevokePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestUUIDApprovePost - Approve institution into the system
	restricted.POST("/government/enrollment-requests/:enrollmentRequestUUID/approve", s.GovernmentEnrollmentRequestsEnrollmentRequestUUIDApprovePost)

	// GovernmentEnrollmentRequestsEnrollmentRequestUUIDDenyPost - Deny institution into the system
	restricted.POST("/government/enrollment-requests/:enrollmentRequestUUID/deny", s.GovernmentEnrollmentRequestsEnrollmentRequestUUIDDenyPost)

	// GovernmentEnrollmentRequestsGet - List request to approve institution into government
	restricted.GET("/government/enrollment-requests", s.GovernmentEnrollmentRequestsGet)

	// GovernmentEnrollmentRequestsPost - Send request to approve institution into government
	restricted.POST("/government/enrollment-requests", s.GovernmentEnrollmentRequestsPost)

	// HealthRecordHealthReacordUUIDDelete - Deletes a health-record on the DB ONLY
	restricted.DELETE("/health-record/:healthReacordUUID", s.HealthRecordHealthReacordUUIDDelete)

	// HealthRecordHealthReacordUUIDGet - Find health-record by UUID
	restricted.GET("/health-record/:healthReacordUUID", s.HealthRecordHealthReacordUUIDGet)

	// HealthRecordPost - Add a new health-record to the system
	restricted.POST("/health-record", s.HealthRecordPost)

	// InstitutionsApprovedGet - List ALL approved institutions
	restricted.GET("/institutions/approved", s.InstitutionsApprovedGet)

	// InstitutionsEnrollmentDoctorUUIDRevokePost - Deny doctor into institution
	restricted.POST("/institutions/enrollment/:doctor_uuid/revoke", s.InstitutionsEnrollmentDoctorUUIDRevokePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestUUIDApprovePost - Approve doctor into institution
	restricted.POST("/institutions/enrollment-requests/:enrollmentRequestUUID/approve", s.InstitutionsEnrollmentRequestsEnrollmentRequestUUIDApprovePost)

	// InstitutionsEnrollmentRequestsEnrollmentRequestUUIDDenyPost - Deny doctor into institution
	restricted.POST("/institutions/enrollment-requests/:enrollmentRequestUUID/deny", s.InstitutionsEnrollmentRequestsEnrollmentRequestUUIDDenyPost)

	// InstitutionsEnrollmentRequestsGet - List request to approve doctor into institution
	restricted.GET("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsGet)

	// InstitutionsEnrollmentRequestsPost - Send request to approve doctor into institution
	restricted.POST("/institutions/enrollment-requests", s.InstitutionsEnrollmentRequestsPost)

	// InstitutionsGet - List ALL institutions
	restricted.GET("/institutions", s.InstitutionsGet)

	// InstitutionsGovIdGet - Returns a single institution by gov_id
	restricted.GET("/institutions/:gov_id", s.InstitutionsGovIdGet)

	// InstitutionsInstitutionUUIDDelete - Delete an institution
	restricted.DELETE("/institutions/:institution_uuid", s.InstitutionsInstitutionUUIDDelete)

	// InstitutionsInstitutionUUIDUsersGet - list all institutions users on the system
	restricted.GET("/institutions/:institution_uuid/users", s.InstitutionsInstitutionUUIDUsersGet)

	// InstitutionsPost - Add a new institutions to the system
	unrestricted.POST("/institutions", s.InstitutionsPost)

	// InstitutionsPut - Update an existing institutions by Id
	restricted.PUT("/institutions", s.InstitutionsPut)

	// InstitutionsInstitutionUUIDUsersGovIdGet - Returns a single institution user by gov id
	restricted.GET("/institutions/:institution_uuid/users/:gov_id", s.InstitutionsInstitutionUUIDUsersGovIdGet)

	// InstitutionsInstitutionUUIDUsersLoginPost -
	unrestricted.POST("/institutions/:institution_uuid/users/login", s.InstitutionsInstitutionUUIDUsersLoginPost)

	// InstitutionsInstitutionUUIDUsersPost - Add a new institutions user to the system
	unrestricted.POST("/institutions/:institution_uuid/users", s.InstitutionsInstitutionUUIDUsersPost)

	// InstitutionsInstitutionUUIDUsersPut - Update an existing institutions user by Id
	restricted.PUT("/institutions/:institution_uuid/users", s.InstitutionsInstitutionUUIDUsersPut)

	// InstitutionsInstitutionUUIDUsersUserUUIDDelete - Deletes a institution user
	restricted.DELETE("/institutions/:institution_uuid/users/:userUUID", s.InstitutionsInstitutionUUIDUsersUserUUIDDelete)

	// NursesGet - List ALL nurses
	restricted.GET("/nurses", s.NursesGet)

	// NursesInstitutionUUIDGet - List ALL nurses in an institution
	restricted.GET("/nurses/:institution_uuid", s.NursesInstitutionUUIDGet)

	// NursesLoginPost -
	unrestricted.POST("/nurses/login", s.NursesLoginPost)

	// NursesNurseUUIDDelete - Deletes a nurse
	restricted.DELETE("/nurses/:nurseUUID", s.NursesNurseUUIDDelete)

	// NursesNurseUUIDGet - Find nurse by UUID
	restricted.GET("/nurses/:nurseUUID", s.NursesNurseUUIDGet)

	// NursesPost - Add a new nurse to the system
	unrestricted.POST("/nurses", s.NursesPost)

	// NursesPut - Update an existing nurse by UUID
	restricted.PUT("/nurses", s.NursesPut)

	// PatientsAccessDoctorUUIDRevokePost - Deny doctor access to patient records
	restricted.POST("/patients/access/:doctor_uuid/revoke", s.PatientsAccessDoctorUUIDRevokePost)

	// PatientsAccessRequestsAccessRequestUUIDApprovePost - Approve doctor access to patient records
	restricted.POST("/patients/access-requests/:accessRequestUUID/approve", s.PatientsAccessRequestsAccessRequestUUIDApprovePost)

	// PatientsAccessRequestsAccessRequestUUIDDenyPost - Deny doctor access to patient records
	restricted.POST("/patients/access-requests/:accessRequestUUID/deny", s.PatientsAccessRequestsAccessRequestUUIDDenyPost)

	// PatientsAccessRequestsGet - List requests from doctors to access patient records
	restricted.GET("/patients/access-requests", s.PatientsAccessRequestsGet)

	// PatientsGet - List ALL patients
	restricted.GET("/patients", s.PatientsGet)

	// PatientsGovIdDoctorsGet - Returns a list of doctors treating patients
	restricted.GET("/patients/:gov_id/doctors", s.PatientsGovIdDoctorsGet)

	// PatientsGovIdGet - Find patient by gov_id
	restricted.GET("/patients/:gov_id", s.PatientsGovIdGet)

	// PatientsGovIdHealthRecordsGet - List health records by patient
	restricted.GET("/patients/:gov_id/health-records", s.PatientsGovIdHealthRecordsGet)

	// PatientsGovIdHealthRecordsSpecialtiesGet - List health records by patient and specialty Id
	restricted.GET("/patients/:gov_id/health-records/specialties", s.PatientsGovIdHealthRecordsSpecialtiesGet)

	// PatientsGovIdHealthRecordsSpecialtyIdGet - List health records by patient and specialty Id
	restricted.GET("/patients/:gov_id/health-records/:specialtyId", s.PatientsGovIdHealthRecordsSpecialtyIdGet)

	// PatientsGovIdOrdersGet - List health orders by patient
	restricted.GET("/patients/:gov_id/orders", s.PatientsGovIdOrdersGet)

	// PatientsLoginPost -
	unrestricted.POST("/patients/login", s.PatientsLoginPost)

	// PatientsPatientUUIDAccessRequestsPost - Make request for doctor to access patient records
	restricted.POST("/patients/:patient_uuid/access-requests", s.PatientsPatientUUIDAccessRequestsPost)

	// PatientsPatientUUIDDelete - Deletes a patient
	restricted.DELETE("/patients/:patient_uuid", s.PatientsPatientUUIDDelete)

	// PatientsPost - Add a new patient to the system
	unrestricted.POST("/patients", s.PatientsPost)

	// PatientsPut - Update an existing patient by uuid
	restricted.PUT("/patients", s.PatientsPut)

	// SpecialtiesGet - Returns a list of specialties
	restricted.GET("/specialties", s.SpecialtiesGet)
}
