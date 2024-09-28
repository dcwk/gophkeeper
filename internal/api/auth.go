package api

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) Auth(ctx context.Context, req gophkeeper.AuthRequest) (*gophkeeper.AuthResponse, error) {
	user := entity.NewUser(req.Login, req.Password)
	jwtToken, err := c.userService.Auth(ctx, user)
	if err != nil {
		return nil, err
	}

	return &gophkeeper.AuthResponse{JwtToken: jwtToken}, nil
}
