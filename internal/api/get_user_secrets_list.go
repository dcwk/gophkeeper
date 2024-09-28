package api

import (
	"context"

	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) GetUserSecretsList(ctx context.Context, req *gophkeeper.GetUserSecretsRequest) (*gophkeeper.GetUserSecretsResponse, error) {
	secrets, err := c.secretService.GetSecretsList(ctx)
	if err != nil {
		return nil, err
	}

	return &gophkeeper.GetUserSecretsResponse{}
}
