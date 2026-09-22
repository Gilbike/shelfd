package book

import (
	"strings"
	"time"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	Pages         int       `json:"pages"`
	ISBN          *string   `json:"isbn"`
	CoverUrl      *string   `json:"cover_url"`
	PublishedYear *int      `json:"published_year"`
	Description   *string   `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (b *Book) Validate() error {
	errors := errs.NewValidationError()

	title := strings.TrimSpace(b.Title)

	if len(title) < 1 {
		errors.AddWithParams("title", errs.CodeMinLen, map[string]any{"min": 1})
	}

	if len(b.Authors) < 1 {
		errors.Add("authors", errs.CodeRequired)
	}

	if b.Pages < 1 {
		// TODO: improve error code
		errors.Add("pages", errs.CodeRequired)
	}

	return errors.ToError()
}
