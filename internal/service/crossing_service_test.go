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
	outboxRepo := new(repository.MockOutboxRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo, outboxRepo)

	req := &crossing.CreateCrossingRequest{
		CrossroadID: 1,
		EventTime:   time.Now(),
	}

	crossingRepo.
		On("BeginTx", mock.Anything).
		Return(nil, nil)

	crossingRepo.
		On("CreateTx", mock.Anything, nil, mock.Anything).
		Return(nil)

	crossroadRepo.
		On("Exists", mock.Anything, 1).
		Return(true, nil)

	outboxRepo.
		On("AddEvent", mock.Anything, nil, mock.Anything).
		Return(nil)

	created, err := svc.CreateCrossing(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, created)
}

func TestCreateCrossing_CrossroadNotFound(t *testing.T) {
	crossingRepo := new(repository.MockCrossingRepository)
	crossroadRepo := new(repository.MockCrossroadRepository)
	outboxRepo := new(repository.MockOutboxRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo, outboxRepo)

	dto := &crossing.CreateCrossingRequest{
		CrossroadID: 99,
	}

	crossroadRepo.
		On("Exists", mock.Anything, dto.CrossroadID).
		Return(false, nil)

	created, err := svc.CreateCrossing(context.Background(), dto)
	assert.Nil(t, created)
	assert.EqualError(t, err, "crossroad does not exists")
}

func TestListCrossings(t *testing.T) {
	crossingRepo := new(repository.MockCrossingRepository)
	crossroadRepo := new(repository.MockCrossroadRepository)
	outboxRepo := new(repository.MockOutboxRepository)

	svc := NewCrossingService(crossingRepo, crossroadRepo, outboxRepo)

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
