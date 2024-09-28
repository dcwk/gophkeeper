package user

import (
	"context"
	"fmt"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/infra/auth"
)

func (s *Service) Auth(ctx context.Context, user *entity.User) (string, error) {
	currentUser, err := s.userRepository.GeUserByLogin(ctx, user.Login)
	if err != nil || currentUser.ID == 0 {
		return "", fmt.Errorf("failed to find user by login: %w", err)
	}

	if !currentUser.VerifyPassword(user.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.BuildJWTString(currentUser.ID)
	if err != nil {
		return "", fmt.Errorf("failed to build JWT token: %w", err)
	}

	return token, nil
}
