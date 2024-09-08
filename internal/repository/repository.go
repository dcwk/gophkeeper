package repository

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, request proto.AddUserRequest) (int64, error)
}
