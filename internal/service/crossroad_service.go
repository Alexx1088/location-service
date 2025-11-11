package service

import (
	"context"
	"database/sql"
	"fmt"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CrossroadService struct {
	repo repository.CrossroadRepositoryInterface
}

func NewCrossroadService(repo repository.CrossroadRepositoryInterface) *CrossroadService {
	return &CrossroadService{repo: repo}
}

func (s *CrossroadService) CreateCrossroad(ctx context.Context, crossroad *model.Crossroad) error {

	existing, err := s.repo.FindByStreetAndCity(ctx, crossroad.StreetId, crossroad.CityId)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existing != nil {
		return fmt.Errorf("crossroad already exists for this street and city")
	}

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
