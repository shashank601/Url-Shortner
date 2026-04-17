package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() *pgxpool.Pool {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	log.Println("connected to postgres")

	return pool
}
