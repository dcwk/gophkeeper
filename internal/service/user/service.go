package user

import (
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
)

type Service struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) *Service {
	return &Service{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
