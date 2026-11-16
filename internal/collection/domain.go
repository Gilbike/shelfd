package collection

import (
	"slices"
	"time"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

var validCollections = []string{"want_to_read", "currently_reading", "backlog", "completed"}

type UserBook struct {
	UserId         int64     `json:"user_id"`
	BookId         int64     `json:"book_id"`
	CollectionName string    `json:"collection"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TODO: improve error message
func (ub *UserBook) IsCollectionNameValid() error {
	if slices.Contains(validCollections, ub.CollectionName) {
		return nil
	}
	return errs.ErrInvalidInput
}
