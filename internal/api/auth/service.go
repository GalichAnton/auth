package auth

import (
	"github.com/GalichAnton/auth/internal/services"
	desc "github.com/GalichAnton/auth/pkg/auth_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService services.AuthService
}

// NewImplementation ...
func NewImplementation(authService services.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
