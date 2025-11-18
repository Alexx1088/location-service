package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
)

type MockCrossroadRepository struct {
	mock.Mock
}

func (m *MockCrossroadRepository) Create(ctx context.Context, crossroad *model.Crossroad) error {
	args := m.Called(ctx, crossroad)
	return args.Error(0)
}

func (m *MockCrossroadRepository) GetAll(ctx context.Context) ([]model.Crossroad, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Crossroad), args.Error(1)
}

func (m *MockCrossroadRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCrossroadRepository) GetByID(ctx context.Context, id int) (*model.Crossroad, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Crossroad), args.Error(1)
}

func (m *MockCrossroadRepository) Update(ctx context.Context, crossroad *model.Crossroad) error {
	args := m.Called(ctx, crossroad)
	return args.Error(0)
}

func (m *MockCrossroadRepository) FindByStreetAndCity(ctx context.Context, streetId, cityId int64) (*model.Crossroad, error) {
	args := m.Called(ctx, streetId, cityId)
	return args.Get(0).(*model.Crossroad), args.Error(1)
}
