package auth

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/claims"
	"github.com/GalichAnton/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRefreshToken(_ context.Context, refreshToken string) (string, error) {
	claim, err := utils.VerifyToken(refreshToken, []byte(s.tokensConfig.Config().RefreshSecret))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	newRefreshToken, err := utils.GenerateToken(
		claims.UserClaims{
			Email: claim.Email,
			Role:  claim.Role,
		},
		[]byte(s.tokensConfig.Config().RefreshSecret),
		s.tokensConfig.Config().RefreshExpiration,
	)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}
