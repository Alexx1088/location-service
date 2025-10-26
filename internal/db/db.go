package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/location?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}

	log.Println("✅ Подключено к PostgreSQL")
	return pool
}
