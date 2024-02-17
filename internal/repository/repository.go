package repository

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/user"
)

// UserRepository - .
type UserRepository interface {
	Create(ctx context.Context, user *user.Info) (int64, error)
	Get(ctx context.Context, id int64) (*user.User, error)
	Update(ctx context.Context, id int64, info *user.Info) error
	Delete(ctx context.Context, id int64) error
}
