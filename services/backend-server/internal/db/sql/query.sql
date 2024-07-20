-- name: GetDoctorByID :one
SELECT *
FROM doctor
WHERE id = $1;
-- name: GetDoctorByLogin :one
SELECT *
FROM doctor
WHERE email = $1 AND password = crypt($2, password);
-- name: ListDoctors :many
SELECT *
FROM doctor;
-- name: ListDoctorsByInstitutionID :many
SELECT *
FROM doctor
WHERE institution_id = $1;
-- name: CreateDoctor :one
INSERT INTO doctor(
        institution_id,
        firstname,
        lastname,
        gov_id,
        birthdate,
        password,
        email,
        phone_number,
        credentials,
        sex
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        crypt($6, gen_salt('bf')),
        $7,
        $8,
        $9,
        $10
    )
RETURNING *;
-- name: UpdateDoctorByID :one
UPDATE doctor
SET institution_id = $1,
    firstname = $2,
    lastname = $3,
    gov_id = $4,
    birthdate = $5,
    password = crypt($6, gen_salt('bf')),
    email = $7,
    phone_number = $8,
    credentials = $9,
    pending = $10,
    patient_pending = $11,
    sex = $12
WHERE id = $13
RETURNING *;
-- name: DeleteDoctorByID :exec
DELETE FROM doctor
WHERE id = $1;
-- name: GetAccessRequestsByID :one
SELECT *
FROM doctor_access_request
WHERE id = $1;
-- name: GetAccessRequestsByPatientAndDoctorID :one
SELECT *
FROM doctor_access_request
WHERE patient_id = $1 AND doctor_id = $2;
-- name: ListAccessRequestsByPatientID :many
SELECT *
FROM doctor_access_request
WHERE patient_id = $1;
-- name: ListApprovedAccessRequestsByPatientID :many
SELECT *
FROM doctor_access_request
WHERE patient_id = $1 AND approved = TRUE;
-- name: ListAccessRequestsByDoctorID :many
SELECT *
FROM doctor_access_request
WHERE doctor_id = $1;
-- name: CreateAccessRequest :one
INSERT INTO doctor_access_request(patient_id, doctor_id)
VALUES ($1, $2)
RETURNING *;
-- name: UpdateAccessRequestByID :one
UPDATE doctor_access_request
SET patient_id = $1,
    doctor_id = $2,
    pending = $3,
    approved = $4
