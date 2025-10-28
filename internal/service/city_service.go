package service

import (
	"context"
	"location-service/internal/model"
	"location-service/internal/repository"
)

type CityService struct {
	repo *repository.CityRepository
}

func NewCityService(repo *repository.CityRepository) *CityService {
	return &CityService{repo: repo}
}

func (s *CityService) CreateCity(ctx context.Context, city *model.City) error {
	return s.repo.Create(ctx, city)
}

func (s *CityService) ListCities(ctx context.Context) ([]model.City, error) {
	return s.repo.GetAll(ctx)
}

func (s *CityService) DeleteCity(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *CityService) UpdateCity(ctx context.Context, city *model.City) error {
	return s.repo.Update(ctx, city)
}

func (s *CityService) GetCity(ctx context.Context, id int) (*model.City, error) {
	return s.repo.GetByID(ctx, id)
}
