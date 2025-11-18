package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"location-service/internal/repository"
	"location-service/internal/service"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cityRepo := repository.NewCityRepository(pool)
	cityService := service.NewCityService(cityRepo)
	r.Mount("/cities", CityRoutes(cityService))

	streetRepo := repository.NewStreetRepository(pool)
	streetService := service.NewStreetService(streetRepo)
	r.Mount("/streets", StreetRoutes(streetService))

	crossroadRepo := repository.NewCrossroadRepository(pool)
	crossroadService := service.NewCrossroadService(crossroadRepo)
	r.Mount("/crossroads", CrossroadRoutes(crossroadService))

	return r
}
