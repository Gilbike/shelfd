package collection

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/middleware"
)

type service interface {
	SetBookCollection(ctx context.Context, payload *UserBook) error
}

type Handler struct {
	service service
}

func NewHandler(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, pattern string, middlewares *middleware.Manager) {
	mux.Handle("PUT "+pattern+"/", middlewares.HelperSecure(h.HandleBookCollectionSet))
}

// TODO: check for book existense
func (h *Handler) HandleBookCollectionSet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	bookId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || bookId <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	userId, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		api.Error(w, api.NewApiError(errs.ErrUnauthenticated))
		return
	}

	request, err := api.FromBody[collectionSetRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	userBook := &UserBook{
		UserId:         userId,
		BookId:         bookId,
		CollectionName: request.Collection,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = h.service.SetBookCollection(r.Context(), userBook)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	api.JSON(w, http.StatusCreated, map[string]any{"ok": true})
}
