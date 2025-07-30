# HiSeas Backend

This folder contains the backend Go modular monolith for the sailing app. All backend modules, migrations, and entry points are located here.

## Structure
- `cmd/server/main.go`: Entry point
- `internal/`: All backend modules
- `db/migrations/`: Goose migration scripts

## Usage
- See the main project README for setup and running instructions.
- Environment variables are loaded from `.env` in the project root.
- Database migrations must be run manually using Goose before starting the server.

## Licensing
Distributed under GPLv2+ for compatibility with PostGIS.
