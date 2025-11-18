package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"location-service/internal/model"
	"os"
	"strings"
	"testing"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/location_test?sslmode=disable"
	}
	t.Logf("connecting to %s", dsn)

	dsn = strings.Replace(dsn, "location?", "location_test?", 1)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("failed to connect to db: %v", err)
	}

	if _, err = db.Exec(context.Background(), "TRUNCATE cities RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate cities: %v", err)

	}

	return db
}

func TestCityRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCityRepository(db)

	city := model.City{Name: "New City"}
	err := repo.Create(context.Background(), &city)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if city.Id == 0 {
		t.Errorf("expected city.Id to be assigned")
	}
}

func TestCityRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCityRepository(db)

	_, _ = db.Exec(context.Background(), "INSERT INTO cities (name) VALUES ('Omsk'), ('Kurgan')")

	cities, err := repo.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cities) != 2 {
		t.Errorf("expected 2 cities, got %d", len(cities))
	}
}

func TestCityRepository_GetById(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCityRepository(db)

	_, _ = db.Exec(context.Background(), "INSERT INTO cities (id, name) VALUES ('1', 'Omsk')")

	city, err := repo.GetByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if city.Name != "Omsk" {
		t.Errorf("expected city.name to be 'Omsk', got %s", city.Name)
	}
}

func TestCityRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCityRepository(db)

	_, _ = db.Exec(context.Background(), "INSERT INTO cities (id, name) VALUES ('1', 'Kurgan')")

	city := &model.City{Id: 1, Name: "Almaty"}
	err := repo.Update(context.Background(), city)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updated, _ := repo.GetByID(context.Background(), 1)
	if updated.Name != "Almaty" {
		t.Errorf("expected city.name to be 'Almaty', got %s", updated.Name)
	}
}

func TestCityRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCityRepository(db)

	_, _ = db.Exec(context.Background(), "INSERT INTO cities (id, name) VALUES (1, 'Astana')")

	err := repo.Delete(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = repo.GetByID(context.Background(), 1)

	if err == nil {
		t.Errorf("expected error, got nil(city should be deleted)")
	}

}
