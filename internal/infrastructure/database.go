package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
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
