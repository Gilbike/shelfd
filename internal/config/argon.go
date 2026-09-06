package config

import (
	"log/slog"
	"os"
	"strconv"
)

type ArgonConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
}

func getArgonConfig() ArgonConfig {
	const argonDefaultTime uint32 = 3
	const argonDefaultMemory uint32 = 19 * 1024
	const argonDefaultThread uint8 = 2

	argonConfig := ArgonConfig{
		Time:    argonDefaultTime,
		Memory:  argonDefaultMemory,
		Threads: argonDefaultThread,
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

	return argonConfig
}
