package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/model"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CityRepositoryInterface interface {
	Create(ctx context.Context, city *model.City) error
	GetAll(ctx context.Context) ([]model.City, error)
	GetByID(ctx context.Context, id int) (*model.City, error)
	Update(ctx context.Context, city *model.City) error
	Delete(ctx context.Context, id int) error
}

type CityRepository struct {
	db *pgxpool.Pool
}

func NewCityRepository(db *pgxpool.Pool) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) Create(ctx context.Context, city *model.City) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id`

	err := r.db.QueryRow(ctx, query, city.Name).Scan(&city.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_city_name") {
					return errors.New("city with this name already exists")
				}
			}
		}
		return err
	}
	return nil
}

func (r *CityRepository) GetAll(ctx context.Context) ([]model.City, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := []model.City{}
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
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_city_name") {
					return errors.New("city with this name already exists")
				}
			}
		}
		return err
	}

	return nil
}

var _ CityRepositoryInterface = (*CityRepository)(nil)
