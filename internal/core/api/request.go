package api

import (
	"encoding/json"
	"io"
	"log/slog"
)

func FromBody[T any](body io.Reader) (T, error) {
	var target T
	err := json.NewDecoder(body).Decode(&target)
	if err != nil {
		// TODO: might remove log
		slog.Error("Failed to decode request", "error", err)
		var empty T
		return empty, ErrBadRequest
	}
	return target, nil
}
