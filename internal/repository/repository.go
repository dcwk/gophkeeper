package repository

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
}
