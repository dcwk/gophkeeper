package secret

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
)

const (
	tableName = "secret"

	idColumn        = "id"
	userIDColumn    = "login"
	typeColumn      = "type"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SecretRepository {
	return &repo{db: db}
}

func (r *repo) CreateSecret(ctx context.Context, secret *entity.Secret) (int64, error) {
	return 10, nil
}
