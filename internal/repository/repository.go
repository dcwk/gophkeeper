package repository

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	GeUserByLogin(ctx context.Context, login string) (*entity.User, error)
}

type SecretRepository interface {
	CreateSecret(ctx context.Context, secret *entity.Secret) (int64, error)
	GetSecretsList(ctx context.Context) ([]*entity.Secret, error)
}

type MetadataRepository interface {
	CreateMetadata(ctx context.Context, metadata *entity.Metadata) (int64, error)
}

type AuthPairRepository interface {
	CreateAuthPair(ctx context.Context, pair *entity.AuthPair) (int64, error)
}
