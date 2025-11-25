package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"location-service/internal/dto/street"
	"location-service/internal/http/response"
	"location-service/internal/model"
	"location-service/internal/service"
	"net/http"
	"strconv"
	"strings"
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
		response.JSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if err := c.validate.Struct(req); err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	streetModel := model.Street{
		Name: req.Name,
	}

	if err := c.service.CreateStreet(r.Context(), &streetModel); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			response.JSONError(w, http.StatusConflict, err.Error())
			return
		}
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, streetModel)

}

func (c *StreetController) ListStreets(w http.ResponseWriter, r *http.Request) {
	streets, err := c.service.ListStreets(r.Context())
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, streets)
}

func (c *StreetController) GetStreet(w http.ResponseWriter, r *http.Request) {

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

	streetModel, err := c.service.GetStreet(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"errorCode":    404,
				"errorMessage": fmt.Sprintf("Street with id: %d not found", id),
			})
			if err != nil {
				return
			}
			return
		}

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
	w.WriteHeader(http.StatusOK)

	response.JSON(w, http.StatusCreated, streetModel)
}

func (c *StreetController) UpdateStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req street.UpdateStreetRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if err := c.validate.Struct(req); err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	streets := model.Street{
		Id:   int64(id),
		Name: req.Name,
	}

	if err := c.service.UpdateStreet(r.Context(), &streets); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			response.JSONError(w, http.StatusConflict, err.Error())
			return
		}
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, streets)
}

func (c *StreetController) DeleteStreet(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.service.DeleteStreet(r.Context(), id); err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
