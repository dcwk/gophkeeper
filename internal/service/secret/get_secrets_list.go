package secret

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

func (s *Service) GetSecretsList(ctx context.Context) ([]*entity.Secret, error) {
	secrets, err := s.secretRepository.GetSecretsList(ctx)
	if err != nil {
		return nil, err
	}

	return secrets, nil
}
