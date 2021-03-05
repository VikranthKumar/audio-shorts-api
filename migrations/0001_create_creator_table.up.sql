BEGIN;

CREATE TYPE creator_status AS ENUM (
    'active',
    'banned',
    'suspended'
);

CREATE TABLE IF NOT EXISTS creators (
    "id" SERIAL PRIMARY KEY,
    "username" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "email" varchar(100) NOT NULL UNIQUE,
    "status" creator_status NOT NULL,
    "created_at" timestamp with time zone DEFAULT now(),
    "updated_at" timestamp with time zone DEFAULT now()
);

COMMIT;