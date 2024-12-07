package db

import (
	"context"

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
