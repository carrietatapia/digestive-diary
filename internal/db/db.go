package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(connString string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewTestDB() (*pgxpool.Pool, error) {
	connTestString := os.Getenv("TEST_DB_CONN_STRING")
	if connTestString == "" {
		connTestString = "postgres://postgres:admin@localhost:5432/test"
	}

	db, err := pgxpool.Connect(context.Background(), connTestString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CleanupDB(db *pgxpool.Pool) {
	_, err := db.Exec(context.Background(), "TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	if err != nil {
		log.Fatalf("Failed to clean up database: %v", err)
	}
}
