package auth

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/claims"
	"github.com/GalichAnton/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRefreshToken(ctx context.Context, refreshToken string) (*string, error) {
	refreshSecret := s.tokensConfig.Config().RefreshSecret
	claim, err := utils.VerifyToken(refreshToken, []byte(refreshSecret))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	user, err := s.userRepository.GetByEmail(ctx, claim.Email)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateToken(
		claims.UserClaims{
			Email: user.Info.Email,
			Role:  user.Info.Role,
		},
		[]byte(refreshSecret),
		s.tokensConfig.Config().RefreshExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &newRefreshToken, nil
}
