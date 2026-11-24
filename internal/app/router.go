package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/web"
)

type ResourceHandler interface {
	RegisterRoutes(mux *http.ServeMux, prefix string, middlewares *middleware.Manager)
}

type RouteMap map[string]ResourceHandler

func newRouter(routeMap RouteMap, middlewares *middleware.Manager, isDevelopment bool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Shelfd!")
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for route, handler := range routeMap {
		fullRoute := fmt.Sprintf("/api/v1/%s", route)
		handler.RegisterRoutes(mux, fullRoute, middlewares)
	}

	if isDevelopment {
		return middlewares.Logging(middlewares.CORS(mux))
	}

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	web.EmbedWeb(mux)
	return middlewares.Logging(mux)
}
