package user

import (
	"context"

	modelService "github.com/GalichAnton/auth/internal/models/user"
	modelRepo "github.com/GalichAnton/auth/internal/repository/user/model"
)

func (s *service) Get(ctx context.Context, id int64) (*modelService.User, error) {
	user, err := s.userRepository.Get(ctx, modelRepo.Filter{ID: &id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
