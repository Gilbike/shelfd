package book

import (
	"errors"
	"testing"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

func TestValidateTitle(t *testing.T) {
	mock_book := Book{
		Title:   "",
		Authors: []string{},
		Pages:   0,
	}

	err := mock_book.Validate()

	var validationError errs.ValidationError
	if !errors.As(err, &validationError) {
		t.Errorf("%+v, expected validation error, got nil", mock_book)
		return
	}

	if len(validationError["title"]) == 0 {
		t.Errorf(`Title = %s, expected title error`, mock_book.Title)
	}

	if len(validationError["authors"]) == 0 {
		t.Errorf(`Authors = %v, expected authors error`, mock_book.Authors)
	}

	if len(validationError["pages"]) == 0 {
		t.Errorf(`Title = %d, expected pages error`, mock_book.Pages)
	}
}
