package shelf

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Gilbike/shelfd/internal/core/errs"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FetchBookShelfForUser(ctx context.Context, userId int64, bookId int64) (string, error) {
	const query = `
		SELECT collection
		FROM user_books
		WHERE user_id = ? AND book_id = ?;
	`

	var collection string

	row := r.db.QueryRowContext(ctx, query, userId, bookId)
	if err := row.Err(); err != nil {
		return "", fmt.Errorf("failed to get book collection for user: %w", err)
	}
	if err := row.Scan(&collection); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errs.ErrNotFound
		}
		return "", fmt.Errorf("failed to get collection by user and book: %w", err)
	}

	return collection, nil
}

func (r *Repository) FetchBookUserCount(ctx context.Context, bookId int64) (int64, error) {
	const query = `
		SELECT COUNT(user_id)
		FROM user_books
		WHERE book_id = ?;
	`

	var count int64

	row := r.db.QueryRowContext(ctx, query, bookId)
	if err := row.Scan(&count); err != nil {
		return -1, fmt.Errorf("failed to get user count for book: %w", err)
	}

	return count, nil
}

func (r *Repository) UpsertShelf(ctx context.Context, userId int64, bookId int64, collection string) error {
	const query = `
		INSERT INTO user_books (user_id, book_id, collection)
		VALUES (?, ?, ?)
		ON CONFLICT(user_id, book_id) DO UPDATE SET
		collection = excluded.collection;
	`

	_, err := r.db.ExecContext(ctx, query, userId, bookId, collection)
	if err != nil {
		var sqliteError *sqlite.Error
		if errors.As(err, &sqliteError) {
			// if the book doesnt exists (or the user, but thats a low chance)
			if sqliteError.Code() == sqlite3.SQLITE_CONSTRAINT_FOREIGNKEY {
				return errs.ErrNotFound
			}
		}
		return fmt.Errorf("failed to insert book: %w", err)
	}

	return nil
}
