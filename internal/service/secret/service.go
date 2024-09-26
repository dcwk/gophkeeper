package secret

import (
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
)

type Service struct {
	secretRepository   repository.SecretRepository
	authPairRepository repository.AuthPairRepository
	metadataRepository repository.MetadataRepository
	txManager          db.TxManager
}

func NewService(
	secretRepository repository.SecretRepository,
	authPairRepository repository.AuthPairRepository,
	metadataRepository repository.MetadataRepository,
	txManager db.TxManager,
) *Service {
	return &Service{
		secretRepository:   secretRepository,
		authPairRepository: authPairRepository,
		metadataRepository: metadataRepository,
		txManager:          txManager,
	}
}
