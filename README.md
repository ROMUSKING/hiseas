# HiSeas Project

HiSeas is a modular monolith backend for a sailing application, built in Go and using PostgreSQL with the PostGIS extension for advanced geospatial features. The backend is located in the `/backend` folder and follows a domain-driven modular structure.

## Features
- User management, authentication, and reputation system
- Voyage planning, crew postings, and matching
- Geospatial queries powered by PostGIS
- Real-time chat via WebSocket
- Weather API proxy and caching
- Privacy controls and encrypted data storage
- Database migrations managed with Goose

## Technology Stack
- Go 1.21+
- PostgreSQL + PostGIS
- chi (HTTP router)
- pgx (PostgreSQL driver)
- sqlx (DB access)
- gorilla/websocket (WebSocket)
- goose (migrations)

## Structure
- `/backend/cmd/server/main.go`: Application entry point
- `/backend/internal/`: Domain modules
- `/backend/db/migrations/`: Migration scripts

## Licensing
HiSeas is distributed under the GNU General Public License v2.0 or later (GPLv2+), compatible with the requirements of PostGIS. See the LICENSE file for details.

## Local Setup Instructions
1. **Install required tools:**
   - Go 1.21+ ([go.dev/dl](https://go.dev/dl/))
   - PostgreSQL ([postgresql.org/download](https://www.postgresql.org/download/))
   - PostGIS ([postgis.net/install](https://postgis.net/install/))
   - Goose ([github.com/pressly/goose](https://github.com/pressly/goose))

2. **Clone the repository:**
   ```bash
   git clone https://github.com/ROMUSKING/hiseas.git
   cd hiseas
   ```

3. **Configure environment:**
   - Create a `.env` file in the project root:
     ```env
     DB_USER=postgres
     DB_PASSWORD=yourpassword
     DB_HOST=localhost
     DB_PORT=5432
     DB_NAME=hiseas
     DB_SSLMODE=disable
     SERVER_PORT=8080
     ```

4. **Install Go dependencies:**
   ```bash
   cd backend
   go mod tidy
   ```

5. **Create database and enable extensions:**
   ```bash
   createdb hiseas
   psql -d hiseas -c "CREATE EXTENSION IF NOT EXISTS postgis;"
   psql -d hiseas -c "CREATE EXTENSION IF NOT EXISTS pgcrypto;"
   ```

6. **Run database migrations:**
   ```bash
   goose -dir ./db/migrations postgres "postgres://postgres:yourpassword@localhost:5432/hiseas?sslmode=disable" up
   ```

7. **Start the backend server:**
   ```bash
   go run ./cmd/server/main.go
   ```

8. The backend will be available at `http://localhost:8080` (or as configured in `.env`).

## References
- [PostGIS License](https://github.com/postgis/postgis/blob/master/LICENSE.TXT)
- [GNU GPL v2.0](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html)
