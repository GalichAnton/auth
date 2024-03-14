package repository

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/log"
	"github.com/GalichAnton/auth/internal/models/role"
	"github.com/GalichAnton/auth/internal/models/user"
)

// UserRepository - .
type UserRepository interface {
	Create(ctx context.Context, user *user.Info) (int64, error)
	Get(ctx context.Context, id int64) (*user.User, error)
	Update(ctx context.Context, id int64, info *user.Info) error
	Delete(ctx context.Context, id int64) error
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

// LogRepository - .
type LogRepository interface {
	Create(ctx context.Context, log *log.Info) error
}

// RoleRepository - .
type RoleRepository interface {
	GetAllRolePermissions(ctx context.Context) ([]role.Permission, error)
}
