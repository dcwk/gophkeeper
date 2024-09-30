package api

import (
	"context"
	"fmt"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) SaveAuthPair(ctx context.Context, req *gophkeeper.SaveAuthPairRequest) (*gophkeeper.SaveAuthPairResponse, error) {
	authPair := entity.NewAuthPair(req.Login, req.Password, req.SecretName, req.Metadata)

	secretID, err := c.secretService.CreateAuthPair(ctx, authPair)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Success get auth pair data: id %v", secretID))

	return &gophkeeper.SaveAuthPairResponse{SecretId: secretID}, nil
}
