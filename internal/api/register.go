package api

import (
	"context"
	"fmt"

	"github.com/dcwk/gophkeeper/proto"
)

func (c *Controller) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println(fmt.Sprintf("Success get data with login %s and password %s", req.Login, req.Password))

	return &proto.RegisterResponse{UserId: "10"}, nil
}
