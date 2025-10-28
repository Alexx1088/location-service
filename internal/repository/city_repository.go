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

func (r *CityRepository) Create(ctx context.Context, city *model.City) error {
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

func (r *CityRepository) GetByID(ctx context.Context, id int) (*model.City, error) {
	var city model.City
	err := r.db.QueryRow(ctx, "SELECT id, name FROM cities WHERE id=$1", id).
		Scan(&city.Id, &city.Name)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *CityRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM cities WHERE id=$1", id)
	return err
}

func (r *CityRepository) Update(ctx context.Context, city *model.City) error {
	query := "UPDATE cities SET name=$1 where id=$2"
	_, err := r.db.Exec(ctx, query, city.Name, city.Id)
	return err
}
