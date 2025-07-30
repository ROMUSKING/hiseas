# db/migrations

This folder contains Goose migration scripts for managing the PostgreSQL database schema.

## Usage
- Run migrations manually before starting the backend server:
  ```bash
  goose -dir ./db/migrations postgres "postgres://postgres:yourpassword@localhost:5432/hiseas?sslmode=disable" up
  ```
- See the main project README for full setup instructions.

## Notes
- Includes PostGIS and pgcrypto extensions for geospatial and encrypted data support.
