package env

import (
	"errors"
	"os"
)

const (
	loglevel = "LOGLEVEL"
)

// LogConfig ...
type LogConfig interface {
	Level() string
}

type logConfig struct {
	level string
}

// NewLogConfig ...
func NewLogConfig() (LogConfig, error) {
	level := os.Getenv(loglevel)
	if len(level) == 0 {
		return nil, errors.New("loglevel not found")
	}

	return &logConfig{
		level: level,
	}, nil
}

func (cfg *logConfig) Level() string {
	return cfg.level
}
