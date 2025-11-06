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

// CreateCity godoc
// @Summary Create a new city
// @Description Add a new city to the database
// @Tags cities
// @Accept json
// @Produce json
// @Param data body city.CreateCityRequest true "City data"
// @Success 201 {object} model.City
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "City already exists"
// @Router /cities [post]
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

// ListCities godoc
// @Summary Get all cities
// @Description Get a list of all cities
// @Tags cities
// @Produce json
// @Success 200 {array} model.City
// @Router /cities [get]
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

// GetCity godoc
// @Summary Get city by ID
// @Description Get city details by ID
// @Tags cities
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} model.City
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "City not found"
// @Router /cities/{id} [get]
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

// UpdateCity godoc
// @Summary Update city
// @Description Update city name by ID
// @Tags cities
// @Accept json
// @Produce json
// @Param id path int true "City ID"
// @Param data body city.UpdateCityRequest true "City data"
// @Success 200 {object} model.City
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "City not found"
// @Router /cities/{id} [put]
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

// DeleteCity godoc
// @Summary Delete city
// @Description Remove a city by ID
// @Tags cities
// @Param id path int true "City ID"
// @Success 204 "No content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "City not found"
// @Router /cities/{id} [delete]
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
