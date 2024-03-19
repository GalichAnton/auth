package repository

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/log"
	"github.com/GalichAnton/auth/internal/models/role"
	"github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/repository/user/model"
)

// UserRepository - .
type UserRepository interface {
	Create(ctx context.Context, user *user.Info) (int64, error)
	Get(ctx context.Context, filter model.Filter) (*user.User, error)
	Update(ctx context.Context, id int64, info *user.Info) error
	Delete(ctx context.Context, id int64) error
}

// LogRepository - .
type LogRepository interface {
	Create(ctx context.Context, log *log.Info) error
}

// RoleRepository - .
type RoleRepository interface {
	GetAllRolePermissions(ctx context.Context) ([]role.Permission, error)
}
