package env

import (
	"errors"
	"os"

	"github.com/GalichAnton/auth/internal/config"
)

var _ config.PGConfig = (*PGConfig)(nil)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig struct {
	dsn string
}

func NewPGConfig() (*PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("dsn not found")
	}

	return &PGConfig{dsn: dsn}, nil
}

func (cfg PGConfig) DSN() string {
	return cfg.dsn
}
