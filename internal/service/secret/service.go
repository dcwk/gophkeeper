package secret

import (
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
)

type Service struct {
	secretRepository repository.SecretRepository
	txManager        db.TxManager
}

func NewService(
	secretRepository repository.SecretRepository,
	txManager db.TxManager,
) *Service {
	return &Service{
		secretRepository: secretRepository,
		txManager:        txManager,
	}
}
