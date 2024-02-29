package user

import (
	"github.com/GalichAnton/auth/internal/repository"
	"github.com/GalichAnton/auth/internal/services"
	"github.com/GalichAnton/platform_common/pkg/db"
)

var _ services.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(
	userRepository repository.UserRepository, txManager db.TxManager, logRepository repository.LogRepository,
) *service {
	return &service{
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}
