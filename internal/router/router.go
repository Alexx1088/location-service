package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"location-service/internal/controller"
	"location-service/internal/repository"
	"location-service/internal/service"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ---- Cities ----
	cityRepo := repository.NewCityRepository(pool)
	cityService := service.NewCityService(cityRepo)
	cityController := controller.NewCityController(cityService)

	r.Route("/cities", func(r chi.Router) {
		r.Get("/", cityController.ListCities)
		r.Post("/", cityController.CreateCity)
		r.Get("/{id}", cityController.GetCity)
		r.Put("/{id}", cityController.UpdateCity)
		r.Delete("/{id}", cityController.DeleteCity)
	})

	// ---- Streets ----
	streetRepo := repository.NewStreetRepository(pool)
	streetService := service.NewStreetService(streetRepo)
	streetController := controller.NewStreetController(streetService)

	r.Route("/streets", func(r chi.Router) {
		r.Get("/", streetController.ListStreets)
		r.Post("/", streetController.CreateStreet)
		r.Get("/{id}", streetController.GetStreet)
		r.Put("/{id}", streetController.UpdateStreet)
		r.Delete("/{id}", streetController.DeleteStreet)
	})

	// ----- Crossroads
	crossroadRepo := repository.NewCrossroadRepository(pool)
	crossroadService := service.NewCrossroadService(crossroadRepo)
	crossroadController := controller.NewCrossroadController(crossroadService)

	r.Route("/crossroads", func(r chi.Router) {
		r.Get("/", crossroadController.ListCrossroads)
		r.Post("/", crossroadController.CreateCrossroad)
		r.Put("/{id}", crossroadController.UpdateCrossroad)
		r.Delete("/{id}", crossroadController.DeleteCrossroad)
		r.Get("/{id}", crossroadController.GetCrossroad)
	})

	return r
}
