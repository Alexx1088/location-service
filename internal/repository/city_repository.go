package repository

import (
	"context"
	"location-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CityRepository struct {
	db *pgxpool.Pool
}

func NewCityRepository(db *pgxpool.Pool) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) Create(ctx context.Context, city model.City) error {
	return r.db.QueryRow(ctx, "INSERT INTO cities (name) VALUES ($1) RETURNING id", city.Name).Scan(&city.Id)
}

func (r *CityRepository) GetAll(ctx context.Context) ([]model.City, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []model.City
	for rows.Next() {
		var c model.City
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}
