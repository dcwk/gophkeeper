package repository

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
}
