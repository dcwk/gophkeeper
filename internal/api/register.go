package api

import (
	"context"
	"fmt"

	"github.com/dcwk/gophkeeper/internal/model"
	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

func (c *Controller) Register(ctx context.Context, req *gophkeeper.RegisterRequest) (*gophkeeper.RegisterResponse, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}

	userID, err := c.userService.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Success get data with login %s and password %s", req.Login, req.Password))

	return &gophkeeper.RegisterResponse{UserId: userID}, nil
}
