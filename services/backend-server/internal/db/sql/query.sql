------------------------------------------- GOVERNMENT -------------------------
-- name: GetGovernmentByLogin :one
SELECT *
FROM government
WHERE email = $1
    AND password = crypt($2, password);
-- name: CreateGovernment :one
INSERT INTO government(email, password)
VALUES ($1, $2)
RETURNING *;
-- name: UpdateGovernmentByID :one
UPDATE government
SET email = $1
WHERE id = $2
RETURNING *;
-- name: GetGovernment :one
SELECT *
FROM government
LIMIT 1;
------------------------------------------ INSTITUTION -------------------------
-- name: GetInstitutionByID :one
SELECT *
FROM institution
WHERE id = $1;
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
            AND pending = FALSE
    );
-- name: CreateInstitution :one
INSERT INTO institution (
        government_id,
        name,
        address,
        credentials,
        type,
        gov_id
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
    pending = $6,
    government_id = $7
WHERE id = $8
RETURNING *;
-- name: DeleteInstitutionByID :exec
DELETE FROM institution
WHERE id = $1;
------------------------------------- DOCTOR -----------------------------------
-- name: GetDoctorByID :one
SELECT *
FROM doctor
WHERE id = $1;
-- name: GetDoctorByLogin :one
SELECT d.*
FROM doctor d
    JOIN institution_enrollment_request er ON er.doctor_id = d.id
WHERE email = $1
    AND password = crypt($2, password)
    AND er.pending = FALSE
    AND er.approved = TRUE;
-- name: ListDoctors :many
SELECT *
FROM doctor;
-- name: ListDoctorsByInstitutionID :many
SELECT d.*
FROM doctor d
JOIN institution_enrollment_request er ON er.doctor_id = d.id
WHERE d.institution_id = $1 AND er.approved = TRUE;
-- name: CreateDoctor :one
INSERT INTO doctor(
        institution_id,
        firstname,
        lastname,
        gov_id,
        birthdate,
        email,
        password,
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
        $6,
        crypt($7, gen_salt('bf')),
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
    email = $6,
    phone_number = $7,
    credentials = $8,
    pending = $9,
    patient_pending = $10,
    sex = $11
WHERE id = $12
RETURNING *;
-- name: DeleteDoctorByID :exec
DELETE FROM doctor
WHERE id = $1;
--------------------------------- PATIENT --------------------------------------
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
WHERE email = $1
    AND password = crypt($2, password);
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
        bed,
        private_key,
        blockchain_address
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
        $10,
        $11,
        '',
        ''
    )
RETURNING *;
-- name: UpdatePatientByID :one
UPDATE patient
SET firstname = $1,
    lastname = $2,
    gov_id = $3,
    birthdate = $4,
    email = $5,
    phone_number = $6,
    sex = $7,
    pending = $8,
    status = $9,
    bed = $10,
    institution_id = $11
WHERE id = $12
RETURNING *;
-- name: SetPatientAddressAndPrivateKey :one
UPDATE patient
SET blockchain_address = $1,
    private_key = $2
WHERE id = $3
RETURNING *;
-- name: DeletePatientByID :exec
DELETE FROM patient
WHERE id = $1;
-- name: ListPatientsTreatedByDoctorID :many
SELECT p.*
FROM patient p
    JOIN doctor_access_request dar ON p.id = dar.patient_id
WHERE dar.doctor_id = $1
    AND dar.pending = FALSE
    and dar.approved = TRUE;
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
    )
    AND (
        SELECT COUNT(*)
        FROM health_record
        WHERE specialty_id = $2
    ) > 0;
-- name: ListPatientsByInstitutionID :many
SELECT p.*
FROM patient p
WHERE institution_id = $1;
----------------------------------- NURSE --------------------------------------
-- name: GetNurseByID :one
SELECT *
FROM nurse
WHERE id = $1;
-- name: ListNurses :many
SELECT *
FROM nurse;
-- name: ListNursesByInstitutionID :many
SELECT n.*
FROM nurse n
JOIN institution_enrollment_request er ON er.doctor_id = n.id
WHERE n.institution_id = $1 AND er.approved = TRUE;
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
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        crypt($9, gen_salt('bf')),
        $10,
        $11
    )
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
    pending = $9,
    sex = $10
WHERE id = $11
RETURNING *;
-- name: DeleteNurseByID :exec
DELETE FROM nurse
WHERE id = $1;
-- name: GetNurseByLogin :one
SELECT n.*
FROM nurse n
    JOIN institution_enrollment_request er ON er.nurse_id = n.id
WHERE email = $1
    AND password = crypt($2, password)
    AND er.pending = FALSE
    AND er.approved = TRUE;
------------------------------ INSTITUTION USER --------------------------------
-- name: GetInstitutionUserByLogin :one
SELECT iu.*
FROM institution_user iu
    JOIN government_enrollment_request er ON er.institution_id = iu.institution_id
WHERE email = $1
    AND password = crypt($2, password)
    AND er.pending = FALSE
    AND er.approved = TRUE;
