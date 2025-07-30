package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ROMUSKING/hiseas/backend/internal/user"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for sqlx
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Load .env file for local development. This is non-fatal.
	if err := godotenv.Load("../.env"); err != nil {
		logger.Info("No .env file found, relying on environment variables")
	}

	cfg := loadConfig()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	// Use "pgx" as the driver name for sqlx with pgx/v5.
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("Database connection established")

	// Run goose migrations manually before starting the server
	// goose -dir ./db/migrations postgres "%s" up

	// Initialize router
	r := chi.NewRouter()

	// Create an auth middleware instance with the configured token.
	// This decouples the middleware from global state (os.Getenv).
	authMiddleware := user.NewAuthMiddleware(cfg.APIToken)
	r.Use(authMiddleware)

	// Register module routes
	user.RegisterRoutes(r, db)

	// Setup server with graceful shutdown
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: r,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		logger.Info("Starting server", "address", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server startup failed", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exiting gracefully")
}
