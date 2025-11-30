package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"location-service/internal/model"
)

type OutboxRepositoryInterface interface {
	AddEvent(ctx context.Context, tx pgx.Tx, e *model.OutboxEvent) error
}
