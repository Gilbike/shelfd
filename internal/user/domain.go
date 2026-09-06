package user

import (
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Gilbike/shelfd/internal/core/errs"
)

const (
	usernameMinLen    = 5
	usernameMaxLen    = 25
	displayNameMinLen = 5
	displayNameMaxLen = 25
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	errors := errs.NewValidationError()

	username := strings.TrimSpace(u.Username)
	displayName := strings.TrimSpace(u.DisplayName)

	// validate username
	usernameRunes := utf8.RuneCountInString(username)
	if username == "" {
		errors.Add("username", errs.CodeRequired)
	}
	if usernameRunes < usernameMinLen {
		errors.AddWithParams("username", errs.CodeMinLen, map[string]any{"min": usernameMinLen})
	}
	if usernameRunes > usernameMaxLen {
		errors.AddWithParams("username", errs.CodeMaxLen, map[string]any{"max": usernameMaxLen})
	}
	if !usernameRegex.MatchString(username) {
		errors.Add("username", errs.CodeInvalidCharSet)
	}

	// validate display name
	displayRunes := utf8.RuneCountInString(displayName)
	if displayName == "" {
		errors.Add("display_name", errs.CodeRequired)
	}
	if displayRunes < displayNameMinLen {
		errors.AddWithParams("display_name", errs.CodeMinLen, map[string]any{"min": displayNameMinLen})
	}
	if displayRunes > displayNameMaxLen {
		errors.AddWithParams("display_name", errs.CodeMaxLen, map[string]any{"max": displayNameMaxLen})
	}

	return errors.ToError()
}
