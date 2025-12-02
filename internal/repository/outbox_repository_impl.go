package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"location-service/internal/model"
)

type OutboxRepository struct{}

func NewOutboxRepository() *OutboxRepository {
	return &OutboxRepository{}
}

func (r *OutboxRepository) AddEvent(ctx context.Context, tx pgx.Tx, e *model.OutboxEvent) error {
	_, err := tx.Exec(ctx,
		`INSERT INTO outbox (aggregate_type, aggregate_id, event_type, payload, processed, processing, retry_count, last_error)
         VALUES ($1, $2, $3, $4, false, false, 0, NULL)`,
		e.AggregateType, nullableInt64(e.AggregateID), e.EventType, e.Payload,
	)
	return err
}

func nullableInt64(v *int64) interface{} {
	if v == nil {
		return nil
	}
	return *v
}

func (r *OutboxRepository) GetUnprocessed(ctx context.Context, tx pgx.Tx, limit int) ([]*model.OutboxEvent, error) {
	rows, err := tx.Query(ctx,
		`SELECT id, aggregate_type, aggregate_id, event_type, payload, retry_count
         FROM outbox
         WHERE processed = false AND processing = false
         ORDER BY id
         LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*model.OutboxEvent
	for rows.Next() {
		var e model.OutboxEvent
		err := rows.Scan(&e.ID, &e.AggregateType, &e.AggregateID, &e.EventType, &e.Payload, &e.RetryCount)
		if err != nil {
			return nil, err
		}
		events = append(events, &e)
	}

	return events, nil
}

func (r *OutboxRepository) MarkProcessed(ctx context.Context, tx pgx.Tx, id int64) error {
	_, err := tx.Exec(ctx,
		`UPDATE outbox SET processed = true, processing = false WHERE id = $1`, id)
	return err
}

func (r *OutboxRepository) MarkFailed(ctx context.Context, tx pgx.Tx, id int64, lastErr string) error {
	_, err := tx.Exec(ctx,
		`UPDATE outbox SET retry_count = retry_count + 1, processing = false, last_error = $2 WHERE id = $1`,
		id, lastErr)
	return err
}
