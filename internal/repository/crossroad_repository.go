package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/model"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrossroadRepository struct {
	db *pgxpool.Pool
}

func NewCrossroadRepository(db *pgxpool.Pool) *CrossroadRepository {
	return &CrossroadRepository{db: db}
}

func (r *CrossroadRepository) Create(ctx context.Context, crossroad *model.Crossroad) error {
	query := `INSERT INTO crossroads (street_id, city_id) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(ctx, query, crossroad.StreetId, crossroad.CityId).Scan(&crossroad.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				if strings.Contains(pgErr.Message, "street_id") {
					return errors.New("street does not exist")
				}
				if strings.Contains(pgErr.Message, "city_id") {
					return errors.New("city does not exist")
				}
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_crossroad") {
					return errors.New("crossroad already exists for this street and city")
				}
			}
		}
		return err
	}

	return nil
}

func (r *CrossroadRepository) GetAll(ctx context.Context) ([]model.Crossroad, error) {
	rows, err := r.db.Query(ctx, "SELECT id, street_id, city_id FROM crossroads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crossroads []model.Crossroad
	for rows.Next() {
		var s model.Crossroad
		if err := rows.Scan(&s.Id, &s.StreetId, &s.CityId); err != nil {
			return nil, err
		}
		crossroads = append(crossroads, s)
	}
	return crossroads, nil
}

func (r *CrossroadRepository) GetByID(ctx context.Context, id int) (*model.Crossroad, error) {
	var crossroad model.Crossroad
	err := r.db.QueryRow(ctx, "SELECT id, street_id, city_id FROM crossroads WHERE id=$1", id).
		Scan(&crossroad.Id, &crossroad.StreetId, &crossroad.CityId)
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
	query := `UPDATE crossroads SET street_id = $1, city_id = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, crossroad.StreetId, crossroad.CityId, crossroad.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				if strings.Contains(pgErr.Message, "street_id") {
					return errors.New("street does not exist")
				}
				if strings.Contains(pgErr.Message, "city_id") {
					return errors.New("city does not exist")
				}

			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_crossroad") {
					return errors.New("crossroad already exists for this street and city")
				}
			}
		}
		return err
	}

	return nil
}

func (r *CrossroadRepository) FindByStreetAndCity(ctx context.Context, streetId, cityId int64) (*model.Crossroad, error) {
	var crossroad model.Crossroad
	err := r.db.QueryRow(ctx, `
        SELECT id, street_id, city_id
        FROM crossroads
        WHERE street_id = $1 AND city_id = $2
    `, streetId, cityId).Scan(&crossroad.Id, &crossroad.StreetId, &crossroad.CityId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &crossroad, nil
}
