package access

import (
	"github.com/GalichAnton/auth/internal/services"
	desc "github.com/GalichAnton/auth/pkg/access_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService services.AccessService
}

// NewImplementation ...
func NewImplementation(accessService services.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
