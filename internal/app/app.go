package app

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/Gilbike/shelfd/internal/config"
	"github.com/Gilbike/shelfd/internal/database"
	_ "modernc.org/sqlite"
)

type App struct {
	config config.Config
	server *http.Server
	db     *sql.DB
}

func New(cfg config.Config) (*App, error) {
	app := &App{
		config: cfg,
	}

	// setup logger
	var logHandler slog.Handler
	switch cfg.LogFormat {
	case "json":
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	case "kv-pair":
		logHandler = slog.NewTextHandler(os.Stdout, nil)
	}

	if logHandler != nil {
		logger := slog.New(logHandler)
		slog.SetDefault(logger)
	}

	if app.config.Env == "development" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	// setup database
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	app.db = db

	// wire up
	handler := newRouter(app.wire())

	// setup http server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}
	app.server = server

	return app, nil
}

func (app *App) Start() error {
	slog.Info("Server starting", slog.String("port", app.config.Port))
	err := app.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	if err := app.server.Shutdown(ctx); err != nil {
		return err
	}

	if err := app.db.Close(); err != nil {
		return err
	}

	slog.Info("Gracefully shutting down")
	return nil
}
