package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		"postgresql",
		"postgresql",
		"localhost",
		5434,
		"api",
	)

	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Could not ping DB: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL")
	return &Database{Conn: pool}, nil
}
