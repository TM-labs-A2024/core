CREATE EXTENSION pgcrypto;
CREATE TYPE PROFESSIONAL_TYPE AS ENUM ('doctor', 'nurse');
CREATE TYPE SPECIALTY_NAME AS ENUM (
    'Allergy and immunology',
    'Anesthesiology',
    'Dermatology',
    'Diagnostic radiology',
    'Emergency medicine',
    'Family medicine',
    'Internal medicine',
    'Medical genetics',
    'Neurology',
    'Nuclear medicine',
    'Obstetrics and gynecology',
    'Ophthalmology',
    'Pathology',
    'Pediatrics',
    'Physical medicine and rehabilitation',
    'Preventive medicine',
    'Psychiatry',
    'Radiation oncology',
    'Surgery',
    'Urology',
);
CREATE TABLE IF NOT EXISTS Doctor (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    credentials TEXT NOT NULL pending BOOLEAN NOT NULL DEFAULT false,
    patient_pending BOOLEAN NOT NULL DEFAULT false,
    specialities INTEGER[] NOT NULL
);
CREATE TABLE IF NOT EXISTS DoctorAccessRequest (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_uuid TEXT NOT NULL,
    doctor_uuid TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false
);
CREATE TABLE IF NOT EXISTS InstitutionEnrollmentRequest(
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid TEXT NOT NULL,
    doctor_uuid TEXT,
    nurse_uuid TEXT,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN DEFAULT false,
    professional_type PROFESSIONAL_TYPE NOT NULL
);
-- CREATE TABLE IF NOT EXISTS _doctors__doctor_uuid__
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     sex TEXT NOT NULL,
--     pending BOOLEAN NOT NULL, DEFAULT false,
--     status TEXT NOT NULL,
--     bed TEXT NOT NULL,
-- );
-- CREATE TABLE IF NOT EXISTS _doctors_get_200_respo
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,
--     specialities JSON NOT NULL,
--     pending BOOLEAN NOT NULL, DEFAULT false,
--     patient_pending BOOLEAN NOT NULL DEFAULT false,
-- );
-- CREATE TABLE IF NOT EXISTS _doctors_post_request 
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,
--     password TEXT NOT NULL,
--     specialties JSON NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _doctors_put_request (
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,
--     password TEXT NOT NULL,
--     specialties JSON NOT NULL
-- );
CREATE TABLE IF NOT EXISTS Government (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
);
CREATE TABLE IF NOT EXISTS HealthRecord (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_uuid UUID NOT NULL,
    content TEXT NOT NULL,
    type TEXT NOT NULL,
    specialty TEXT NOT NULL,
    content_format TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS Institution (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    gov_id TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS InstitutionEnrollmentRequest (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT false,
    approved BOOLEAN DEFAULT false DEFAULT false,
);
CREATE TABLE IF NOT EXISTS InstitutionUser (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    role TEXT NOT NULL
);
-- CREATE TABLE IF NOT EXISTS _institutions_enrollme
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     doctor_uuid TEXT NOT NULL,
--     pending BOOLEAN NOT NULL, DEFAULT false,
--     approved BOOLEAN DEFAULT false, DEFAULT false,
--     professional_type PROFESSIONAL_TYPE
-- );
-- CREATE TABLE IF NOT EXISTS _institutions_get_200_
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     name TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
-- );
-- CREATE TABLE IF NOT EXISTS _institutions__institu
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     role TEXT NOT NULL,
-- );
-- CREATE TABLE IF NOT EXISTS _institutions__institu
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     role TEXT NOT NULL,
--     password TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _institutions__institu
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     role TEXT NOT NULL,,
--     password TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _institutions_put_200_
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     name TEXT NOT NULL,
--     gov_id TEXT NOT NULL,,
--     pending BOOLEAN NOT NULL DEFAULT false,
-- );
CREATE TABLE IF NOT EXISTS Nurse (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    credentials TEXT NOT NULL
);
-- CREATE TABLE IF NOT EXISTS _nurses_get_200_respon
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,,
--     pending BOOLEAN NOT NULL DEFAULT false,
-- );
-- CREATE TABLE IF NOT EXISTS _nurses_post_request (
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,
--     password TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _nurses_put_request (
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     institution_uuid TEXT NOT NULL,
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     credentials TEXT NOT NULL,,
--     password TEXT NOT NULL
-- );
CREATE TABLE IF NOT EXISTS Patient (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    sex TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT false,
    status TEXT NOT NULL,
    bed TEXT NOT NULL
);
-- CREATE TABLE IF NOT EXISTS _patients_access_reque
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     patient_uuid TEXT NOT NULL,
--     doctor_uuid TEXT NOT NULL,,
--     completed BOOLEAN NOT NULL DEFAULT false,
-- );
-- CREATE TABLE IF NOT EXISTS _patients__gov_id__medi
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     content TEXT NOT NULL,
--     type TEXT NOT NULL,
--     specialty TEXT NOT NULL,
--     content_format TEXT NOT NULL,
-- );
-- CREATE TABLE IF NOT EXISTS _patients__gov_id__orde
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     content TEXT NOT NULL,
--     type TEXT,
--     specialty TEXT NOT NULL,
--     content_format TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _patients_post_request
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     sex TEXT NOT NULL,
--     pending BOOLEAN NOT NULL, DEFAULT false,
--     status TEXT NOT NULL,
--     bed TEXT NOT NULL,
--     password TEXT NOT NULL
-- );
-- CREATE TABLE IF NOT EXISTS _patients_put_request 
--     uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
--     firstname TEXT NOT NULL,
--     lastname TEXT NOT NULL,
--     gov_id TEXT NOT NULL,
--     birthdate DATE NOT NULL,
--     email TEXT NOT NULL,
--     password TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     sex TEXT NOT NULL,
--     pending BOOLEAN NOT NULL, DEFAULT false,
--     status TEXT NOT NULL,
--     bed TEXT NOT NULL,,
--     password TEXT NOT NULL
-- );
CREATE TABLE IF NOT EXISTS Specialty (
    created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    id INT NOT NULL,
    description TEXT NOT NULL,
    name SPECIALTY_NAME NOT NULL
);