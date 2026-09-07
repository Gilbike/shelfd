package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/Gilbike/shelfd/internal/core/errs"
	"github.com/Gilbike/shelfd/internal/user"
)

type userRepository interface {
	FindById(ctx context.Context, id int64) (*user.User, error)
	FetchPasswordHashByUsername(ctx context.Context, username string) (string, int64, error)
}

type repository interface {
	FindById(ctx context.Context, id string) (*Session, error)
	Create(ctx context.Context, session *Session) error
	Delete(ctx context.Context, id string) error
}

type passwordHasher interface {
	HashPassword(password string) ([]byte, []byte, error)
	EncodePassword(hash []byte, salt []byte) string
	VerifyPasswordHash(password string, hash string) bool
}

type authenticatePayload struct {
	username  string
	password  string
	ipAddress string
	userAgent string
}

type Service struct {
	repository repository
	hasher     passwordHasher
	userRepo   userRepository
	dummyHash  string
}

func NewService(repo repository, userRepo userRepository, hasher passwordHasher) *Service {
	dummyPass, dummySalt, err := hasher.HashPassword("timing-attack-avoiding-dummy-password")
	if err != nil {
		panic(fmt.Errorf("failed to create dummy hash: %w", err))
	}

	dummyHash := hasher.EncodePassword(dummyPass, dummySalt)

	return &Service{
		repository: repo,
		hasher:     hasher,
		userRepo:   userRepo,
		dummyHash:  dummyHash,
	}
}

func (s *Service) Authenticate(ctx context.Context, payload authenticatePayload) (*user.User, *Session, error) {
	dummy := false

	dbHash, userId, err := s.userRepo.FetchPasswordHashByUsername(ctx, payload.username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return nil, nil, err
		}
		dummy = true
	}

	var hash string
	if dummy {
		hash = s.dummyHash
	} else {
		hash = dbHash
	}

	isValid := s.hasher.VerifyPasswordHash(payload.password, hash)

	if dummy || !isValid {
		return nil, nil, errs.ErrInvalidCredentials
	}

	user, err := s.userRepo.FindById(ctx, userId)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove magic number
	sessIdBytes := make([]byte, 16)
	_, err = rand.Read(sessIdBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create session id: %w", err)
	}

	sessId := base64.RawURLEncoding.EncodeToString(sessIdBytes)

	// TODO: move expiry time to config
	expiresAt := time.Now().Add(30 * 24 * time.Hour).UTC()

	session := &Session{
		ID:        sessId,
		UserId:    user.ID,
		ExpiresAt: expiresAt,
		UserAgent: payload.userAgent,
		IpAddress: payload.ipAddress,
	}

	err = s.repository.Create(ctx, session)
	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}

func (s *Service) VerifyCookie(ctx context.Context, sessionId string) (int64, error) {
	session, err := s.repository.FindById(ctx, sessionId)
	if err != nil {
		return -1, err
	}

	if !session.IsValid() {
		return -1, errs.ErrExpiredSession
	}

	return session.UserId, nil
}

func (s *Service) RevokeSession(ctx context.Context, sessionId string) error {
	err := s.repository.Delete(ctx, sessionId)
	if err != nil {
		return err
	}

	return nil
}
