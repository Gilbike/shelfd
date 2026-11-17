package collection

import (
	"errors"
	"testing"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

func TestCollectionNameValidation(t *testing.T) {
	mock := UserBook{
		UserId:         1,
		BookId:         2,
		CollectionName: "invalid_name",
	}

	err := mock.IsCollectionNameValid()

	if !errors.Is(err, errs.ErrInvalidInput) {
		t.Errorf(`CollectionName = %+v, expected invalid input error`, mock.CollectionName)
	}
}
