package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	// Docker Compose の設定に合わせた接続 URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("Could not ping DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
}
