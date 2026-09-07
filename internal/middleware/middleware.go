package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
)

type contextKey string

const userIDKey contextKey = "userId"

func UserIDFromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}

type CookieVerifier interface {
	VerifyCookie(ctx context.Context, sessionId string) (int64, error)
}

type Manager struct {
	cookieVerfier CookieVerifier
}

func New(verifier CookieVerifier) *Manager {
	return &Manager{
		cookieVerfier: verifier,
	}
}

func (mw *Manager) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		userAttrs := slog.Group("user", slog.String("ip", r.RemoteAddr))
		requestAttrs := slog.Group("request", slog.String("protocol", r.Proto), slog.String("method", r.Method), slog.String("URL", r.URL.Path))

		slog.Info("Processed request", userAttrs, requestAttrs, slog.Duration("duration", duration))
	})
}

func (mw *Manager) WithSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(api.SessionCookieName)
		if err == nil {
			id, err := mw.cookieVerfier.VerifyCookie(r.Context(), cookie.Value)
			if err == nil {
				ctx := context.WithValue(r.Context(), userIDKey, id)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     api.SessionCookieName,
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: api.SessionCookieHttpOnly,
				Secure:   api.SessionCookieSecure,
				SameSite: http.SameSiteLaxMode,
			})
		}

		next.ServeHTTP(w, r)
	})
}

func (mw *Manager) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := UserIDFromContext(r.Context()); !ok {
			api.Error(w, api.NewApiError(errs.ErrUnauthenticated))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mw *Manager) RequireGuest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := UserIDFromContext(r.Context()); ok {
			api.Error(w, api.NewApiError(errs.ErrForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}