-- name: GetInstitutionUserByID :many
SELECT *
FROM institution_user
WHERE id = $1;
-- name: GetInstitutionUserByGovAndInstitutionID :one
SELECT *
FROM institution_user
WHERE gov_id = $1
    AND institution_id = $2;
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
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        crypt($7, gen_salt('bf')),
        $8,
        $9
    )
RETURNING *;
-- name: UpdateInstitutionUserByGovID :one
UPDATE institution_user
SET firstname = $2,
    lastname = $3,
    birthdate = $5,
    email = $6,
    phone_number = $7,
    role = $8
WHERE gov_id = $4
    AND institution_id = $1
RETURNING *;
-- name: DeleteInstitutionUserByInsitutionAndUserID :exec
DELETE FROM institution_user
WHERE id = $1
    AND institution_id = $2;
-------------------------------- HEALTH RECORD ---------------------------------
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
        author,
        title,
        description,
        public_key,
        type,
        specialty_id,
        content_format
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
-- name: DeleteHealthRecordByID :exec
DELETE FROM health_record
WHERE id = $1;
-- name: DeleteHealthRecordDataByID :exec
UPDATE health_record
SET public_key = null
WHERE id = $1;
-- name: GetHealthRecordByID :one
SELECT *
FROM health_record
WHERE id = $1;
------------------------------- SPECIALTIES ------------------------------------
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
--------------------------------------------------------------------------------
------------------------------- DOC ACCESS REQUEST -----------------------------
-- name: GetAccessRequestsByID :one
SELECT *
FROM doctor_access_request
WHERE id = $1;
-- name: GetAccessRequestsByPatientAndDoctorID :one
SELECT *
FROM doctor_access_request
WHERE patient_id = $1
    AND doctor_id = $2;
-- name: ListAccessRequestsByPatientID :many
SELECT *
FROM doctor_access_request
WHERE patient_id = $1;
-- name: ListApprovedAccessRequestsByPatientID :many
SELECT *
FROM doctor_access_request
WHERE patient_id = $1
    AND approved = TRUE
    AND pending = FALSE;
-- name: ListAccessRequestsByDoctorID :many
SELECT *
FROM doctor_access_request
WHERE doctor_id = $1;
-- name: CountPendingAccessRequestsByDoctorID :one
SELECT COUNT (*)
FROM doctor_access_request
WHERE doctor_id = $1
    AND pending = TRUE;
-- name: CountPendingAccessRequestsByPatientID :one
SELECT COUNT (*)
FROM doctor_access_request
WHERE patient_id = $1
    AND pending = TRUE;
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
--------------------------- INSTITUTION ENROLLMENT -----------------------------
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
-- name: CountPendingInstitutionEnrollmentRequestByDoctorID :one
SELECT COUNT (*)
FROM institution_enrollment_request
WHERE doctor_id = $1
    AND pending = TRUE;
-- name: GetInstitutionEnrollmentRequestByDoctorIDAndInstitutionID :one
SELECT *
FROM institution_enrollment_request
WHERE doctor_id = $1
    AND institution_id = $2;
-- name: ListInstitutionEnrollmentRequestByNurseID :many
SELECT *
FROM institution_enrollment_request
WHERE nurse_id = $1;
-- name: CountPendingInstitutionEnrollmentRequestByNurseID :one
SELECT COUNT(*)
FROM institution_enrollment_request
WHERE nurse_id = $1
    AND pending = TRUE;
-- name: CreateInstitutionEnrollmentRequest :one
INSERT INTO institution_enrollment_request(
        institution_id,
        doctor_id,
        nurse_id
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: DeleteInstitutionEnrollmentRequestByProfID :exec
DELETE FROM institution_enrollment_request
WHERE nurse_id = $1
    OR doctor_id = $1;
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
------------------------------- GOV ENROLLMENT ---------------------------------
-- name: ListGovernmentEnrollmentRequests :many
SELECT *
FROM government_enrollment_request;
-- name: CountPendingGovernmentEnrollmentRequestsByInstitutionID :one
SELECT COUNT (*)
FROM government_enrollment_request
WHERE institution_id = $1
    AND pending = TRUE;
-- name: CreateGovernmentEnrollmentRequests :one
INSERT INTO government_enrollment_request (institution_id, government_id)
VALUES ($1, $2)
RETURNING *;
-- name: UpdatePendingGovernmentEnrollmentRequestsByID :one
UPDATE government_enrollment_request
SET institution_id = $1,
    pending = $2,
    approved = $3
WHERE id = $4
    AND pending = TRUE
RETURNING *;
-- name: DeleteGovernmentEnrollmentRequestByInsitutionID :exec
DELETE FROM government_enrollment_request
WHERE institution_id = $1;
-- name: GetGovernmentEnrollmentRequestByID :one
SELECT *
FROM government_enrollment_request
WHERE id = $1
    AND pending = TRUE;
-- name: GetGovernmentEnrollmentRequestByInsitutionID :one
SELECT *
FROM government_enrollment_request
WHERE institution_id = $1;