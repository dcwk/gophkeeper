package user

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

func (s *Service) Register(ctx context.Context, user entity.User) (int64, error) {
	err := user.HashPassword()
	if err != nil {
		return 0, err
	}

	id, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
