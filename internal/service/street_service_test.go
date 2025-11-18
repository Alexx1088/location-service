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

func TestStreetService_CreateStreet(t *testing.T) {
	mockRepo := new(repository.MockStreetRepository)
	service := NewStreetService(mockRepo)

	street := &model.Street{
		Name: "Test Street",
	}

	mockRepo.On("Create", mock.Anything, street).Return(nil)

	err := service.CreateStreet(context.Background(), street)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStreetService_GetStreet(t *testing.T) {
	mockRepo := new(repository.MockStreetRepository)
	street := &model.Street{Id: 1, Name: "Test Street"}

	mockRepo.On("GetByID", mock.Anything, 1).Return(street, nil)

	s := &StreetService{repo: mockRepo}

	got, err := s.GetStreet(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "Test Street", got.Name)
	mockRepo.AssertExpectations(t)
}

func TestStreetService_GetStreet_NotFound(t *testing.T) {
	mockRepo := new(repository.MockStreetRepository)
	mockRepo.On("GetByID", mock.Anything, 999).Return((*model.Street)(nil), errors.New("not found"))

	s := &StreetService{repo: mockRepo}

	got, err := s.GetStreet(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, got)
	mockRepo.AssertExpectations(t)
}
