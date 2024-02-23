package user

import (
	"github.com/GalichAnton/auth/internal/services"
	desc "github.com/GalichAnton/auth/pkg/user_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService services.UserService
}

// NewImplementation ...
func NewImplementation(userService services.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
