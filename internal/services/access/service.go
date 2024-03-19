package auth

import (
	"github.com/GalichAnton/auth/internal/config/env"
	"github.com/GalichAnton/auth/internal/repository"
	"github.com/GalichAnton/auth/internal/services"
)

const (
	authPrefix = "Bearer "
)

var _ services.AccessService = (*service)(nil)

type service struct {
	roleRepository repository.RoleRepository
	tokensConfig   env.TokensConfig
}

// NewService ...
func NewService(
	roleRepository repository.RoleRepository,
	tokensConfig env.TokensConfig,
) *service {
	return &service{
		roleRepository: roleRepository,
		tokensConfig:   tokensConfig,
	}
}
