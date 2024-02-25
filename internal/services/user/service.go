package user

import (
	"github.com/GalichAnton/auth/internal/client/db"
	"github.com/GalichAnton/auth/internal/repository"
	"github.com/GalichAnton/auth/internal/services"
)

var _ services.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(userRepository repository.UserRepository, txManager db.TxManager, logRepository repository.LogRepository) *service {
	return &service{
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}
