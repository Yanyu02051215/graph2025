package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// Docker Compose の設定に合わせた接続 URL
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		"postgresql",  // ユーザー名 (POSTGRES_USER)
		"postgresql",  // パスワード (POSTGRES_PASSWORD)
		"localhost",          // ホスト名 (サービス名)
		5434,          // ポート番号 (Docker内部のデフォルト)
		"api",         // データベース名 (POSTGRES_DB)
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
