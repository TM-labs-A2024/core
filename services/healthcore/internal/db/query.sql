
-- name: GetDoctorByUUID :one
SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials FROM Doctor WHERE uuid = $1;

-- INSERT INTO Doctor(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials) VALUES ($1, $2, $3, ?, ?, ?, ?, ?);

-- UPDATE Doctor SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ? WHERE 1;

-- DELETE FROM Doctor WHERE 0;


-- SELECT patientUuid, doctorUuid FROM DoctorAccessRequest WHERE 1;
-- INSERT INTO DoctorAccessRequest(patientUuid, doctorUuid) VALUES (?, ?);

-- UPDATE DoctorAccessRequest SET patientUuid = ?, doctorUuid = ? WHERE 1;

-- DELETE FROM DoctorAccessRequest WHERE 0;


-- SELECT institutionUuid, doctorUuid, pending, approved, professional-type FROM DoctorEnrollmentRequest WHERE 1;
-- INSERT INTO DoctorEnrollmentRequest(institutionUuid, doctorUuid, pending, approved, professional-type) VALUES (?, ?, ?, ?, ?);

-- UPDATE DoctorEnrollmentRequest SET institutionUuid = ?, doctorUuid = ?, pending = ?, approved = ?, professional-type = ? WHERE 1;

-- DELETE FROM DoctorEnrollmentRequest WHERE 0;


-- SELECT firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, uuid FROM _doctors__doctorUuid__patients_get_200_response_inner WHERE 1;
-- INSERT INTO _doctors__doctorUuid__patients_get_200_response_inner(firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _doctors__doctorUuid__patients_get_200_response_inner SET firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, sex = ?, pending = ?, status = ?, bed = ?, uuid = ? WHERE 1;

-- DELETE FROM _doctors__doctorUuid__patients_get_200_response_inner WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, specialities, pending, patientPending FROM _doctors_get_200_response_inner WHERE 1;
-- INSERT INTO _doctors_get_200_response_inner(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, specialities, pending, patientPending) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _doctors_get_200_response_inner SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, uuid = ?, specialities = ?, pending = ?, patientPending = ? WHERE 1;

-- DELETE FROM _doctors_get_200_response_inner WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, password, specialties FROM _doctors_post_request WHERE 1;
-- INSERT INTO _doctors_post_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, password, specialties) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _doctors_post_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, password = ?, specialties = ? WHERE 1;

-- DELETE FROM _doctors_post_request WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, password, specialties FROM _doctors_put_request WHERE 1;
-- INSERT INTO _doctors_put_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, password, specialties) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _doctors_put_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, uuid = ?, password = ?, specialties = ? WHERE 1;

-- DELETE FROM _doctors_put_request WHERE 0;


-- SELECT uuid, name FROM Government WHERE 1;
-- INSERT INTO Government(uuid, name) VALUES (?, ?);

-- UPDATE Government SET uuid = ?, name = ? WHERE 1;

-- DELETE FROM Government WHERE 0;


-- SELECT institutionUuid, pending, approved, uuid FROM _government_enrollment_requests_get_200_response_inner WHERE 1;
-- INSERT INTO _government_enrollment_requests_get_200_response_inner(institutionUuid, pending, approved, uuid) VALUES (?, ?, ?, ?);

-- UPDATE _government_enrollment_requests_get_200_response_inner SET institutionUuid = ?, pending = ?, approved = ?, uuid = ? WHERE 1;

-- DELETE FROM _government_enrollment_requests_get_200_response_inner WHERE 0;


-- SELECT content, type, specialty, content-format FROM HealthRecord WHERE 1;
-- INSERT INTO HealthRecord(content, type, specialty, content-format) VALUES (?, ?, ?, ?);

-- UPDATE HealthRecord SET content = ?, type = ?, specialty = ?, content-format = ? WHERE 1;

-- DELETE FROM HealthRecord WHERE 0;


-- SELECT name, govId FROM Institution WHERE 1;
-- INSERT INTO Institution(name, govId) VALUES (?, ?);

