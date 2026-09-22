package config

import (
	"log/slog"
	"os"
	"strconv"
)

type ArgonConfig struct {
	Time     uint32
	Memory   uint32
	Threads  uint8
	SaltSize uint32
	KeyLen   uint32
}

func getArgonConfig() ArgonConfig {
	const argonDefaultTime uint32 = 3
	const argonDefaultMemory uint32 = 19 * 1024
	const argonDefaultThread uint8 = 2
	const argonDefaultSaltSize uint32 = 16
	const argonDefaultKeyLen uint32 = 32

	argonConfig := ArgonConfig{
		Time:     argonDefaultTime,
		Memory:   argonDefaultMemory,
		Threads:  argonDefaultThread,
		SaltSize: argonDefaultSaltSize,
		KeyLen:   argonDefaultKeyLen,
	}

	timeStr := os.Getenv(EnvArgonTime)
	if timeStr != "" {
		timeVal, err := strconv.ParseUint(timeStr, 10, 32)
		if err != nil {
			slog.Error("invalid configuration value, using fallback",
				"env", EnvArgonTime,
				"error", err,
				"default", argonConfig.Time,
			)
		} else {
			argonConfig.Time = uint32(timeVal)
		}
	}

	memoryStr := os.Getenv(EnvArgonMemory)
	if memoryStr != "" {
		memoryVal, err := strconv.ParseUint(memoryStr, 10, 32)
		if err != nil {
			slog.Error("invalid configuration value, using fallback",
				"env", EnvArgonMemory,
				"error", err,
				"default", argonConfig.Memory,
			)
		} else {
			argonConfig.Memory = uint32(memoryVal)
		}
	}

	threadStr := os.Getenv(EnvArgonThreads)
	if threadStr != "" {
		threadVal, err := strconv.ParseUint(threadStr, 10, 8)
		if err != nil {
			slog.Error("invalid configuration value, using fallback",
				"env", EnvArgonThreads,
				"error", err,
				"default", argonConfig.Threads,
			)
		} else {
			argonConfig.Threads = uint8(threadVal)
		}
	}

	saltSizeStr := os.Getenv(EnvArgonSaltSize)
	if saltSizeStr != "" {
		saltSizeVal, err := strconv.ParseUint(saltSizeStr, 10, 32)
		if err != nil {
			slog.Error("invalid configuration value, using fallback",
				"env", EnvArgonSaltSize,
				"error", err,
				"default", argonConfig.SaltSize,
			)
		} else {
			argonConfig.SaltSize = uint32(saltSizeVal)
		}
	}

	keyLenStr := os.Getenv(EnvArgonKeyLen)
	if keyLenStr != "" {
		keyLenVal, err := strconv.ParseUint(keyLenStr, 10, 32)
		if err != nil {
			slog.Error("invalid configuration value, using fallback",
				"env", EnvArgonKeyLen,
				"error", err,
				"default", argonConfig.KeyLen,
			)
		} else {
			argonConfig.KeyLen = uint32(keyLenVal)
		}
	}

	return argonConfig
}
