package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"location-service/internal/dto/city"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type CityController struct {
	service  *service.CityService
	validate *validator.Validate
}

func NewCityController(service *service.CityService) *CityController {
	return &CityController{
		service:  service,
		validate: validator.New(),
	}
}

func (c *CityController) CreateCity(w http.ResponseWriter, r *http.Request) {

	var req city.CreateCityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cities := model.City{Name: req.Name}
	if err := c.service.CreateCity(r.Context(), &cities); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			if strings.Contains(err.Error(), "already exists") {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			http.Error(w, err.Error(), http.StatusConflict)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(cities)
	if err != nil {
		return
	}
}

func (c *CityController) ListCities(w http.ResponseWriter, r *http.Request) {
	cities, err := c.service.ListCities(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		return
	}
}

func (c *CityController) GetCity(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	citi, err := c.service.GetCity(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(citi)
	if err != nil {
		return
	}
}

func (c *CityController) UpdateCity(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req city.UpdateCityRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cities := model.City{
		Id:   int64(id),
		Name: req.Name,
	}

	if err := c.service.UpdateCity(r.Context(), &cities); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		return
	}
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
