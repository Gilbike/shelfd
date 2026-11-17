package errs

import (
	"errors"
)

var (
	ErrUnauthenticated    = errors.New("unauthenticated")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidInput       = errors.New("invalid input")
	ErrNotFound           = errors.New("not found")
	ErrAlreadyExists      = errors.New("already exists")
	ErrInUse              = errors.New("resource is in use")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrExpiredSession     = errors.New("expired session")
)
