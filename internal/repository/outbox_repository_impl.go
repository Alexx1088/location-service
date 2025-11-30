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
