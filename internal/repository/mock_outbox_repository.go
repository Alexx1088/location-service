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

func (m *MockOutboxRepository) GetUnprocessed(ctx context.Context, limit int) ([]*model.OutboxEvent, error) {
	args := m.Called(ctx, limit)
	return args.Get(0).([]*model.OutboxEvent), args.Error(1)
}

func (m *MockOutboxRepository) MarkProcessed(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockOutboxRepository) MarkFailed(ctx context.Context, id int64, lastErr string) error {
	args := m.Called(ctx, id, lastErr)
	return args.Error(0)
}

func (m *MockOutboxRepository) AddEvent(ctx context.Context, tx pgx.Tx, e *model.OutboxEvent) error {
	args := m.Called(ctx, tx, e)
	return args.Error(0)
}
