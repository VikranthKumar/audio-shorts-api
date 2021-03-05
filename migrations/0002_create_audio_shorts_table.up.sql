BEGIN;

CREATE TYPE audio_shorts_status AS ENUM (
    'active',
    'banned',
    'deleted'
);

CREATE TYPE category AS ENUM (
    'news',
    'gossip',
    'review',
    'story'
);

CREATE TABLE IF NOT EXISTS audio_shorts (
    "id" SERIAL PRIMARY KEY,
    "title" varchar(100) NOT NULL,
    "description" text NOT NULL,
    "status" audio_shorts_status NOT NULL,
    "category" category NOT NULL,
    "audio_file" varchar(300) NOT NULL,
    "creator_id" int NOT NULL,
    "created_at" timestamp with time zone DEFAULT now(),
    "updated_at" timestamp with time zone DEFAULT now(),
    UNIQUE ("title", "creator_id"),
    CONSTRAINT fk_creator FOREIGN KEY("creator_id") references creators("id")
);

COMMIT;