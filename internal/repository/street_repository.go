package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StreetRepository struct {
	db *pgxpool.Pool
}

func NewStreetRepository(db *pgxpool.Pool) *StreetRepository {
	return &StreetRepository{db: db}
}

func (r *StreetRepository) Create(ctx context.Context, street *model.Street) error {
	query := `INSERT INTO streets (name, city_id) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(ctx, query, street.Name, street.CityId).Scan(&street.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return errors.New("city does not exist")
			}
		}
		return err
	}

	return nil
}

func (r *StreetRepository) GetAll(ctx context.Context) ([]model.Street, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, city_id FROM streets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var streets []model.Street
	for rows.Next() {
		var s model.Street
		if err := rows.Scan(&s.Id, &s.Name, &s.CityId); err != nil {
			return nil, err
		}
		streets = append(streets, s)
	}
	return streets, nil
}

func (r *StreetRepository) GetByID(ctx context.Context, id int) (*model.Street, error) {
	var street model.Street
	err := r.db.QueryRow(ctx, "SELECT id, name, city_id FROM streets WHERE id=$1", id).
		Scan(&street.Id, &street.Name, &street.CityId)
	if err != nil {
		return nil, err
	}
	return &street, nil
}

func (r *StreetRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM streets WHERE id=$1", id)
	return err
}

func (r *StreetRepository) Update(ctx context.Context, street *model.Street) error {
	query := `UPDATE streets SET name = $1, city_id = $2 where id = $3`
	_, err := r.db.Exec(ctx, query, street.Name, street.CityId, street.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return errors.New("city does not exist")
			}
		}
		return err
	}

	return nil
}
