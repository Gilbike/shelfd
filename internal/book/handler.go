package book

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/middleware"
)

type service interface {
	List(ctx context.Context, filter listFilters) (*listResult, error)
	Create(ctx context.Context, payload createPayload) (*Book, error)
	Delete(ctx context.Context, id int64) error
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

	mux.HandleFunc("GET /", h.HandleBookList)
	mux.HandleFunc("POST /", h.HandleBookCreate)
	mux.HandleFunc("DELETE /{id}", h.HandleBookDelete)

	return middlewares.WithSession(middlewares.RequireAuth(mux))
}

func (h *Handler) HandleBookList(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	result, err := h.service.List(r.Context(), listFilters{page: page})
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusOK, listResponse{
		CurrentPage: page,
		PerPage:     result.PageSize,
		TotalPages:  result.TotalPages,
		TotalBooks:  result.TotalCount,
		Data:        result.Books,
	})
}

func (h *Handler) HandleBookCreate(w http.ResponseWriter, r *http.Request) {
	request, err := api.FromBody[createRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	book, err := h.service.Create(r.Context(), createPayload{
		title:         request.Title,
		authors:       request.Authors,
		pages:         request.Pages,
		isbn:          request.ISBN,
		publishedYear: request.PublishedYear,
		description:   request.Description,
	})
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusCreated, book)
}

func (h *Handler) HandleBookDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusOK, map[string]any{"success": true})
}
