package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
)

type CityController struct {
	service *service.CityService
}

func NewCityController(s *service.CityService) *CityController {
	return &CityController{service: s}
}

func (c *CityController) CreateCity(w http.ResponseWriter, r *http.Request) {
	var city model.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateCity(r.Context(), &city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(city)
}

func (c *CityController) ListCities(w http.ResponseWriter, r *http.Request) {
	cities, err := c.service.ListCities(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}

func (c *CityController) GetCity(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	city, err := c.service.GetCity(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(city)
}

func (c *CityController) UpdateCity(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var city model.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	city.Id = int64(id)
	if err := c.service.UpdateCity(r.Context(), &city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(city)
}

func (c *CityController) DeleteCity(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteCity(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
