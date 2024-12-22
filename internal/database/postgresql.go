package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"gofiber-fullstack/internal/config"
	"os"
)

var DB *pgx.Conn

func ConnectPostgreSQL() (*pgx.Conn, error) {
	// CALL LOGS
	logs := config.NewLogs()
	// CREATE CONNECTION STRING
	connStr := os.Getenv("DATABASE_URL") // Ensure the environment variable is correctly named
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		logs.Error("Failed to connect to PostgreSQL database", err.Error())
		return nil, err
	}

	// if successful, show log message
	logs.Info("Connected to PostgreSQL database")

	return conn, nil
}
