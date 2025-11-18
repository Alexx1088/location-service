package router

import (
	"github.com/go-chi/chi/v5"
	"location-service/internal/controller"
	"location-service/internal/service"
)

func StreetRoutes(streetService *service.StreetService) chi.Router {
	r := chi.NewRouter()
	c := controller.NewStreetController(streetService)

	r.Get("/", c.ListStreets)
	r.Post("/", c.CreateStreet)
	r.Get("/{id}", c.GetStreet)
	r.Put("/{id}", c.UpdateStreet)
	r.Delete("/{id}", c.DeleteStreet)

	return r
}
