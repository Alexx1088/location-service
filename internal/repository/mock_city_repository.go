package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
)

type MockCityRepository struct {
	mock.Mock
}

func (m *MockCityRepository) Create(ctx context.Context, city *model.City) error {
	args := m.Called(ctx, city)
	return args.Error(0)
}

func (m *MockCityRepository) GetAll(ctx context.Context) ([]model.City, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.City), args.Error(1)
}

func (m *MockCityRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCityRepository) GetByID(ctx context.Context, id int) (*model.City, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.City), args.Error(1)
}

func (m *MockCityRepository) Update(ctx context.Context, city *model.City) error {
	args := m.Called(ctx, city)
	return args.Error(0)
}
