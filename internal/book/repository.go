package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Gilbike/shelfd/internal/core/errs"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

const authorsSeparator = ";"

var sortByAllowList = map[string]bool{
	"title": true,
}

var updateFieldAllowList = map[string]bool{
	"title":          true,
	"authors":        true,
	"pages":          true,
	"isbn":           true,
	"cover_url":      true,
	"published_year": true,
	"description":    true,
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

func (r *Repository) FindById(ctx context.Context, id int64) (*Book, error) {
	const query = `
		SELECT id, title, authors, pages, isbn, cover_url, published_year, description, created_at, updated_at
		FROM books WHERE id = ?;
	`

	var book Book
	var authorsRaw string

	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&book.ID, &book.Title, &authorsRaw, &book.Pages, &book.ISBN, &book.CoverUrl, &book.PublishedYear, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to read book from database: %w", err)
	}

	book.Authors = strings.Split(authorsRaw, ";")

	return &book, nil
}

func (r *Repository) Create(ctx context.Context, book *Book) (int64, error) {
	const query = `
		INSERT INTO books (title, authors, pages, isbn, published_year, description)
		VALUES (?, ?, ?, ?, ?, ?);
	`

	authors := strings.Join(book.Authors, authorsSeparator)
	result, err := r.db.ExecContext(ctx, query, book.Title, authors, book.Pages, book.ISBN, book.PublishedYear, book.Description)
	if err != nil {
		var sqliteError *sqlite.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
				return -1, errs.ErrAlreadyExists
			}
		}
		return -1, fmt.Errorf("failed to insert book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("failed to get last id: %w", err)
	}

	return id, nil
}

func (r *Repository) Update(ctx context.Context, id int64, dirtyValues map[string]any) error {
	fields, values, err := r.sanitizeFields(dirtyValues)
	if err != nil {
		return err
	}
	values = append(values, id)

	query := fmt.Sprintf("UPDATE books SET %s WHERE id = ?", fields)
	result, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("failed to update book %d: %w", id, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM books WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

// TODO: get slices and remove hardcoded capacity (derived from: updateFieldAllowList)
func (r *Repository) sanitizeFields(fields map[string]any) (string, []any, error) {
	keys := make([]string, 0, len(fields))
	values := make([]any, 0, len(fields)+1)

	for field, val := range fields {
		if !updateFieldAllowList[field] {
			continue
		}
		keys = append(keys, fmt.Sprintf("%s = ?", field))
		values = append(values, val)
	}

	if len(keys) == 0 {
		return "", nil, errors.New("no valid fields")
	}

	return strings.Join(keys, ", "), values, nil
}
