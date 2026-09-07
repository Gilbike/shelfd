package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

const dbFileName = "app.db"

var dbConnectParams = []string{
	"_pragma=journal_mode(WAL)",
	"_pragma=synchronous(NORMAL)",
	"_pragma=busy_timeout(5000)",
	"_pragma=foreign_keys(ON)",
	"_time_format=sqlite",
	"_texttotime=true",
}

func Connect() (*sql.DB, error) {
	connParams := strings.Join(dbConnectParams, "&")
	connString := fmt.Sprintf("file:%s?%s", dbFileName, connParams)

	db, err := sql.Open("sqlite", connString)
	if err != nil {
		slog.Error("Failed to open connection to database", "error", err)
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(4)

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
