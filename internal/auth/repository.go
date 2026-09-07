package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindById(ctx context.Context, id string) (*Session, error) {
	const query = `
		SELECT id, user_id, created_at, expires_at, user_agent, ip_address
		FROM sessions
		WHERE id = ?;
	`

	var session Session
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&session.ID, &session.UserId, &session.CreatedAt, &session.ExpiresAt, &session.UserAgent, &session.IpAddress); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find session by: %w", err)
	}

	return &session, nil
}

func (r *Repository) Create(ctx context.Context, session *Session) error {
	const query = `
		INSERT INTO sessions (id, user_id, expires_at, user_agent, ip_address)
		VALUES (?, ?, ?, ?, ?);
	`

	_, err := r.db.ExecContext(ctx, query, session.ID, session.UserId, session.ExpiresAt.Format(time.RFC3339), session.UserAgent, session.IpAddress)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	const query = `
		DELETE FROM sessions
		WHERE id = ?;
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if affected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
