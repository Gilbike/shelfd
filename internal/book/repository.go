package book

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const authorsSeparator = ";"

var sortByAllowList = map[string]bool{
	"title": true,
}

type pagination struct {
	page     int
	pageSize int
}

type sorting struct {
	by         string
	descending bool
}

func (s sorting) Column() string {
	if !sortByAllowList[s.by] {
		return "id"
	}
	return s.by
}

func (s sorting) Direction() string {
	if s.descending {
		return "DESC"
	}
	return "ASC"
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FetchAll(ctx context.Context, p pagination, s sorting) ([]Book, int64, error) {
	const rawQuery = `
		SELECT id, title, authors, pages, isbn, cover_url, published_year, description, created_at, updated_at
		FROM books
		ORDER BY %s %s
		LIMIT ?
		OFFSET ?;
	`

	query := fmt.Sprintf(rawQuery, s.Column(), s.Direction())

	var totalCount int64
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM books").Scan(&totalCount)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to fetch book count: %w", err)
	}

	offset := (p.page - 1) * p.pageSize
	rows, err := r.db.QueryContext(ctx, query, p.pageSize, offset)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get books: %w", err)
	}
	defer rows.Close()

	var books []Book = []Book{}

	for rows.Next() {
		var authorsRaw string
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&authorsRaw,
			&book.Pages,
			&book.ISBN,
			&book.CoverUrl,
			&book.PublishedYear,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return books, totalCount, fmt.Errorf("failed to scan row to book: %w", err)
		}
		book.Authors = strings.Split(authorsRaw, authorsSeparator)
		books = append(books, book)
	}

	return books, totalCount, nil
}
