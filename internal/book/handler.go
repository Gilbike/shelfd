package book

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Gilbike/shelfd/internal/core/api"
	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/middleware"
)

// TODO: move to config
const maxFileSize = 2 << 20 //2 MB

type service interface {
	List(ctx context.Context, filter listFilters) (*listResult, error)
	Get(ctx context.Context, id int64) (*Book, error)
	Create(ctx context.Context, payload createPayload) (*Book, error)
	Delete(ctx context.Context, id int64) error
	UploadCover(ctx context.Context, id int64, mimeType string, content []byte, urlTemplate string) error
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
	mux.HandleFunc("GET /{id}", h.HandleBookGet)
	mux.HandleFunc("PUT /{id}/cover", h.HandleBookCoverUpload)
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

func (h *Handler) HandleBookGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	book, err := h.service.Get(r.Context(), id)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusOK, book)
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

func (h *Handler) HandleBookCoverUpload(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	file, header, err := r.FormFile("cover")
	if err != nil {
		api.Error(w, api.NewApiError(api.ErrBadRequest))
		return
	}

	mimeType := header.Header.Get("Content-Type")
	if mimeType == "" {
		// TODO: make error better
		api.Error(w, api.NewApiError(api.ErrInternalServer))
		return
	}

	var content = make([]byte, header.Size)
	_, err = file.Read(content)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}
	defer file.Close()

	h.service.UploadCover(r.Context(), id, mimeType, content, "/api/v1/assets/%s")
	api.JSON(w, http.StatusCreated, map[string]any{"success": true})
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
