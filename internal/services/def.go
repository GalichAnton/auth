package services

import (
	"context"

	modelService "github.com/GalichAnton/auth/internal/models/user"
)

// UserService ...
type UserService interface {
	Create(ctx context.Context, user *modelService.ToCreate) (int64, error)
	Get(ctx context.Context, id int64) (*modelService.User, error)
	Update(ctx context.Context, id int64, info *modelService.Info) error
	Delete(ctx context.Context, id int64) error
}
