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
const maxFileSize = 2 << 20 // 2 MB

type service interface {
	List(ctx context.Context, filter listFilters) (*listResult, error)
	ListForUser(ctx context.Context, userId int64) (*listForUserResponse, error)
	Get(ctx context.Context, userId int64, bookId int64) (*Book, error)
	Create(ctx context.Context, payload createPayload) (*Book, error)
	Update(ctx context.Context, id int64, payload updatePayload) error
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

func (h *Handler) RegisterRoutes(mux *http.ServeMux, pattern string, middlewares *middleware.Manager) {
	mux.Handle("GET "+pattern, middlewares.HelperSecure(h.HandleBookList))
	mux.Handle("GET "+pattern+"/me", middlewares.HelperSecure(h.HandleBookListForUser))
	mux.Handle("POST "+pattern, middlewares.HelperSecure(h.HandleBookCreate))
	mux.Handle("GET "+pattern+"/{id}", middlewares.HelperSecure(h.HandleBookGet))
	mux.Handle("PATCH "+pattern+"/{id}", middlewares.HelperSecure(h.HandleBookUpdate))
	mux.Handle("PUT "+pattern+"/{id}/cover", middlewares.HelperSecure(h.HandleBookCoverUpload)) // TODO: probably should change to PATCH
	mux.Handle("DELETE "+pattern+"/{id}", middlewares.HelperSecure(h.HandleBookDelete))
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

func (h *Handler) HandleBookListForUser(w http.ResponseWriter, r *http.Request) {
	userId, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		api.Error(w, api.NewApiError(errs.ErrUnauthenticated))
		return
	}

	result, err := h.service.ListForUser(r.Context(), userId)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	api.JSON(w, http.StatusOK, result)
}

func (h *Handler) HandleBookGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	userId, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		api.Error(w, api.NewApiError(errs.ErrUnauthenticated))
		return
	}

	book, err := h.service.Get(r.Context(), userId, id)
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

func (h *Handler) HandleBookUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, api.NewApiError(errs.ErrNotFound))
		return
	}

	request, err := api.FromBody[updateRequest](r.Body)
	if err != nil {
		api.Error(w, api.NewApiError(err))
		return
	}

	err = h.service.Update(r.Context(), id, updatePayload{
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

	api.JSON(w, http.StatusOK, map[string]any{"success": true})
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

	err = h.service.UploadCover(r.Context(), id, mimeType, content, "/api/v1/assets/%s")
	if err != nil {
		api.Error(w, api.NewApiError(err))
	}

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
