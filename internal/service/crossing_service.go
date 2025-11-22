package service

import (
	"context"
	"errors"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CrossingService struct {
	repo          repository.CrossingRepositoryInterface
	crossroadRepo repository.CrossroadRepositoryInterface
}

func NewCrossingService(repo repository.CrossingRepositoryInterface, crossroadRepo repository.CrossroadRepositoryInterface) *CrossingService {
	return &CrossingService{
		repo:          repo,
		crossroadRepo: crossroadRepo,
	}
}

func (s *CrossingService) CreateCrossing(ctx context.Context, dto *crossing.CreateCrossingRequest) error {

	exists, err := s.crossroadRepo.Exists(ctx, dto.CrossroadID)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("crossroad does not exists")
	}

	oneCrossing := &model.Crossing{
		CrossroadID: dto.CrossroadID,
		EventTime:   dto.EventTime,
	}

	return s.repo.Create(ctx, oneCrossing)
}

func (s *CrossingService) ListCrossings(ctx context.Context, filter crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {

	return s.repo.GetAll(ctx, filter)
}
