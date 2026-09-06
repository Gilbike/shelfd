package config

import "os"

const (
	EnvPort         = "SHELFD_PORT"
	EnvEnv          = "SHELFD_ENV"
	EnvLogFormat    = "SHELFD_LOG_FORMAT"
	EnvArgonTime    = "SHELFD_ARGON_TIME"
	EnvArgonMemory  = "SHELFD_ARGON_MEMORY"
	EnvArgonThreads = "SHELFD_ARGON_THREADS"
)

type Config struct {
	Port      string
	Env       string
	LogFormat string
	Argon     ArgonConfig
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

	return Config{
		Port:      port,
		Env:       env,
		LogFormat: logFormat,
		Argon:     argonConfig,
	}
}
