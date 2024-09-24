package api

import (
	"github.com/dcwk/gophkeeper/internal/service"
	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

type Controller struct {
	gophkeeper.UnimplementedGophkeeperServer
	userService service.UserService
}

func NewController(service service.UserService) *Controller {
	return &Controller{userService: service}
}
