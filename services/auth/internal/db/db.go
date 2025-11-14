package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	return pool
}
