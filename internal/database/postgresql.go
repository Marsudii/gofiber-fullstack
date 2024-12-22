package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// ConnectPostgreSQL establishes a connection to the PostgreSQL database
func ConnectPostgreSQL() (*pgx.Conn, error) {
	connStr := os.Getenv("DATABASE_URL") // Ensure the environment variable is correctly named
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
