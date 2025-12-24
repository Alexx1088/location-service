package service

import (
	"context"
	"encoding/json"
	"fmt"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CrossingService struct {
	repo          repository.CrossingRepositoryInterface
	crossroadRepo repository.CrossroadRepositoryInterface
	outboxRepo    repository.OutboxRepositoryInterface
}

func NewCrossingService(
	repo repository.CrossingRepositoryInterface,
	crossroadRepo repository.CrossroadRepositoryInterface,
	outboxRepo repository.OutboxRepositoryInterface,
) *CrossingService {
	return &CrossingService{
		repo:          repo,
		crossroadRepo: crossroadRepo,
		outboxRepo:    outboxRepo,
	}
}

func (s *CrossingService) CreateCrossing(
	ctx context.Context,
	dto *crossing.CreateCrossingRequest,
) (*model.Crossing, error) {

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	c := &model.Crossing{
		CrossroadID: dto.CrossroadID,
		EventTime:   dto.EventTime,
	}

	err = s.repo.CreateTx(ctx, tx, c)
	if err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(c)
	id := int64(c.ID)

	event := &model.OutboxEvent{
		AggregateType: "crossing",
		AggregateID:   &id,
		EventType:     "crossing_created",
		Payload:       payload,
	}

	if err = s.outboxRepo.AddEvent(ctx, tx, event); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CrossingService) ListCrossings(ctx context.Context, filter crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {

	return s.repo.GetAll(ctx, filter)
}
