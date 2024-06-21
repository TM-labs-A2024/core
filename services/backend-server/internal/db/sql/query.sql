-- name: GetDoctorByID :one
SELECT id,
    institution_id,
    firstname,
    lastname,
    gov_id,
    birthdate,
    email,
    phone_number,
    credentials,
    pending,
    patient_pending
FROM doctor
WHERE id = $1;
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
        pending,
        patient_pending
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    )
RETURNING *;
-- name: UpdateDoctorByID :one
UPDATE doctor
SET institution_id = $1,
    firstname = $2,
    lastname = $3,
    gov_id = $4,
    birthdate = $5,
    password = $6,
    email = $7,
    phone_number = $8,
    credentials = $9,
    pending = $10,
    patient_pending = $11
WHERE id = $12
RETURNING *;
-- name: DeleteDoctorByID :exec
DELETE FROM doctor
WHERE id = $1;
-- name: GetAccessRequestByPatientID :one
SELECT id,
    patient_id,
    doctor_id,
    pending,
    approved
FROM doctor_access_request
WHERE patient_id = $1;
-- name: GetAccessRequestByDoctorID :one
SELECT id,
    patient_id,
    doctor_id,
    pending,
    approved
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
SELECT id,
    institution_id,
    doctor_id,
    nurse_id,
    pending,
    approved
FROM institution_enrollment_request
WHERE institution_id = $1;
-- name: GetInstitutionEnrollmentRequestsByID :one
SELECT id,
    institution_id,
    doctor_id,
    nurse_id,
    pending,
    approved
FROM institution_enrollment_request
WHERE id = $1;
-- name: CreateInstitutionEnrollmentRequest :one
INSERT INTO institution_enrollment_request(
        institution_id,
        doctor_id,
        nurse_id
    )
VALUES ($1, $2, $3)
RETURNING *;
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
-- name: GetGovernment :one
SELECT id,
    email
FROM government
LIMIT 1;
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
-- name: GetHealthRecordsBySpecialtyID :many
SELECT id,
    patient_id,
    private_key,
    type,
    specialty_id,
    content_format
FROM health_record
WHERE specialty_id = $1;
-- name: GetHealthRecordsBySpecialtyIDAndPatientID :many
SELECT id,
    patient_id,
    private_key,
    type,
    specialty_id,
    content_format
FROM health_record
WHERE specialty_id = $1
    AND patient_id = $2;
-- name: GetHealthRecordsByPatientID :many
SELECT id,
    patient_id,
    private_key,
    type,
    specialty_id,
    content_format
FROM health_record
WHERE patient_id = $1;
-- name: CreateHealthRecord :one
INSERT INTO health_record (
        patient_id,
        private_key,
        type,
        specialty_id,
        content_format
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateHealthRecordByID :one
UPDATE health_record
SET patient_id = $1,
    private_key = $2,
    type = $3,
    specialty_id = $4,
    content_format = $5
WHERE id = $6
RETURNING *;
-- name: DeleteHealthRecordByID :exec
DELETE FROM health_record
WHERE id = $1;
-- name: GetInstitutionByID :one
SELECT id,
    name,
    address,
    credentials,
    type,
    gov_id,
    pending
FROM institution
WHERE id = $1;
-- name: ListInstitutions :many
SELECT id,
    name,
    address,
    credentials,
    type,
    gov_id,
    pending
FROM institution;
-- name: ListApprovedInstitutions :many
SELECT id,
    name,
    address,
    credentials,
    type,
    gov_id,
    pending
FROM institution
WHERE id IN (SELECT institution_id FROM government_enrollment_request WHERE approved = TRUE);
-- name: InsertInstitution :one
INSERT INTO institution (
        name,
        address,
        credentials,
        type,
        gov_id,
        pending
    )
VALUES ($1, $2, $3, $4, $5, $6)
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
SELECT id,
    institution_id,
    pending,
    approved
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
-- name: GetInstitutionUserByID :many
SELECT id,
    institution_id,
    firstname,
    lastname,
    gov_id,
    birthdate,
    email,
    phone_number,
    role
FROM institution_user
WHERE id = $1;
-- name: ListInstitutionUserByInstitutionID :many
SELECT id,
    institution_id,
    firstname,
    lastname,
    gov_id,
    birthdate,
    email,
    phone_number,
    role
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;
-- name: UpdateInstitutionUserByGovID :one
UPDATE institution_user
SET institution_id = $1,
    firstname = $2,
    lastname = $3,
    gov_id = $4,
    birthdate = $5,
    email = $6,
    password = $7,
    phone_number = $8,
    role = $9
WHERE gov_id = $10
RETURNING *;
-- name: DeleteInstitutionUserByID :exec
DELETE FROM institution_user
WHERE id = $1;
-- SELECT  FROM InstitutionUserRole WHERE 1;
-- INSERT INTO InstitutionUserRole() VALUES ();
-- UPDATE InstitutionUserRole SET  WHERE 1;
-- DELETE FROM institution_user_role WHERE id = $1;
-- name: GetNurseByID :one
SELECT id,
    institution_id,
    firstname,
    lastname,
    gov_id,
    birthdate,
    email,
    phone_number,
    credentials,
    pending
FROM nurse
WHERE id = $1;
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
        pending
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
    password = $9,
    pending = $10
WHERE id = $11
RETURNING *;
-- name: DeleteNurseByID :exec
DELETE FROM nurse
WHERE id = $1;
-- name: GetPatientByID :one
SELECT id,
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
FROM patient
WHERE id = $1;
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;
-- name: UpdatePatientByID :one
UPDATE patient
SET firstname = $1,
    lastname = $2,
    gov_id = $3,
    birthdate = $4,
    email = $5,
    password = $6,
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
-- name: GetSpecialtyByID :one
SELECT id,
    description,
    name
FROM specialty
WHERE 1;
-- name: ListSpecialties :many
SELECT id,
    description,
    name
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
SELECT doctor_id,
    specialty_id
FROM doctor_specialty
WHERE doctor_id = $1;
-- name: ListSpecialtyDoctorJunctionsBySpecialtyID :many
SELECT doctor_id,
    specialty_id
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