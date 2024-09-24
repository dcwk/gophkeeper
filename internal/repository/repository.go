package repository

import (
	"context"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

type UserRepository interface {
	CreateUser(ctx context.Context, request gophkeeper.RegisterRequest) (int64, error)
}
