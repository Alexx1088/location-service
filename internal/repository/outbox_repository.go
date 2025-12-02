package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"location-service/internal/model"
)

type OutboxRepositoryInterface interface {
	AddEvent(ctx context.Context, tx pgx.Tx, e *model.OutboxEvent) error
	GetUnprocessed(ctx context.Context, limit int) ([]*model.OutboxEvent, error)
	MarkProcessed(ctx context.Context, id int64) error
	MarkFailed(ctx context.Context, id int64, lastErr string) error
}
