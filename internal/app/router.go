package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/web"
)

type RouteMap map[string]http.Handler

func newRouter(routeMap RouteMap, middlewares *middleware.Manager, isDevelopment bool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Shelfd!")
	})

	for route, handler := range routeMap {
		pattern := fmt.Sprintf("/api/v1/%s/", route)
		prefix := strings.TrimSuffix(pattern, "/")

		mux.Handle(pattern, http.StripPrefix(prefix, handler))
	}

	if isDevelopment {
		return middlewares.Logging(middlewares.CORS(mux))
	}

	web.EmbedWeb(mux)
	return middlewares.Logging(mux)
}
