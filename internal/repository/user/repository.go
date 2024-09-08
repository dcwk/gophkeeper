package user

import (
	"context"

	"github.com/dcwk/gophkeeper/proto"
)

type UserRepository struct {
}

func (r *UserRepository) CreateUser(ctx context.Context, request proto.RegisterRequest) (int64, error) {
	return 1, nil
}
