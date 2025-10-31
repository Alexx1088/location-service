package service

import (
	"context"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CrossroadService struct {
	repo *repository.CrossroadRepository
}

func NewCrossroadService(repo *repository.CrossroadRepository) *CrossroadService {
	return &CrossroadService{repo: repo}
}

func (s *CrossroadService) CreateCrossroad(ctx context.Context, crossroad *model.Crossroad) error {
	return s.repo.Create(ctx, crossroad)
}

func (s *CrossroadService) ListCrossroads(ctx context.Context) ([]model.Crossroad, error) {
	return s.repo.GetAll(ctx)
}

func (s *CrossroadService) DeleteCrossroad(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *CrossroadService) UpdateCrossroad(ctx context.Context, crossroad *model.Crossroad) error {
	return s.repo.Update(ctx, crossroad)
}

func (s *CrossroadService) GetCrossroad(ctx context.Context, id int) (*model.Crossroad, error) {
	return s.repo.GetByID(ctx, id)
}
