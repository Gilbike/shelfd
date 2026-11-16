package collection

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

func (r *Repository) UpsertBuiltIn(ctx context.Context, userId int64, bookId int64, collection string) error {
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
