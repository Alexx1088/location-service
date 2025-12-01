package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
)

type MockOutboxRepository struct {
	mock.Mock
}

func (m *MockOutboxRepository) AddEvent(
	ctx context.Context,
	tx pgx.Tx,
	e *model.OutboxEvent,
) error {
	args := m.Called(ctx, tx, e)
	return args.Error(0)
}