-- UPDATE Institution SET name = ?, govId = ? WHERE 1;

-- DELETE FROM Institution WHERE 0;


-- SELECT institutionUuid, pending, approved FROM InstitutionEnrollmentRequest WHERE 1;
-- INSERT INTO InstitutionEnrollmentRequest(institutionUuid, pending, approved) VALUES (?, ?, ?);

-- UPDATE InstitutionEnrollmentRequest SET institutionUuid = ?, pending = ?, approved = ? WHERE 1;

-- DELETE FROM InstitutionEnrollmentRequest WHERE 0;


-- SELECT institutionUuid, doctorUuid, pending, approved, professional-type, uuid FROM _institutions_enrollment_requests_get_200_response_inner WHERE 1;
-- INSERT INTO _institutions_enrollment_requests_get_200_response_inner(institutionUuid, doctorUuid, pending, approved, professional-type, uuid) VALUES (?, ?, ?, ?, ?, ?);

-- UPDATE _institutions_enrollment_requests_get_200_response_inner SET institutionUuid = ?, doctorUuid = ?, pending = ?, approved = ?, professional-type = ?, uuid = ? WHERE 1;

-- DELETE FROM _institutions_enrollment_requests_get_200_response_inner WHERE 0;


-- SELECT name, govId, uuid FROM _institutions_get_200_response_inner WHERE 1;
-- INSERT INTO _institutions_get_200_response_inner(name, govId, uuid) VALUES (?, ?, ?);

-- UPDATE _institutions_get_200_response_inner SET name = ?, govId = ?, uuid = ? WHERE 1;

