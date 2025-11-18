package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func RunMigrations() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = GetDatabaseURL()
	}
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open sql connection: %v", err)
	}
	defer sqlDB.Close()

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to init migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	log.Println("✅ Migrations applied successfully")

}
