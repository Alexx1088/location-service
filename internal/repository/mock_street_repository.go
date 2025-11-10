package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
)

type MockStreetRepository struct {
	mock.Mock
}

func (m *MockStreetRepository) Create(ctx context.Context, street *model.Street) error {
	args := m.Called(ctx, street)
	return args.Error(0)
}

func (m *MockStreetRepository) GetAll(ctx context.Context) ([]model.Street, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Street), args.Error(1)
}

func (m *MockStreetRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStreetRepository) GetByID(ctx context.Context, id int) (*model.Street, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Street), args.Error(1)
}

func (m *MockStreetRepository) Update(ctx context.Context, street *model.Street) error {
	args := m.Called(ctx, street)
	return args.Error(0)
}
