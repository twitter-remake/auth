BEGIN;

CREATE DATABASE auth;

USE auth;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS identities(
    "id" uuid DEFAULT uuid_generate_v4(),
    "uid" varchar(255) NOT NULL, -- unique id (since we're using firebase auth)
    "email" varchar(255) NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz NULL DEFAULT NULL,
    PRIMARY KEY ("id")
);

COMMIT;

