package user

import (
	"context"

	modelService "github.com/GalichAnton/auth/internal/models/user"
	modelRepo "github.com/GalichAnton/auth/internal/repository/user/model"
	"github.com/pkg/errors"
)

func (s *service) Get(ctx context.Context, id int64) (*modelService.User, error) {
	if id == 0 {
		return nil, errors.Errorf("id is empty")
	}

	user, err := s.userRepository.Get(ctx, modelRepo.Filter{ID: &id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
