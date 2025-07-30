package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Load config from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	serverPort := os.Getenv("SERVER_PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run goose migrations manually before starting the server
	// goose -dir ./db/migrations postgres "%s" up

	// Initialize router
	r := chi.NewRouter()

	// Register module routes
	// user.RegisterRoutes(r, db)
	// voyage.RegisterRoutes(r, db)
	// geospatial.RegisterRoutes(r, db)
	// chat.RegisterRoutes(r)
	// weather.RegisterRoutes(r)

	log.Printf("Starting server on :%s", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
