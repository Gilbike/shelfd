package auth

import (
	"context"
	"net/http"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/internal/user"
)

type service interface {
	Authenticate(ctx context.Context, payload authenticatePayload) (*user.User, *Session, error)
}

type Handler struct {
	service service
}

func NewHandler(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(middlewares *middleware.Manager) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /", middlewares.RequireGuest(http.HandlerFunc(h.HandleUserAuthenticate)))

	return mux
}

func (h *Handler) HandleUserAuthenticate(w http.ResponseWriter, r *http.Request) {
	body, err := api.FromBody[authenticateRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	ip := r.RemoteAddr
	userAgent := r.Header.Get("User-Agent")

	user, session, err := h.service.Authenticate(r.Context(), authenticatePayload{
		username:  body.Username,
		password:  body.Password,
		ipAddress: ip,
		userAgent: userAgent,
	})
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	cookie := &http.Cookie{
		Name:     api.SessionCookieName,
		Value:    session.ID,
		Path:     "/",
		MaxAge:   api.SessionCookieMaxAge,
		Expires:  session.ExpiresAt,
		HttpOnly: api.SessionCookieHttpOnly,
		Secure:   api.SessionCookieSecure,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)

	api.JSON(w, http.StatusCreated, map[string]any{"user": user})
}
