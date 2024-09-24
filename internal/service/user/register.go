package user

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/model"
)

func (s *Service) Register(ctx context.Context, user model.User) (int64, error) {
	return 11, nil
}
