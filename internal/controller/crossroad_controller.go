package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"location-service/internal/dto/crossroad"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type CrossroadController struct {
	service  *service.CrossroadService
	validate *validator.Validate
}

func NewCrossroadController(s *service.CrossroadService) *CrossroadController {
	return &CrossroadController{
		service:  s,
		validate: validator.New(),
	}
}

func (c *CrossroadController) CreateCrossroad(w http.ResponseWriter, r *http.Request) {

	var req crossroad.CreateCrossroadRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	crossroads := model.Crossroad{
		CityId:   req.CityId,
		StreetId: req.StreetId,
	}

	if err := c.service.CreateCrossroad(r.Context(), &crossroads); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(crossroads)
	if err != nil {
		return
	}

}

func (c *CrossroadController) ListCrossroads(w http.ResponseWriter, r *http.Request) {
	crossroads, err := c.service.ListCrossroads(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(crossroads)
	if err != nil {
		return
	}
}

func (c *CrossroadController) GetCrossroad(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	crossroad, err := c.service.GetCrossroad(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(crossroad)
	if err != nil {
		return
	}
}

func (c *CrossroadController) UpdateCrossroad(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req crossroad.UpdateStreetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := c.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	crossroads := model.Crossroad{
		Id:       int64(id),
		CityId:   *req.CityId,
		StreetId: *req.StreetId,
	}
	if err := c.service.UpdateCrossroad(r.Context(), &crossroads); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(crossroads)
	if err != nil {
		return
	}
}

func (c *CrossroadController) DeleteCrossroad(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteCrossroad(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
