package user

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/model"
)

func (s *Service) Register(ctx context.Context, user model.User) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.CreateUser(ctx, user)
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
