package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"location-service/internal/dto/crossing"
	"location-service/internal/service"
	"net/http"
	"strconv"
	"time"
)

type CrossingController struct {
	service  *service.CrossingService
	validate *validator.Validate
}

func NewCrossingController(service *service.CrossingService) *CrossingController {
	return &CrossingController{
		service:  service,
		validate: validator.New(),
	}
}

func (c *CrossingController) CreateCrossing(w http.ResponseWriter, r *http.Request) {
	var req crossing.CreateCrossingRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.validate.Struct(req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CreateCrossing(r.Context(), &req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		return
	}
}

func (c *CrossingController) ListCrossings(w http.ResponseWriter, r *http.Request) {

	var filter crossing.FilterDTO

	query := r.URL.Query()

	if from := query.Get("from"); from != "" {
		t, err := time.Parse(time.RFC3339, from)
		if err != nil {
			http.Error(w, "invalid 'from' date format", http.StatusBadRequest)
			return
		}
		filter.From = &t
	}

	if to := query.Get("to"); to != "" {
		t, err := time.Parse(time.RFC3339, to)
		if err != nil {
			http.Error(w, "invalid 'to' date format", http.StatusBadRequest)
			return
		}
		filter.To = &t
	}

	if crossroadID := query.Get("crossroad_id"); crossroadID != "" {
		id, err := strconv.Atoi(crossroadID)
		if err != nil {
			http.Error(w, "invalid crossroad_id", http.StatusBadRequest)
			return
		}
		filter.CrossroadID = &id
	}

	limit := 10
	offset := 0

	if l := query.Get("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
	}
	if o := query.Get("offset"); o != "" {
		offset, _ = strconv.Atoi(o)
	}

	filter.Limit = limit
	filter.Offset = offset

	crossings, err := c.service.ListCrossings(r.Context(), filter)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(crossings); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
	}
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
	if err != nil {
		return
	}
}
