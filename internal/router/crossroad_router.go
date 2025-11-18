package router

import (
	"github.com/go-chi/chi/v5"
	"location-service/internal/controller"
	"location-service/internal/service"
)

func CrossroadRoutes(crossroadService *service.CrossroadService) chi.Router {
	r := chi.NewRouter()
	c := controller.NewCrossroadController(crossroadService)

	r.Get("/", c.ListCrossroads)
	r.Post("/", c.CreateCrossroad)
	r.Put("/{id}", c.UpdateCrossroad)
	r.Delete("/{id}", c.DeleteCrossroad)
	r.Get("/{id}", c.GetCrossroad)

	return r
}