-- DELETE FROM _institutions_get_200_response_inner WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, uuid FROM _institutions__institutionUuid__users_get_200_response_inner WHERE 1;
-- INSERT INTO _institutions__institutionUuid__users_get_200_response_inner(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, uuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _institutions__institutionUuid__users_get_200_response_inner SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, role = ?, uuid = ? WHERE 1;

-- DELETE FROM _institutions__institutionUuid__users_get_200_response_inner WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, password FROM _institutions__institutionUuid__users_post_request WHERE 1;
-- INSERT INTO _institutions__institutionUuid__users_post_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _institutions__institutionUuid__users_post_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, role = ?, password = ? WHERE 1;

-- DELETE FROM _institutions__institutionUuid__users_post_request WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, uuid, password FROM _institutions__institutionUuid__users_put_request WHERE 1;
-- INSERT INTO _institutions__institutionUuid__users_put_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role, uuid, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _institutions__institutionUuid__users_put_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, role = ?, uuid = ?, password = ? WHERE 1;

-- DELETE FROM _institutions__institutionUuid__users_put_request WHERE 0;


-- SELECT name, govId, uuid, pending FROM _institutions_put_200_response WHERE 1;
-- INSERT INTO _institutions_put_200_response(name, govId, uuid, pending) VALUES (?, ?, ?, ?);

-- UPDATE _institutions_put_200_response SET name = ?, govId = ?, uuid = ?, pending = ? WHERE 1;

-- DELETE FROM _institutions_put_200_response WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role FROM InstitutionUser WHERE 1;
-- INSERT INTO InstitutionUser(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE InstitutionUser SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, role = ? WHERE 1;

-- DELETE FROM InstitutionUser WHERE 0;


-- SELECT  FROM InstitutionUserRole WHERE 1;
-- INSERT INTO InstitutionUserRole() VALUES ();

-- UPDATE InstitutionUserRole SET  WHERE 1;

-- DELETE FROM InstitutionUserRole WHERE 0;


-- SELECT email, password FROM Login WHERE 1;
-- INSERT INTO Login(email, password) VALUES (?, ?);

-- UPDATE Login SET email = ?, password = ? WHERE 1;

-- DELETE FROM Login WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials FROM Nurse WHERE 1;
-- INSERT INTO Nurse(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE Nurse SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ? WHERE 1;

-- DELETE FROM Nurse WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, pending FROM _nurses_get_200_response_inner WHERE 1;
-- INSERT INTO _nurses_get_200_response_inner(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, pending) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _nurses_get_200_response_inner SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, uuid = ?, pending = ? WHERE 1;

-- DELETE FROM _nurses_get_200_response_inner WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, password FROM _nurses_post_request WHERE 1;
-- INSERT INTO _nurses_post_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _nurses_post_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, password = ? WHERE 1;

-- DELETE FROM _nurses_post_request WHERE 0;


-- SELECT institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, password FROM _nurses_put_request WHERE 1;
-- INSERT INTO _nurses_put_request(institutionUuid, firstname, lastname, govId, birthdate, email, phoneNumber, credentials, uuid, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _nurses_put_request SET institutionUuid = ?, firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, credentials = ?, uuid = ?, password = ? WHERE 1;

-- DELETE FROM _nurses_put_request WHERE 0;


-- SELECT firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed FROM Patient WHERE 1;
-- INSERT INTO Patient(firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE Patient SET firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, sex = ?, pending = ?, status = ?, bed = ? WHERE 1;

-- DELETE FROM Patient WHERE 0;


-- SELECT patientUuid, doctorUuid, uuid, completed FROM _patients_access_requests_get_200_response_inner WHERE 1;
-- INSERT INTO _patients_access_requests_get_200_response_inner(patientUuid, doctorUuid, uuid, completed) VALUES (?, ?, ?, ?);

-- UPDATE _patients_access_requests_get_200_response_inner SET patientUuid = ?, doctorUuid = ?, uuid = ?, completed = ? WHERE 1;

-- DELETE FROM _patients_access_requests_get_200_response_inner WHERE 0;


-- SELECT content, type, specialty, content-format, uuid FROM _patients__govId__medical_records_get_200_response_inner WHERE 1;
-- INSERT INTO _patients__govId__medical_records_get_200_response_inner(content, type, specialty, content-format, uuid) VALUES (?, ?, ?, ?, ?);

-- UPDATE _patients__govId__medical_records_get_200_response_inner SET content = ?, type = ?, specialty = ?, content-format = ?, uuid = ? WHERE 1;

-- DELETE FROM _patients__govId__medical_records_get_200_response_inner WHERE 0;


-- SELECT content, type, specialty, content-format FROM _patients__govId__orders_get_200_response_inner WHERE 1;
-- INSERT INTO _patients__govId__orders_get_200_response_inner(content, type, specialty, content-format) VALUES (?, ?, ?, ?);

-- UPDATE _patients__govId__orders_get_200_response_inner SET content = ?, type = ?, specialty = ?, content-format = ? WHERE 1;

-- DELETE FROM _patients__govId__orders_get_200_response_inner WHERE 0;


-- SELECT firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, password FROM _patients_post_request WHERE 1;
-- INSERT INTO _patients_post_request(firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _patients_post_request SET firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, sex = ?, pending = ?, status = ?, bed = ?, password = ? WHERE 1;

-- DELETE FROM _patients_post_request WHERE 0;


-- SELECT firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, uuid, password FROM _patients_put_request WHERE 1;
-- INSERT INTO _patients_put_request(firstname, lastname, govId, birthdate, email, phoneNumber, sex, pending, status, bed, uuid, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- UPDATE _patients_put_request SET firstname = ?, lastname = ?, govId = ?, birthdate = ?, email = ?, phoneNumber = ?, sex = ?, pending = ?, status = ?, bed = ?, uuid = ?, password = ? WHERE 1;

-- DELETE FROM _patients_put_request WHERE 0;


-- SELECT id, description, name FROM Specialty WHERE 1;
-- INSERT INTO Specialty(id, description, name) VALUES (?, ?, ?);

-- UPDATE Specialty SET id = ?, description = ?, name = ? WHERE 1;

-- DELETE FROM Specialty WHERE 0;


-- SELECT  FROM SpecialtyName WHERE 1;
-- INSERT INTO SpecialtyName() VALUES ();

-- UPDATE SpecialtyName SET  WHERE 1;

-- DELETE FROM SpecialtyName WHERE 0;

