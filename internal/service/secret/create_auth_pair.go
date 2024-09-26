package secret

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

func (s *Service) CreateAuthPair(ctx context.Context, authPair *entity.AuthPair) (int64, error) {
	return 10, nil
}
