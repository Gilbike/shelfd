package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

var (
	ErrInternalServer = errors.New("internal server")
	ErrBadRequest     = errors.New("bad request")
)

type ApiErrorDetails map[string]any

type ApiError struct {
	Status  int             `json:"-"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Details ApiErrorDetails `json:"details,omitempty"`
}

func NewApiError(err error) *ApiError {
	apiError := &ApiError{}

	switch {
	// validation error
	case errors.Is(err, errs.ErrInvalidInput):
		var validationError errs.ValidationError
		if errors.As(err, &validationError) {
			// copy to details
			details := make(ApiErrorDetails, len(validationError))
			for k, v := range validationError {
				details[k] = v
			}

			apiError.Details = details
		}
		apiError.Status = http.StatusUnprocessableEntity
		apiError.Code = "INVALID_INPUT"
		apiError.Message = "One or more fields contain invalid input"
	case errors.Is(err, ErrBadRequest):
		apiError.Status = http.StatusBadRequest
		apiError.Code = "BAD_REQUEST"
		apiError.Message = "Request does not match expected schema"

	// authentication errors
	case errors.Is(err, errs.ErrUnauthenticated):
		apiError.Status = http.StatusUnauthorized
		apiError.Code = "UNAUTHENTICATED"
		apiError.Message = "Cannot verify user identity"
	case errors.Is(err, errs.ErrForbidden):
		apiError.Status = http.StatusForbidden
		apiError.Code = "FORBIDDEN"
		apiError.Message = "You do not have permission to perform this action"

	// resource errors
	case errors.Is(err, errs.ErrNotFound):
		apiError.Status = http.StatusNotFound
		apiError.Code = "RESOURCE_NOT_FOUND"
		apiError.Message = "Requested resource cannot be found"
	case errors.Is(err, errs.ErrAlreadyExists):
		apiError.Status = http.StatusConflict
		apiError.Code = "RESOURCE_ALREADY_EXISTS"
		apiError.Message = "Resource conflicts with already created resource"

	// session errors
	case errors.Is(err, errs.ErrInvalidCredentials):
		apiError.Status = http.StatusUnauthorized
		apiError.Code = "INVALID_CREDENTIALS"
		apiError.Message = "Invalid username or password"
	default:
		slog.Error("Failed to server request", "error", err)
		apiError.Status = http.StatusInternalServerError
		apiError.Code = "INTERNAL_SERVER_ERROR"
		apiError.Message = "Unexpected error happened"
	}

	return apiError
}
