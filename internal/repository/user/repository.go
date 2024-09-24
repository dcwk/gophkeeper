package user

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
	"github.com/dcwk/gophkeeper/proto"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, request proto.RegisterRequest) (int64, error) {
	return 1, nil
}
