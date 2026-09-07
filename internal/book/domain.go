package book

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	Pages         int       `json:"pages"`
	ISBN          string    `json:"isbn"`
	CoverUrl      string    `json:"cover_url"`
	PublishedYear int       `json:"published_year"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
