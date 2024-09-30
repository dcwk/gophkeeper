package api

import (
	"context"
	"time"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) GetUserSecretsList(ctx context.Context, req *gophkeeper.GetUserSecretsRequest) (*gophkeeper.GetUserSecretsResponse, error) {
	secrets, err := c.secretService.GetSecretsList(ctx)
	if err != nil {
		return nil, err
	}

	respSecrets := []*gophkeeper.Secret{}
	for _, secret := range secrets {
		respSecret := gophkeeper.Secret{
			Name:      secret.Name,
			Type:      secret.SecretType,
			CreatedAt: secret.CreatedAt.Format(time.RFC3339),
		}

		respSecrets = append(respSecrets, &respSecret)
	}

	return &gophkeeper.GetUserSecretsResponse{Secrets: respSecrets}, nil
}
