package user

import (
	"context"

	"github.com/GalichAnton/auth/internal/models/log"
	modelService "github.com/GalichAnton/auth/internal/models/user"
)

func (s *service) Update(ctx context.Context, id int64, info *modelService.Info) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		newLog := log.Info{
			Action:   "update",
			EntityID: id,
		}

		errTx = s.logRepository.Create(ctx, &newLog)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
