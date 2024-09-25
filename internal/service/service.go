package service

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

type UserService interface {
	Register(ctx context.Context, user entity.User) (int64, error)
}

type SecretService interface {
	CreateAuthPair(ctx context.Context, authPair entity.AuthPair) (int64, error)
}
