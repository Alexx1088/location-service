package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"location-service/internal/dto/crossroad"
	"location-service/internal/http/response"
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
		if strings.Contains(err.Error(), "street_id") {
			response.JSONError(w, http.StatusBadRequest, "invalid street_id")
			return
		}
		if strings.Contains(err.Error(), "city_id") {
			response.JSONError(w, http.StatusBadRequest, "invalid city_id")
			return
		}

		response.JSONError(w, http.StatusBadRequest, "invalid request body")
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
		w.Header().Set("Content-Type", "application/json")

		switch {
		case strings.Contains(err.Error(), "does not exist"):
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"errorCode":    400,
				"errorMessage": err.Error(),
			})
			if err != nil {
				return
			}
			return

		case strings.Contains(err.Error(), "already exists"):
			w.WriteHeader(http.StatusConflict)
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"errorCode":    409,
				"errorMessage": err.Error(),
			})
			if err != nil {
				return
			}
			return

		default:
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"errorCode":    500,
				"errorMessage": "Internal Server Error",
			})
			if err != nil {
				return
			}
			return
		}
	}
	response.JSON(w, http.StatusCreated, crossroads)
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

	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    400,
			"errorMessage": "Invalid id",
		})
		if err != nil {
			return
		}
		return
	}
	crossroad, err := c.service.GetCrossroad(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"errorCode":    404,
				"errorMessage": fmt.Sprintf("Crossroad with id: %d not found", id),
			})
			if err != nil {
				return
			}
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    500,
			"errorMessage": "Internal Server Error",
		})
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
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
		if strings.Contains(err.Error(), "already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
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
