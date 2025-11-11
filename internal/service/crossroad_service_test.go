package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
	"location-service/internal/repository"
	"testing"
)

func TestCrossroadService_CreateCrossroad(t *testing.T) {
	mockRepo := new(repository.MockCrossroadRepository)
	service := NewCrossroadService(mockRepo)

	crossroad := &model.Crossroad{
		StreetId: 1,
		CityId:   1,
	}

	mockRepo.On("FindByStreetAndCity", mock.Anything, crossroad.StreetId, crossroad.CityId).
		Return((*model.Crossroad)(nil), sql.ErrNoRows)

	mockRepo.On("Create", mock.Anything, crossroad).Return(nil)

	err := service.CreateCrossroad(context.Background(), crossroad)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCrossroadService_GetCrossroad(t *testing.T) {
	mockRepo := new(repository.MockCrossroadRepository)
	crossroad := &model.Crossroad{StreetId: 1, CityId: 1}

	mockRepo.On("GetByID", mock.Anything, 1).Return(crossroad, nil)

	s := &CrossroadService{repo: mockRepo}

	got, err := s.GetCrossroad(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), got.StreetId)
	assert.Equal(t, int64(1), got.CityId)
	mockRepo.AssertExpectations(t)
}

func TestCrossroadService_GetCrossroad_NotFound(t *testing.T) {
	mockRepo := new(repository.MockCrossroadRepository)
	mockRepo.On("GetByID", mock.Anything, 999).Return((*model.Crossroad)(nil), errors.New("not found"))

	s := &CrossroadService{repo: mockRepo}

	got, err := s.GetCrossroad(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, got)
	mockRepo.AssertExpectations(t)
}