WHERE id = $5
RETURNING *;
-- name: DeleteAccessRequestByID :exec
DELETE FROM doctor_access_request
WHERE id = $1;
-- name: ListInstitutionEnrollmentRequestsByInstitutionID :many
SELECT *
FROM institution_enrollment_request
WHERE institution_id = $1;
-- name: GetInstitutionEnrollmentRequestsByID :one
SELECT *
FROM institution_enrollment_request
WHERE id = $1;
-- name: ListInstitutionEnrollmentRequestByDoctorID :many
SELECT *
FROM institution_enrollment_request
WHERE doctor_id = $1;
-- name: GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID :one
SELECT *
FROM institution_enrollment_request
WHERE doctor_id = $1 AND institution_id = $2;
-- name: ListInstitutionEnrollmentRequestByNurseID :many
SELECT *
FROM institution_enrollment_request
WHERE nurse_id = $1;
-- name: CreateInstitutionEnrollmentRequest :one
INSERT INTO institution_enrollment_request(
        institution_id,
        doctor_id,
        nurse_id
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetNurseByLogin :one
SELECT *
FROM nurse
WHERE email = $1 AND password = crypt($2, password);
-- name: UpdateInstitutionEnrollmentRequestByID :one
UPDATE institution_enrollment_request
SET institution_id = $1,
    doctor_id = $2,
    nurse_id = $3,
    pending = $4,
    approved = $5
WHERE id = $6
RETURNING *;
-- name: DeleteInstitutionEnrollmentRequestByID :exec
DELETE FROM institution_enrollment_request
WHERE id = $1;
-- name: GetGovernmentByLogin :one
SELECT *
FROM government
WHERE email = $1 AND password = crypt($2, password);
-- name: CreateGovernment :one
INSERT INTO government(email, password)
VALUES ($1, $2)
RETURNING *;
-- name: UpdateGovernmentByID :one
UPDATE government
SET email = $1,
    password = $2
WHERE id = $3
RETURNING *;
-- name: DeleteGovernmentByID :exec
DELETE FROM government
WHERE id = $1;
-- name: ListHealthRecordsBySpecialtyID :many
SELECT *
FROM health_record
WHERE specialty_id = $1;
-- name: ListHealthRecordsByPatientID :many
SELECT *
FROM health_record
WHERE patient_id = $1;
-- name: ListHealthRecordsBySpecialtyAndPatientID :many
SELECT *
FROM health_record
WHERE specialty_id = $1
    AND patient_id = $2;
-- name: ListHealthRecordsByTypeAndPatientID :many
SELECT *
FROM health_record
WHERE type = $1
    AND patient_id = $2;
-- name: CreateHealthRecord :one
INSERT INTO health_record (
        patient_id,
        private_key,
        public_key,
        type,
        specialty_id,
        content_format
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: DeleteHealthRecordByID :exec
DELETE FROM health_record
WHERE id = $1;
-- name: GetHealthRecordByID :one
SELECT * 
FROM health_record
WHERE id = $1;
-- name: GetInstitutionByGovID :one
SELECT *
FROM institution
WHERE gov_id = $1;
-- name: ListInstitutions :many
SELECT *
FROM institution;
-- name: ListApprovedInstitutions :many
SELECT *
FROM institution
WHERE id IN (
        SELECT institution_id
        FROM government_enrollment_request
        WHERE approved = TRUE
    );
-- name: CreateInstitution :one
INSERT INTO institution (
        name,
        address,
        credentials,
        type,
        gov_id
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateInstitutionByID :one
UPDATE institution
SET name = $1,
    address = $2,
    credentials = $3,
    type = $4,
    gov_id = $5,
    pending = $6
WHERE id = $7
RETURNING *;
-- name: DeleteInstitutionByID :exec
DELETE FROM institution
WHERE id = $1;
-- name: ListGovernmentEnrollmentRequests :many
SELECT *
FROM government_enrollment_request;
-- name: CreateGovernmentEnrollmentRequests :one
INSERT INTO government_enrollment_request (institution_id)
VALUES ($1)
RETURNING *;
-- name: UpdateGovernmentEnrollmentRequestsByID :one
UPDATE government_enrollment_request
SET institution_id = $1,
    pending = $2,
    approved = $3
WHERE id = $4
RETURNING *;
-- name: DeleteGovernmentEnrollmentRequestByID :exec
DELETE FROM government_enrollment_request
WHERE id = $1;
-- name: GetGovernmentEnrollmentRequestByID :one
SELECT *
FROM government_enrollment_request
WHERE id = $1;
-- name: GetInstitutionUserByLogin :one
SELECT *
FROM institution_user
WHERE email = $1 AND password = crypt($2, password);
-- name: GetInstitutionUserByID :many
SELECT *
FROM institution_user
WHERE id = $1;
-- name: GetInstitutionUserByGovAndInstitutionID :one
SELECT *
FROM institution_user
WHERE gov_id = $1 AND institution_id = $2;
-- name: GetFirstInstitutionUserByInstitutionID :one
SELECT *
FROM institution_user
WHERE institution_id = $1
ORDER BY created_at
LIMIT 1;
-- name: ListInstitutionUserByInstitutionID :many
SELECT *
FROM institution_user
WHERE institution_id = $1;
-- name: CreateInstitutionUser :one
INSERT INTO institution_user(
        institution_id,
        firstname,
        lastname,
        gov_id,
        birthdate,
        email,
        password,
        phone_number,
        role
    )
VALUES ($1, $2, $3, $4, $5, $6, crypt($7, gen_salt('bf')), $8, $9)
RETURNING *;
-- name: UpdateInstitutionUserByGovID :one
UPDATE institution_user
SET firstname = $2,
    lastname = $3,
    birthdate = $5,
    email = $6,
    password = crypt($7, gen_salt('bf')),
    phone_number = $8,
    role = $9
WHERE gov_id = $4 AND institution_id = $1
RETURNING *;
-- name: DeleteInstitutionUserByInsitutionAndUserID :exec
DELETE FROM institution_user
WHERE id = $1 AND institution_id = $2;
-- SELECT  FROM InstitutionUserRole WHERE 1;
-- INSERT INTO InstitutionUserRole() VALUES ();
-- UPDATE InstitutionUserRole SET  WHERE 1;
-- DELETE FROM institution_user_role WHERE id = $1;
-- name: GetNurseByID :one
SELECT *
FROM nurse
WHERE id = $1;
-- name: ListNurses :many
SELECT *
FROM nurse;
-- name: ListNursesByInstitutionID :many
SELECT *
FROM nurse
WHERE institution_id = $1;
-- name: CreateNurse :one
INSERT INTO nurse(
        institution_id,
        firstname,
        lastname,
        gov_id,
        birthdate,
        email,
        phone_number,
        credentials,
        password,
        pending,
        sex
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, crypt($9, gen_salt('bf')), $10, $11)
RETURNING *;
-- name: UpdateNurseByID :one
UPDATE nurse
SET institution_id = $1,
    firstname = $2,
    lastname = $3,
    gov_id = $4,
    birthdate = $5,
    email = $6,
    phone_number = $7,
    credentials = $8,
    password = crypt($9, gen_salt('bf')),
    pending = $10,
    sex = $11
WHERE id = $12
RETURNING *;
-- name: DeleteNurseByID :exec
DELETE FROM nurse
WHERE id = $1;
-- name: GetPatientByID :one
SELECT *
FROM patient
WHERE id = $1;
-- name: GetPatientByGovID :one
SELECT *
FROM patient
WHERE gov_id = $1;
-- name: GetPatientByLogin :one
SELECT *
FROM patient
WHERE email = $1 AND password = crypt($2, password);
-- name: CreatePatient :one
INSERT INTO patient(
        firstname,
        lastname,
        gov_id,
        birthdate,
        email,
        password,
        phone_number,
        sex,
        pending,
        status,
        bed
    )
VALUES ($1, $2, $3, $4, $5, crypt($6, gen_salt('bf')), $7, $8, $9, $10, $11)
RETURNING *;
-- name: UpdatePatientByID :one
UPDATE patient
SET firstname = $1,
    lastname = $2,
    gov_id = $3,
    birthdate = $4,
    email = $5,
    password =  crypt($6, password),
    phone_number = $7,
    sex = $8,
    pending = $9,
    status = $10,
    bed = $11
WHERE id = $12
RETURNING *;
-- name: DeletePatientByID :exec
DELETE FROM patient
WHERE id = $1;
-- name: ListPatientsTreatedByDoctorID :many
SELECT *
FROM patient
WHERE id = ANY(
    SELECT patient_id 
    FROM doctor_access_request 
    WHERE doctor_id = $1
);
-- name: ListPatients :many
SELECT *
FROM patient;
-- name: ListPatientsTreatedByDoctorIDWithHealthRecordOfSpecialtyID :many
SELECT *
FROM patient
WHERE id = ANY(
    SELECT patient_id 
    FROM doctor_access_request 
    WHERE doctor_id = $1
) AND
(
    SELECT COUNT(*) 
    FROM health_record 
    WHERE specialty_id = $2
) > 0;
-- name: GetSpecialtyByID :one
SELECT *
FROM specialty
WHERE id = $1;
-- name: ListSpecialties :many
SELECT *
FROM specialty;
-- name: CreateSpecialty :one
INSERT INTO specialty(description, name)
VALUES ($1, $2)
RETURNING *;
-- name: UpdateSpecialtyByID :one
UPDATE specialty
SET description = $1,
    name = $2
WHERE id = $3
RETURNING *;
-- name: DeleteSpecialtyByID :exec
DELETE FROM specialty
WHERE id = $1;
-- name: ListSpecialtyDoctorJunctionsByDoctorID :many
SELECT *
FROM doctor_specialty
WHERE doctor_id = $1;
-- name: ListSpecialtyDoctorJunctionsBySpecialtyID :many
SELECT *
FROM doctor_specialty
WHERE specialty_id = $1;
-- name: CreateSpecialtyDoctorJunction :one
INSERT INTO doctor_specialty(doctor_id, specialty_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteSpecialtyDoctorJunction :exec
DELETE FROM doctor_specialty
WHERE doctor_id = $1
    AND specialty_id = $2;

-- SELECT  FROM SpecialtyName WHERE 1;
-- INSERT INTO SpecialtyName() VALUES ();
-- UPDATE SpecialtyName SET  WHERE 1;
-- DELETE FROM specialty_name WHERE id = $1;