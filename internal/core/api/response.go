package api

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		Error(w, NewApiError(ErrInternalServer))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

func Error(w http.ResponseWriter, error *ApiError) {
	JSON(w, error.Status, error)
}
