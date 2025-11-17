package response

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ApiError{
		Code:    status,
		Message: message,
	})
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
