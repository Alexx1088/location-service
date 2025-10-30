package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"location-service/internal/controller"
	"location-service/internal/repository"
	"location-service/internal/service"
	"net/http"
)

func NewRouter(pool *pgxpool.Pool) http.Handler {

	cityRepo := repository.NewCityRepository(pool)
	cityService := service.NewCityService(cityRepo)
	cityController := controller.NewCityController(cityService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/cities", func(r chi.Router) {
		r.Get("/", cityController.ListCities)
		r.Post("/", cityController.CreateCity)
		r.Get("/{id}", cityController.GetCity)
		r.Put("/{id}", cityController.UpdateCity)
		r.Delete("/{id}", cityController.DeleteCity)
	})

	return r
}
