package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrossroadRepository struct {
	db *pgxpool.Pool
}

func NewCrossroadRepository(db *pgxpool.Pool) *CrossroadRepository {
	return &CrossroadRepository{db: db}
}

func (r *CrossroadRepository) Create(ctx context.Context, crossroad *model.Crossroad) error {
	query := `INSERT INTO crossroads (street_id) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(ctx, query, crossroad.StreetId).Scan(&crossroad.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return errors.New("street does not exist")
			}
		}
		return err
	}

	return nil
}

func (r *CrossroadRepository) GetAll(ctx context.Context) ([]model.Crossroad, error) {
	rows, err := r.db.Query(ctx, "SELECT id, street_id FROM crossroads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crossroads []model.Crossroad
	for rows.Next() {
		var s model.Crossroad
		if err := rows.Scan(&s.Id, &s.StreetId); err != nil {
			return nil, err
		}
		crossroads = append(crossroads, s)
	}
	return crossroads, nil
}

func (r *CrossroadRepository) GetByID(ctx context.Context, id int) (*model.Crossroad, error) {
	var crossroad model.Crossroad
	err := r.db.QueryRow(ctx, "SELECT id, street_id FROM crossroads WHERE id=$1", id).
		Scan(&crossroad.Id, &crossroad.StreetId)
	if err != nil {
		return nil, err
	}
	return &crossroad, nil
}

func (r *CrossroadRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM crossroads WHERE id=$1", id)
	return err
}

func (r *CrossroadRepository) Update(ctx context.Context, crossroad *model.Crossroad) error {
	query := `UPDATE crossroad SET street_id = $1 where id = $2`
	_, err := r.db.Exec(ctx, query, crossroad.StreetId, crossroad.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return errors.New("street does not exist")
			}
		}
		return err
	}

	return nil
}
