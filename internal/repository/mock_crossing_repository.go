package repository

import (
	"context"
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
