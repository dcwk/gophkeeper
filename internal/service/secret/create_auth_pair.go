package secret

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

func (s *Service) CreateAuthPair(ctx context.Context, authPair *entity.AuthPair) (int64, error) {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		authPair.Secret.ID, errTx = s.secretRepository.CreateSecret(ctx, authPair.Secret)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.authPairRepository.CreateAuthPair(ctx, authPair)

		return nil
	})

	if err != nil {
		return 0, err
	}

	return authPair.Secret.ID, nil
}
