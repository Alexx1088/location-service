package controller

import (
	"encoding/json"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
)

type CityController struct {
	service *service.CityService
}

func NewCityController(s *service.CityService) *CityController {
	return &CityController{service: s}
}

func (c *CityController) Create(w http.ResponseWriter, r *http.Request) {
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

func (c *CityController) List(w http.ResponseWriter, r *http.Request) {
	cities, err := c.service.ListCities(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}

func (c *CityController) Get(w http.ResponseWriter, r *http.Request) {

}
