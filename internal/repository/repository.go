package repository

import (
	"context"

	"github.com/dcwk/gophkeeper/proto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, request proto.RegisterRequest) (int64, error)
}
