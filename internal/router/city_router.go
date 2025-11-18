package router

import (
	"github.com/go-chi/chi/v5"
	"location-service/internal/controller"
	"location-service/internal/service"
)

func CityRoutes(cityService *service.CityService) chi.Router {
	r := chi.NewRouter()
	c := controller.NewCityController(cityService)

	r.Get("/", c.ListCities)
	r.Post("/", c.CreateCity)
	r.Get("/{id}", c.GetCity)
	r.Put("/{id}", c.UpdateCity)
	r.Delete("/{id}", c.DeleteCity)

	return r
}
