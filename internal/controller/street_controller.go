package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
)

type StreetController struct {
	service *service.StreetService
}

func NewStreetController(s *service.StreetService) *StreetController {
	return &StreetController{service: s}
}

func (c *StreetController) CreateStreet(w http.ResponseWriter, r *http.Request) {
	var street model.Street
	if err := json.NewDecoder(r.Body).Decode(&street); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateStreet(r.Context(), &street); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(street)
}

func (c *StreetController) ListCities(w http.ResponseWriter, r *http.Request) {
	streets, err := c.service.ListStreets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(streets)
}

func (c *StreetController) GetStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	street, err := c.service.GetStreet(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(street)
}

func (c *StreetController) UpdateStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var street model.Street
	if err := json.NewDecoder(r.Body).Decode(&street); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	street.Id = int64(id)
	if err := c.service.UpdateStreet(r.Context(), &street); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(street)
}

func (c *StreetController) DeleteStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteStreet(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
