package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Gilbike/shelfd/internal/middleware"
)

type RouteMap map[string]http.Handler

func newRouter(routeMap RouteMap, middlewares *middleware.Manager, enableCors bool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Shelfd!")
	})

	for route, handler := range routeMap {
		pattern := fmt.Sprintf("/api/v1/%s/", route)
		prefix := strings.TrimSuffix(pattern, "/")

		mux.Handle(pattern, http.StripPrefix(prefix, handler))
	}

	if enableCors {
		return middlewares.Logging(middlewares.CORS(mux))
	}

	return middlewares.Logging(mux)
}
