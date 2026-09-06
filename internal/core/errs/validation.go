package errs

import (
	"fmt"
	"strings"
)

type ValidationCode string

const (
	CodeRequired          ValidationCode = "errors.required"
	CodeMinLen            ValidationCode = "errors.min_length"
	CodeMaxLen            ValidationCode = "errors.max_length"
	CodeInvalidCharSet    ValidationCode = "errors.invalid_charset"
	CodeMustContainLower  ValidationCode = "errors.must_contain_lower"
	CodeMustContainUpper  ValidationCode = "errors.must_contain_upper"
	CodeMustContainNumber ValidationCode = "errors.must_contain_number"
)

type FieldError struct {
	Code   ValidationCode `json:"code"`
	Params map[string]any `json:"params,omitempty"`
}

type ValidationError map[string][]FieldError

func NewValidationError() ValidationError {
	return make(ValidationError)
}

func (ve ValidationError) Error() string {
	if len(ve) == 0 {
		return ErrInvalidInput.Error()
	}

	var strs []string
	for field, errs := range ve {
		var errStr []string
		for _, err := range errs {
			errStr = append(errStr, string(err.Code))
		}

		strs = append(strs, fmt.Sprintf("%s: %s", field, errStr))
	}
	return strings.Join(strs, "; ")
}

func (ve ValidationError) Unwrap() error {
	return ErrInvalidInput
}

func (ve ValidationError) Add(field string, code ValidationCode) {
	ve.AddWithParams(field, code, nil)
}

func (ve ValidationError) AddWithParams(field string, code ValidationCode, params map[string]any) {
	ve[field] = append(ve[field], FieldError{Code: code, Params: params})
}

func (ve ValidationError) HasErrors() bool {
	return len(ve) > 0
}

func (ve ValidationError) ToError() error {
	if ve.HasErrors() {
		return ve
	}
	return nil
}
