package converter

import (
	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/repository/secret/model"
)

func ToSecretFromRepo(secret *model.Secret) *entity.Secret {
	return &entity.Secret{
		ID:         secret.ID,
		User:       entity.User{ID: secret.UserID},
		SecretType: secret.Type,
		Name:       secret.Name,
	}
}
