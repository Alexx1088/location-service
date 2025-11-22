package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"location-service/internal/dto/crossing"
	"location-service/internal/repository"
	"testing"
	"time"
)

func TestCreateCrossing_Success(t *testing.T) {
	crossingRepo := new(repository.MockCrossingRepository)
	crossroadRepo := new(repository.MockCrossroadRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo)

	eventTime, _ := time.Parse(time.RFC3339, "2025-11-12T10:00:00Z")

	dto := &crossing.CreateCrossingRequest{
		CrossroadID: 10,
		EventTime:   eventTime,
	}

	crossroadRepo.
		On("Exists", mock.Anything, dto.CrossroadID).
		Return(true, nil)

	crossingRepo.
		On("Create", mock.Anything, mock.AnythingOfType("*model.Crossing")).
		Return(nil)

	err := svc.CreateCrossing(context.Background(), dto)
	assert.NoError(t, err)

	crossroadRepo.AssertExpectations(t)
	crossingRepo.AssertExpectations(t)
}

func TestCreateCrossing_CrossroadNotFound(t *testing.T) {
	crossingRepo := new(repository.MockCrossingRepository)
	crossroadRepo := new(repository.MockCrossroadRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo)

	dto := &crossing.CreateCrossingRequest{
		CrossroadID: 99,
	}

	crossroadRepo.
		On("Exists", mock.Anything, dto.CrossroadID).
		Return(false, nil)

	err := svc.CreateCrossing(context.Background(), dto)
	assert.EqualError(t, err, "crossroad does not exists")
}

func TestListCrossings(t *testing.T) {
	crossingRepo := new(repository.MockCrossingRepository)
	crossroadRepo := new(repository.MockCrossroadRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo)

	filter := crossing.FilterDTO{}

	expected := []crossing.WithAddressDTO{
		{ID: 1, Street: "Abay"},
		{ID: 2, Street: "Satpayev"},
	}

	crossingRepo.
		On("GetAll", mock.Anything, filter).
		Return(expected, nil)

	res, err := svc.ListCrossings(context.Background(), filter)

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
