package auth

import (
	"context"

	desc "github.com/GalichAnton/auth/pkg/auth_v1"
)

// GetRefreshToken ...
func (i *Implementation) GetRefreshToken(
	ctx context.Context, req *desc.GetRefreshTokenRequest,
) (*desc.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
