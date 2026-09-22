package config

import (
	"os"
	"strconv"
)

const (
	EnvPort          = "SHELFD_PORT"
	EnvEnv           = "SHELFD_ENV"
	EnvLogFormat     = "SHELFD_LOG_FORMAT"
	EnvArgonTime     = "SHELFD_ARGON_TIME"
	EnvArgonMemory   = "SHELFD_ARGON_MEMORY"
	EnvArgonThreads  = "SHELFD_ARGON_THREADS"
	EnvArgonSaltSize = "SHELFD_ARGON_SALT_SIZE"
	EnvArgonKeyLen   = "SHELFD_ARGON_KEY_LEN"
	EnvSessionLength = "SHELFD_SESSION_LENGTH"
)

type Config struct {
	Port          string
	Env           string
	LogFormat     string
	Argon         ArgonConfig
	SessionLength int
}

func Load() Config {
	port := os.Getenv(EnvPort)
	if port == "" {
		port = "8080"
	}

	env := os.Getenv(EnvEnv)
	if env == "" {
		env = "production"
	}

	logFormat := os.Getenv(EnvLogFormat)
	if logFormat == "" {
		logFormat = "pretty"
	}

	argonConfig := getArgonConfig()

	sessionLengthStr := os.Getenv(EnvSessionLength)
	if sessionLengthStr == "" {
		sessionLengthStr = "43200"
	}

	sessionLength, err := strconv.ParseInt(sessionLengthStr, 10, 32)
	if err != nil {
		sessionLength = 43200
	}

	return Config{
		Port:          port,
		Env:           env,
		LogFormat:     logFormat,
		Argon:         argonConfig,
		SessionLength: int(sessionLength),
	}
}
