package database

import (
	"database/sql"
	"embed"
	"log/slog"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(db *sql.DB) error {
	// disable logging
	goose.SetLogger(goose.NopLogger())

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	goose.SetBaseFS(migrations)

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	slog.Info("Migrated database")

	return nil
}
