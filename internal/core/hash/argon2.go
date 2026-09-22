package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Gilbike/shelfd/internal/config"
	"golang.org/x/crypto/argon2"
)

type Argon2Hasher struct {
	config config.ArgonConfig
}

func NewArgon2Hasher(config config.ArgonConfig) *Argon2Hasher {
	return &Argon2Hasher{
		config: config,
	}
}

func (hasher *Argon2Hasher) HashPassword(password string) ([]byte, []byte, error) {
	salt, err := generateSalt(hasher.config.SaltSize)
	if err != nil {
		return nil, nil, err
	}

	return argon2.IDKey([]byte(password), salt, hasher.config.Time, hasher.config.Memory, hasher.config.Threads, hasher.config.KeyLen), salt, nil
}

func (hasher *Argon2Hasher) EncodePassword(hash []byte, salt []byte) string {
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		hasher.config.Memory,
		hasher.config.Time,
		hasher.config.Threads,
		b64Salt,
		b64Hash,
	)

	return encoded
}

func (hasher *Argon2Hasher) VerifyPasswordHash(password string, hash string) bool {
	parts := strings.Split(hash, "$")

	if len(parts) < 6 {
		slog.Error("Invalid hash format stored in database")
		return false
	}

	if parts[1] != "argon2id" {
		slog.Error("Invalid password hash algorithm", slog.String("expected", "argon2id"), slog.String("received", parts[1]))
		return false
	}
	var hashVer int
	fmt.Sscanf(parts[2], "v=%d", &hashVer)
	if hashVer != 19 {
		slog.Error("Invalid password hash version", slog.Int("expected", 19), slog.Int("received", hashVer))
		return false
	}

	var time, memory uint32
	var thread uint8
	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &thread)

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		slog.Error("Failed to retreive salt", "error", err)
		return false
	}
	dbHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		slog.Error("Failed to retreive password", "error", err)
		return false
	}

	reqHash := argon2.IDKey([]byte(password), salt, time, memory, thread, uint32(len(dbHash)))

	return subtle.ConstantTimeCompare(reqHash, dbHash) == 1
}

func generateSalt(saltSize uint32) ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("salt generation failed: %w", err)
	}
	return salt, nil
}
