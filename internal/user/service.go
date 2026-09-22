package user

import (
	"context"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

const (
	passwordMinLen = 8
)

var passwordLowerCaseRegex = regexp.MustCompile(`[a-z]`)
var passwordUpperCaseRegex = regexp.MustCompile(`[A-Z]`)
var passwordNumberRegex = regexp.MustCompile(`[0-9]`)

type repository interface {
	Create(ctx context.Context, user *User, hash string) (int64, error)
	FindById(ctx context.Context, id int64) (*User, error)
}

type passwordHasher interface {
	HashPassword(password string) ([]byte, []byte, error)
	EncodePassword(hash []byte, salt []byte) string
}

type userCreatePayload struct {
	Username    string
	Password    string
	DisplayName string
}

type Service struct {
	repository repository
	hasher     passwordHasher
}

func NewService(repository repository, hasher passwordHasher) *Service {
	return &Service{
		repository: repository,
		hasher:     hasher,
	}
}

func (s *Service) Create(ctx context.Context, payload userCreatePayload) (*User, error) {
	user := &User{
		Username:    payload.Username,
		DisplayName: payload.DisplayName,
	}

	userErr := user.Validate()
	passErr := s.validatePassword(payload.Password)
	err := errs.MergeValidationErrors(userErr, passErr)
	if err != nil {
		return nil, err
	}

	password, salt, err := s.hasher.HashPassword(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	encodedPass := s.hasher.EncodePassword(password, salt)

	userId, err := s.repository.Create(ctx, user, encodedPass)
	if err != nil {
		return nil, err
	}

	user.ID = userId

	return user, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	user, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) validatePassword(password string) error {
	validationErrors := errs.NewValidationError()

	passwordRunes := utf8.RuneCountInString(password)
	if passwordRunes < passwordMinLen {
		validationErrors.AddWithParams("password", errs.CodeMinLen, map[string]any{"min": passwordMinLen})
	}
	if !passwordLowerCaseRegex.MatchString(password) {
		validationErrors.Add("password", errs.CodeMustContainLower)
	}
	if !passwordUpperCaseRegex.MatchString(password) {
		validationErrors.Add("password", errs.CodeMustContainUpper)
	}
	if !passwordNumberRegex.MatchString(password) {
		validationErrors.Add("password", errs.CodeMustContainNumber)
	}

	return validationErrors.ToError()
}
