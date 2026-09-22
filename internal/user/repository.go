package user

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
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindById(ctx context.Context, id int64) (*User, error) {
	const query = `
		SELECT id, username, display_name, created_at, updated_at
		FROM users
		WHERE id = ?;
	`

	var user User
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&user.ID, &user.Username, &user.DisplayName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user from id: %w", err)
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *User, hash string) (int64, error) {
	const query = `
		INSERT INTO users (username, display_name, password_hash)
		VALUES (?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, user.Username, user.DisplayName, hash)
	if err != nil {
		var sqliteError *sqlite.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
				return -1, errs.ErrAlreadyExists
			}
		}

		return -1, fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("failed to get last id: %w", err)
	}

	return id, nil
}

func (r *Repository) FetchPasswordHashByUsername(ctx context.Context, username string) (string, int64, error) {
	const query = `
		SELECT id, password_hash
		FROM users
		WHERE username = ?;
	`

	var userId int64
	var hash string
	result := r.db.QueryRowContext(ctx, query, username)
	if err := result.Scan(&userId, &hash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", -1, errs.ErrNotFound
		}
		return "", -1, fmt.Errorf("failed to get hash from username: %w", err)
	}

	return hash, userId, nil
}
