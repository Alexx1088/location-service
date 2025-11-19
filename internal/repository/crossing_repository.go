package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CrossingRepositoryInterface interface {
	Create(ctx context.Context, crossing *model.Crossing) error
	GetAll(ctx context.Context, dto crossing.FilterDTO) ([]crossing.WithAddressDTO, error)
}

type CrossingRepository struct {
	db *pgxpool.Pool
}

func NewCrossingRepository(db *pgxpool.Pool) *CrossingRepository {
	return &CrossingRepository{db: db}
}

func (r *CrossingRepository) Create(ctx context.Context, crossing *model.Crossing) error {
	query := `INSERT INTO crossings (crossroad_id, event_time) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(ctx, query, crossing.CrossroadID, crossing.EventTime).Scan(&crossing.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				if strings.Contains(pgErr.Message, "crossroad_id") {
					return errors.New("crossroad does not exist")
				}
			}
		}
		return err
	}
	return nil
}

func (r *CrossingRepository) GetAll(ctx context.Context, filter crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {

	baseQuery := `
		SELECT
c.id,
c.crossroad_id,
c.event_time,
st.name AS street,
ci.name AS city
FROM crossings  c
		JOIN crossroads cr ON cr.id = c.crossroad_id
		JOIN cities ci ON cr.city_id = ci.id
        JOIN streets st ON cr.street_id = st.id
		WHERE 1 = 1
		`
	args := []any{}
	argIndex := 1

	if filter.From != nil {
		baseQuery += fmt.Sprintf(" AND c.event_time >= $%d", argIndex)
		args = append(args, *filter.From)
		argIndex++
	}

	if filter.To != nil {
		baseQuery += fmt.Sprintf(" AND c.event_time <= $%d", argIndex)
		args = append(args, *filter.To)
		argIndex++
	}

	if filter.CrossroadID != nil {
		baseQuery += fmt.Sprintf(" AND c.crossroad_id = $%d", argIndex)
		args = append(args, *filter.CrossroadID)
		argIndex++
	}

	baseQuery += fmt.Sprintf(" ORDER BY c.event_time DESC LIMIT $%d OFFSET $%d",
		argIndex, argIndex+1,
	)
	args = append(args, filter.Limit, filter.Offset)
	argIndex += 2

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []crossing.WithAddressDTO

	for rows.Next() {
		var item crossing.WithAddressDTO

		err := rows.Scan(
			&item.ID,
			&item.CrossroadID,
			&item.EventTime,
			&item.Street,
			&item.City,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	return results, nil
}
