package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"location-service/internal/model"
	"location-service/internal/repository"
	"testing"
)

func TestCityService_CreateCity(t *testing.T) {
	mockRepo := new(repository.MockCityRepository)
	service := NewCityService(mockRepo)

	city := &model.City{
		Name: "TestCity",
	}

	mockRepo.On("Create", mock.Anything, city).Return(nil)

	err := service.CreateCity(context.Background(), city)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCityService_GetCity(t *testing.T) {
	mockRepo := new(repository.MockCityRepository)
	city := &model.City{Id: 1, Name: "Omsk"}

	mockRepo.On("GetByID", mock.Anything, 1).Return(city, nil)

	s := &CityService{repo: mockRepo}

	got, err := s.GetCity(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "Omsk", got.Name)
	mockRepo.AssertExpectations(t)
}

func TestCityService_GetCity_NotFound(t *testing.T) {
	mockRepo := new(repository.MockCityRepository)
	mockRepo.On("GetByID", mock.Anything, 999).Return((*model.City)(nil), errors.New("not found"))

	s := &CityService{repo: mockRepo}

	got, err := s.GetCity(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, got)
	mockRepo.AssertExpectations(t)
}
