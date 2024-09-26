package api

import (
	"github.com/dcwk/gophkeeper/internal/service"
	"github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

type Controller struct {
	gophkeeper.UnimplementedGophkeeperServer
	userService   service.UserService
	secretService service.SecretService
}

func NewController(
	userService service.UserService,
	secretService service.SecretService,
) *Controller {
	return &Controller{
		userService:   userService,
		secretService: secretService,
	}
}
