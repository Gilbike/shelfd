package user

import (
	"context"
	"net/http"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/middleware"
)

type service interface {
	Create(ctx context.Context, request userCreatePayload) (*User, error)
}

type Handler struct {
	service  service
	verifier middleware.CookieVerifier
}

func NewHandler(service service, verifier middleware.CookieVerifier) *Handler {
	return &Handler{
		service:  service,
		verifier: verifier,
	}
}

func (h *Handler) RegisterRoutes(middlewares *middleware.Manager) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.HandleUserCreate)

	return mux
}

func (h *Handler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := api.FromBody[createRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	user, err := h.service.Create(r.Context(), userCreatePayload{
		Username:    body.Username,
		Password:    body.Password,
		DisplayName: body.DisplayName,
	})
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusCreated, user)
}
