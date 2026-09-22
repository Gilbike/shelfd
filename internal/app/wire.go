package app

import (
	"github.com/Gilbike/shelfd/internal/auth"
	"github.com/Gilbike/shelfd/internal/book"
	"github.com/Gilbike/shelfd/internal/core/hash"
	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/internal/user"
)

func (app *App) wire() (RouteMap, *middleware.Manager) {
	argon2hasher := hash.NewArgon2Hasher(app.config.Argon)

	userRepo := user.NewRepository(app.db)
	authRepo := auth.NewRepository(app.db)
	bookRepo := book.NewRepository(app.db)

	userService := user.NewService(userRepo, argon2hasher)
	authService := auth.NewService(authRepo, userRepo, argon2hasher, app.config.SessionLength)
	bookService := book.NewService(bookRepo)

	middlewares := middleware.New(authService)

	userHandler := user.NewHandler(userService, authService)
	authHandler := auth.NewHandler(authService)
	bookHandler := book.NewHandler(bookService)

	return RouteMap{
		"users": userHandler.RegisterRoutes(middlewares),
		"auth":  authHandler.RegisterRoutes(middlewares),
		"books": bookHandler.RegisterRoutes(middlewares),
	}, middlewares
}
