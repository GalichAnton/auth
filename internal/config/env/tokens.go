package env

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	// #nosec G101
	refreshTokenSecretKey = "REFRESH_TOKEN_SECRET"
	accessTokenSecretKey  = "ACCESS_TOKEN_SECRET"
	// #nosec G101
	refreshTokenExpiration = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpiration  = "ACCESS_TOKEN_EXPIRATION"
)

// TokensConfig ...
type TokensConfig interface {
	Config() *tokensConfig
}

// TokensConfig ...
type tokensConfig struct {
	RefreshSecret     string
	AccessSecret      string
	RefreshExpiration time.Duration
	AccessExpiration  time.Duration
}

// NewTokensConfig ...
func NewTokensConfig() (TokensConfig, error) {
	refreshSecretKey := os.Getenv(refreshTokenSecretKey)
	if len(refreshSecretKey) == 0 {
		return nil, errors.New("refreshSecretKey not found")
	}

	refreshExpiration := os.Getenv(refreshTokenExpiration)
	if len(refreshExpiration) == 0 {
		return nil, errors.New("refreshExpiration not found")
	}

	refreshExpSeconds, err := strconv.ParseInt(refreshExpiration, 10, 64)
	if err != nil {
		return nil, errors.New("refreshExpiration is not a valid integer")
	}

	accessSecretKey := os.Getenv(accessTokenSecretKey)
	if len(accessSecretKey) == 0 {
		return nil, errors.New("accessSecretKey not found")
	}

	accessExpiration := os.Getenv(accessTokenExpiration)
	if len(accessExpiration) == 0 {
		return nil, errors.New("accessExpiration not found")
	}

	accessExpSeconds, err := strconv.ParseInt(accessExpiration, 10, 64)
	if err != nil {
		return nil, errors.New("accessExpiration is not a valid integer")
	}

	return &tokensConfig{
		RefreshSecret:     refreshSecretKey,
		AccessSecret:      accessSecretKey,
		RefreshExpiration: time.Duration(refreshExpSeconds) * time.Second,
		AccessExpiration:  time.Duration(accessExpSeconds) * time.Second,
	}, nil
}

func (cfg *tokensConfig) Config() *tokensConfig {
	return cfg
}
