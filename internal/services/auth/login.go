package auth

import (
	"context"
	"errors"

	"github.com/GalichAnton/auth/internal/models/claims"
	modelService "github.com/GalichAnton/auth/internal/models/user"
	"github.com/GalichAnton/auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, login *modelService.Login) (*string, error) {
	user, err := s.userRepository.GetByEmail(ctx, login.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Info.Password), []byte(login.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	refreshToken, err := utils.GenerateToken(
		claims.UserClaims{
			Email: user.Info.Email,
			Role:  user.Info.Role,
		},
		[]byte(s.tokensConfig.Config().RefreshSecret),
		s.tokensConfig.Config().RefreshExpiration,
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &refreshToken, nil
}
