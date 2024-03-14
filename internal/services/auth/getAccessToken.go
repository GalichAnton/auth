package auth

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/claims"
	"github.com/GalichAnton/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (*string, error) {
	refreshTokenSecret := s.tokensConfig.Config().RefreshSecret
	accessTokenSecret := s.tokensConfig.Config().AccessSecret
	claim, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecret))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	user, err := s.userRepository.GetByEmail(ctx, claim.Email)
	if err != nil {
		return nil, err
	}

	newAccessToken, err := utils.GenerateToken(
		claims.UserClaims{
			Email: user.Info.Email,
			Role:  user.Info.Role,
		},
		[]byte(accessTokenSecret),
		s.tokensConfig.Config().AccessExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &newAccessToken, nil
}
