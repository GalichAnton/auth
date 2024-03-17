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

// AuthService ...
type AuthService interface {
	Login(ctx context.Context, login *modelService.Login) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, accessToken string) (string, error)
}

// AccessService ...
type AccessService interface {
	Check(ctx context.Context, address string) error
}
