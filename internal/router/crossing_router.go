package router

import (
	"github.com/go-chi/chi/v5"
	"location-service/internal/controller"
	"location-service/internal/service"
)

func CrossingRoutes(crossingService *service.CrossingService) chi.Router {
	r := chi.NewRouter()
	c := controller.NewCrossingController(crossingService)

	r.Get("/", c.ListCrossings)
	r.Post("/", c.CreateCrossing)

	return r
}
