package api

import "github.com/dcwk/gophkeeper/proto"

type Controller struct {
	proto.UnimplementedGophkeeperServer
}

func NewController() *Controller {
	return &Controller{}
}
