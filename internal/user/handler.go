package user

import (
	"context"
	"net/http"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/middleware"
)

type service interface {
	Create(ctx context.Context, request userCreatePayload) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
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

func (h *Handler) RegisterRoutes(mux *http.ServeMux, pattern string, middlewares *middleware.Manager) {
	mux.HandleFunc("POST "+pattern+"/", h.HandleUserCreate)
	mux.Handle("GET "+pattern+"/me", middlewares.HelperSecure(h.HandleUserCurrent))
}

func (h *Handler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := api.FromBody[createRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	user, err := h.service.Create(r.Context(), userCreatePayload(body))
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusCreated, user)
}

func (h *Handler) HandleUserCurrent(w http.ResponseWriter, r *http.Request) {
	userId, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		api.Error(w, api.NewApiError(errs.ErrUnauthenticated))
		return
	}

	user, err := h.service.GetByID(r.Context(), userId)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusOK, user)
}
