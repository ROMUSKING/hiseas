-- +goose Up
-- This section is executed when the migration is applied.

-- Step 1: Enable necessary extensions
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS postgis;

-- Step 2: Create the users table
-- This table stores user information, including encrypted data and reputation.
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    reputation INTEGER NOT NULL DEFAULT 0,
    privacy_level VARCHAR(50) NOT NULL DEFAULT 'private',
    encrypted_email BYTEA,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Step 3: Create the vessels table
-- Note: An owner_id column has been added to link a vessel to a user.
CREATE TABLE vessels (
    id BIGSERIAL PRIMARY KEY,
    owner_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Step 4: Create the voyages table
-- This table uses PostGIS geography types for accurate location data.
CREATE TABLE voyages (
    id BIGSERIAL PRIMARY KEY,
    skipper_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vessel_id BIGINT REFERENCES vessels(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMPTZ,
    end_time TIMESTAMPTZ,
    start_location geography(Point, 4326),
    planned_route geography(LineString, 4326),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Step 5: Create a spatial index for efficient nearby voyage searches
CREATE INDEX voyages_start_location_idx ON voyages USING GIST (start_location);

-- +goose Down
-- This section is executed when the migration is rolled back.
DROP INDEX IF EXISTS voyages_start_location_idx;
DROP TABLE IF EXISTS voyages;
DROP TABLE IF EXISTS vessels;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS postgis;
DROP EXTENSION IF EXISTS pgcrypto;