CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TYPE PROFESSIONAL_TYPE AS ENUM ('doctor', 'nurse');
CREATE TYPE INSTITUTION_TYPE AS ENUM ('clinic', 'hospital');
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
    'Urology'
);
CREATE TABLE IF NOT EXISTS doctor (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid UUID NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    credentials TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    patient_pending BOOLEAN NOT NULL DEFAULT true,
    specialities INTEGER [] NOT NULL,
    PRIMARY KEY(uuid),
    CONSTRAINT fk_institution FOREIGN KEY(institution_uuid) REFERENCES institution(uuid)
);
CREATE TABLE IF NOT EXISTS doctor_access_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_uuid UUID NOT NULL,
    doctor_uuid UUID NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(uuid),
    CONSTRAINT fk_patient FOREIGN KEY(patient_uuid) REFERENCES patient(uuid),
    CONSTRAINT fk_doctor FOREIGN KEY(doctor_uuid) REFERENCES doctor(uuid)
);
CREATE TABLE IF NOT EXISTS institution_enrollment_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid UUID NOT NULL,
    doctor_uuid UUID,
    nurse_uuid UUID,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    professional_type PROFESSIONAL_TYPE NOT NULL,
    PRIMARY KEY(uuid),
    CONSTRAINT fk_institution FOREIGN KEY(institution_uuid) REFERENCES institution(uuid),
    CONSTRAINT fk_doctor FOREIGN KEY(doctor_uuid) REFERENCES doctor(uuid),
    CONSTRAINT fk_nurse FOREIGN KEY(nurse_uuid) REFERENCES nurse(uuid)
);
CREATE TABLE IF NOT EXISTS government (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS health_record (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_uuid UUID NOT NULL,
    private_key TEXT NOT NULL,
    type TEXT NOT NULL,
    specialty TEXT NOT NULL,
    content_format TEXT NOT NULL,
    PRIMARY KEY(uuid),
    CONSTRAINT fk_patient FOREIGN KEY(patient_uuid) REFERENCES patient(uuid)
);
CREATE TABLE IF NOT EXISTS institution (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    credentials TEXT NOT NULL,
    type TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS government_enrollment_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid UUID NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS institution_user (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid UUID NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    role TEXT NOT NULL,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS nurse (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_uuid UUID NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    credentials TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS patient (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    sex TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    status TEXT NOT NULL,
    bed TEXT NOT NULL,
    PRIMARY KEY(uuid)
);
CREATE TABLE IF NOT EXISTS specialty (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    uuid UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    id INT NOT NULL,
    description TEXT NOT NULL,
    name SPECIALTY_NAME NOT NULL,
    PRIMARY KEY(uuid)
);