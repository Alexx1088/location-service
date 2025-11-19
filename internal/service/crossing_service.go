package service

import (
	"context"
	"location-service/internal/dto/crossing"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CrossingService struct {
	repo repository.CrossingRepositoryInterface
}

func NewCrossingService(repo repository.CrossingRepositoryInterface) *CrossingService {
	return &CrossingService{repo: repo}
}

func (s *CrossingService) CreateCrossing(ctx context.Context, crossing *model.Crossing) error {
	return s.repo.Create(ctx, crossing)
}

func (s *CrossingService) ListCrossings(ctx context.Context, filter crossing.FilterDTO) ([]crossing.WithAddressDTO, error) {

	return s.repo.GetAll(ctx, filter)
}
