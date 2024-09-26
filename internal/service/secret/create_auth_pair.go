package secret

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

func (s *Service) CreateAuthPair(ctx context.Context, authPair *entity.AuthPair) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.secretRepository.CreateSecret(ctx, authPair.Secret)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
