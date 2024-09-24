package service

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/model"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
}
