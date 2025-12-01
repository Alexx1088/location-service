package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
)

type MockCrossingRepository struct {
	mock.Mock
}

func (m *MockCrossingRepository) Create(ctx context.Context, c *model.Crossing) error {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *MockCrossingRepository) GetAll(ctx context.Context, f crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {
	args := m.Called(ctx, f)
	return args.Get(0).([]crossing.WithAddressDTO), args.Error(1)
}

func (m *MockCrossingRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	if tx, ok := args.Get(0).(pgx.Tx); ok {
		return tx, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCrossingRepository) CreateTx(
	ctx context.Context,
	tx pgx.Tx,
	crossing *model.Crossing,
) error {
	args := m.Called(ctx, tx, crossing)
	return args.Error(0)
}
