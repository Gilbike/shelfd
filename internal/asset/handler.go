package asset

import (
	"fmt"
	"net/http"
	"os"

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

	content, err := os.ReadFile(fmt.Sprintf("data/%s", filename))
	if err != nil {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
