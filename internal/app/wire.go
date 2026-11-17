package app

import (
	"github.com/Gilbike/shelfd/internal/asset"
	"github.com/Gilbike/shelfd/internal/auth"
	"github.com/Gilbike/shelfd/internal/book"
	"github.com/Gilbike/shelfd/internal/collection"
	"github.com/Gilbike/shelfd/internal/core/hash"
	"github.com/Gilbike/shelfd/internal/middleware"
	"github.com/Gilbike/shelfd/internal/user"
)

func (app *App) wire() (RouteMap, *middleware.Manager) {
	argon2hasher := hash.NewArgon2Hasher(app.config.Argon)

	userRepo := user.NewRepository(app.db)
	authRepo := auth.NewRepository(app.db)
	bookRepo := book.NewRepository(app.db)
	collectionRepo := collection.NewRepository(app.db)

	userService := user.NewService(userRepo, argon2hasher)
	authService := auth.NewService(authRepo, userRepo, argon2hasher, app.config.SessionLength)

	collectionService := collection.NewService(collectionRepo)
	bookService := book.NewService(bookRepo, collectionService)

	middlewares := middleware.New(authService)

	userHandler := user.NewHandler(userService, authService)
	authHandler := auth.NewHandler(authService)
	bookHandler := book.NewHandler(bookService)
	collectionHandler := collection.NewHandler(collectionService)

	assetHandler := asset.NewHandler()

	return RouteMap{
		"users":      userHandler,
		"auth":       authHandler,
		"books":      bookHandler,
		"books/{id}": collectionHandler,
		"assets":     assetHandler,
	}, middlewares
}
