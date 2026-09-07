package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/internal/user"
)

type service interface {
	Authenticate(ctx context.Context, payload authenticatePayload) (*user.User, *Session, error)
	RevokeSession(ctx context.Context, sessionId string) error
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

	mux.HandleFunc("POST /", h.HandleUserAuthenticate)
	mux.HandleFunc("POST /logout", h.HandleUserLogout)

	return middlewares.WithSession(mux)
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

	if cookie, err := r.Cookie(api.SessionCookieName); err == nil {
		_ = h.service.RevokeSession(r.Context(), cookie.Value)
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

func (h *Handler) HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(api.SessionCookieName); err == nil {
		_ = h.service.RevokeSession(r.Context(), cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     api.SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusNoContent)
}
