package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"location-service/internal/model"
	"strings"
)

type StreetRepositoryInterface interface {
	Create(ctx context.Context, city *model.Street) error
	GetAll(ctx context.Context) ([]model.Street, error)
	GetByID(ctx context.Context, id int) (*model.Street, error)
	Update(ctx context.Context, city *model.Street) error
	Delete(ctx context.Context, id int) error
}

type StreetRepository struct {
	db *pgxpool.Pool
}

func NewStreetRepository(db *pgxpool.Pool) *StreetRepository {
	return &StreetRepository{db: db}
}

func (r *StreetRepository) Create(ctx context.Context, street *model.Street) error {
	query := `INSERT INTO streets (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(ctx, query, street.Name).Scan(&street.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_street_name") {
					return errors.New("street with this name already exists")
				}
			}
		}

		return err
	}
	return nil
}

func (r *StreetRepository) GetAll(ctx context.Context) ([]model.Street, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM streets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	streets := []model.Street{}
	for rows.Next() {
		var s model.Street

		if err := rows.Scan(&s.Id, &s.Name); err != nil {
			return nil, err
		}
		streets = append(streets, s)
	}
	return streets, nil
}

func (r *StreetRepository) GetByID(ctx context.Context, id int) (*model.Street, error) {
	var street model.Street
	err := r.db.QueryRow(ctx, "SELECT id, name FROM streets WHERE id=$1", id).
		Scan(&street.Id, &street.Name)
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
	query := `UPDATE streets SET name = $1 where id = $2`
	_, err := r.db.Exec(ctx, query, street.Name, street.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.Message, "unique_street_name") {
					return errors.New("street with this name already exists")
				}
			}
		}
		return err
	}

	return nil
}
