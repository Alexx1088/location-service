package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
	"location-service/internal/repository"
	"time"
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

func (s *CrossingService) CreateCrossing(ctx context.Context, dto *crossing.CreateCrossingRequest) error {

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	c := &model.Crossing{
		CrossroadID: dto.CrossroadID,
		EventTime:   dto.EventTime,
	}

	err = s.repo.CreateTx(ctx, tx, c)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.New("crossroad does not exists")
		}
		return fmt.Errorf("cannot create crossing: %w", err)
	}

	payloadBytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("cannot marshal crossing payload: %w", err)
	}

	id := int64(c.ID)

	outboxEvent := &model.OutboxEvent{
		AggregateType: "crossing",
		AggregateID:   &id,
		EventType:     "crossing_created",
		Payload:       payloadBytes,
		Processed:     false,
		Processing:    false,
		RetryCount:    0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := s.outboxRepo.AddEvent(ctx, tx, outboxEvent); err != nil {
		return fmt.Errorf("cannot create outbox event: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("cannot commit transaction: %w", err)
	}

	return nil
}

func (s *CrossingService) ListCrossings(ctx context.Context, filter crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {

	return s.repo.GetAll(ctx, filter)
}
