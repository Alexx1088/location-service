package service

import (
	"context"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type StreetService struct {
	repo repository.StreetRepositoryInterface
}

func NewStreetService(repo repository.StreetRepositoryInterface) *StreetService {
	return &StreetService{repo: repo}
}

func (s *StreetService) CreateStreet(ctx context.Context, street *model.Street) error {
	return s.repo.Create(ctx, street)
}

func (s *StreetService) ListStreets(ctx context.Context) ([]model.Street, error) {
	return s.repo.GetAll(ctx)
}

func (s *StreetService) DeleteStreet(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *StreetService) UpdateStreet(ctx context.Context, street *model.Street) error {
	return s.repo.Update(ctx, street)
}

func (s *StreetService) GetStreet(ctx context.Context, id int) (*model.Street, error) {
	return s.repo.GetByID(ctx, id)
}
