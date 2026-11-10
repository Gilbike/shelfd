package asset

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/middleware"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(middlewares *middleware.Manager) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{name}", h.HandleAssetGet)

	return middlewares.WithSession(middlewares.RequireAuth(mux))
}

func (h *Handler) HandleAssetGet(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("name")

	cleanName := filepath.Base(filepath.Clean(filename))
	dataFS := os.DirFS("data")
	content, err := fs.ReadFile(dataFS, cleanName)
	if err != nil {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Disposition", "attachment; filename="+cleanName)
	w.WriteHeader(http.StatusOK)
	//nolint:gosec // disable G705 // Safety: path is secured, and headers set for safety
	w.Write(content)
}
