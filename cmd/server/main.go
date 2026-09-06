package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gilbike/shelfd/internal/app"
	"github.com/Gilbike/shelfd/internal/config"
)

func main() {
	cfg := config.Load()
	app, err := app.New(cfg)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

	go func() {
		if err := app.Start(); err != nil {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		slog.Error("Shutdown failed, forcing exit", "error", err)
		os.Exit(1)
	}
}
