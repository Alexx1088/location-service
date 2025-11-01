package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
)

type CrossroadController struct {
	service *service.CrossroadService
}

func NewCrossroadController(s *service.CrossroadService) *CrossroadController {
	return &CrossroadController{service: s}
}

func (c *CrossroadController) CreateCrossroad(w http.ResponseWriter, r *http.Request) {
	var crossroad model.Crossroad
	if err := json.NewDecoder(r.Body).Decode(&crossroad); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateCrossroad(r.Context(), &crossroad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(crossroad)
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

	var crossroad model.Crossroad
	if err := json.NewDecoder(r.Body).Decode(&crossroad); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	crossroad.Id = int64(id)
	if err := c.service.UpdateCrossroad(r.Context(), &crossroad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(crossroad)
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
