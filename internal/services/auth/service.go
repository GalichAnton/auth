package auth

import (
	"github.com/GalichAnton/auth/internal/config/env"
	"github.com/GalichAnton/auth/internal/repository"
	"github.com/GalichAnton/auth/internal/services"
)

var _ services.AuthService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	tokensConfig   env.TokensConfig
}

// NewService ...
func NewService(
	userRepository repository.UserRepository,
	tokensConfig env.TokensConfig,
) *service {
	return &service{
		userRepository: userRepository,
		tokensConfig:   tokensConfig,
	}
}
