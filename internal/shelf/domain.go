package shelf

import (
	"slices"
	"time"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

var validShelfs = []string{"want_to_read", "currently_reading", "backlog", "completed"}

type UserBook struct {
	UserId    int64     `json:"user_id"`
	BookId    int64     `json:"book_id"`
	ShelfName string    `json:"shelf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TODO: improve error message
func (ub *UserBook) IsCollectionNameValid() error {
	if slices.Contains(validShelfs, ub.ShelfName) {
		return nil
	}
	return errs.ErrInvalidInput
}
