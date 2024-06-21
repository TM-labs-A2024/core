CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TYPE INSTITUTION_TYPE AS ENUM ('clinic', 'hospital');
CREATE TYPE HEALTH_RECORD_TYPE AS ENUM ('study', 'report', 'evolution', 'order');
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
CREATE TABLE IF NOT EXISTS institution (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    credentials TEXT NOT NULL,
    type INSTITUTION_TYPE NOT NULL,
    gov_id TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS specialty (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    description TEXT NOT NULL,
    name SPECIALTY_NAME NOT NULL,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS nurse (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    credentials TEXT NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id)
);
CREATE TABLE IF NOT EXISTS patient (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
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
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id)
);
CREATE TABLE IF NOT EXISTS doctor (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
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
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id)
);
CREATE TABLE IF NOT EXISTS doctor_access_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    doctor_id UUID NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT fk_patient FOREIGN KEY(patient_id) REFERENCES patient(id),
    CONSTRAINT fk_doctor FOREIGN KEY(doctor_id) REFERENCES doctor(id)
);
CREATE TABLE IF NOT EXISTS institution_enrollment_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
    doctor_id UUID,
    nurse_id UUID,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id),
    CONSTRAINT fk_doctor FOREIGN KEY(doctor_id) REFERENCES doctor(id),
    CONSTRAINT fk_nurse FOREIGN KEY(nurse_id) REFERENCES nurse(id),
    CONSTRAINT chk_professional CHECK (
        doctor_id IS NOT NULL
        OR nurse_id IS NOT NULL
    )
);
CREATE TABLE IF NOT EXISTS government (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS health_record (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    patient_id UUID NOT NULL,
    private_key TEXT NOT NULL,
    public_key TEXT NOT NULL,
    type HEALTH_RECORD_TYPE NOT NULL,
    specialty_id UUID NOT NULL,
    content_format TEXT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_patient FOREIGN KEY(patient_id) REFERENCES patient(id),
    CONSTRAINT fk_specialty FOREIGN KEY(specialty_id) REFERENCES specialty(id)
);
CREATE TABLE IF NOT EXISTS government_enrollment_request (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
    pending BOOLEAN NOT NULL DEFAULT true,
    approved BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id)
);
CREATE TABLE IF NOT EXISTS institution_user (
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    institution_id UUID NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    gov_id TEXT NOT NULL,
    birthdate DATE NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    role TEXT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_institution FOREIGN KEY(institution_id) REFERENCES institution(id)
);
CREATE TABLE doctor_specialty(
    doctor_id UUID REFERENCES doctor(id),
    specialty_id UUID REFERENCES specialty(id),
    CONSTRAINT doctor_specialty_pk PRIMARY KEY(doctor_id, specialty_id)
);