# internal/voyage

This module manages voyage planning, crew postings, and matching logic for the HiSeas backend.

## Structure
- `postgres/repository.go`: Implements the repository pattern for voyage data access, including efficient geospatial queries using PostGIS.

## Usage
- Use the repository to query voyages, including proximity searches and route management.
- All geospatial logic is delegated to the database for performance and simplicity.
