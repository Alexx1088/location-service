package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"location-service/internal/dto/street"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
)

type StreetController struct {
	service  *service.StreetService
	validate *validator.Validate
}

func NewStreetController(s *service.StreetService) *StreetController {
	return &StreetController{
		service:  s,
		validate: validator.New(),
	}
}

func (c *StreetController) CreateStreet(w http.ResponseWriter, r *http.Request) {

	var req street.CreateStreetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	streets := model.Street{
		Name: req.Name,
	}

	if err := c.service.CreateStreet(r.Context(), &streets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(streets)
	if err != nil {
		return
	}
}

func (c *StreetController) ListStreets(w http.ResponseWriter, r *http.Request) {
	streets, err := c.service.ListStreets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(streets)
	if err != nil {
		return
	}
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

	err = json.NewEncoder(w).Encode(street)
	if err != nil {
		return
	}
}

func (c *StreetController) UpdateStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req street.UpdateStreetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	streets := model.Street{
		Id:   int64(id),
		Name: req.Name,
	}

	if err := c.service.UpdateStreet(r.Context(), &streets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(streets)
	if err != nil {
		return
	}
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
