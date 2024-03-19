package user

import (
	"context"
	"errors"

	"github.com/GalichAnton/auth/internal/models/log"
	modelService "github.com/GalichAnton/auth/internal/models/user"
	"github.com/jackc/pgconn"
)

func (s *service) Create(ctx context.Context, info *modelService.ToCreate) (int64, error) {
	var newUserID int64

	if info.Password != info.PasswordConfirm {
		return 0, errors.New("password and password confirmation do not match")
	}

	userInfo := modelService.Info{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role,
	}

	err := s.txManager.ReadCommitted(
		ctx, func(ctx context.Context) error {
			id, errTx := s.userRepository.Create(ctx, &userInfo)
			if errTx != nil {
				return errTx
			}

			newUserID = id
			newLog := log.Info{
				Action:   "create",
				EntityID: id,
			}

			errTx = s.logRepository.Create(ctx, &newLog)
			if errTx != nil {
				return errTx
			}

			return nil
		},
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return 0, errors.New("a user with this email already exists")
			}
		}
	}

	return newUserID, nil
}
