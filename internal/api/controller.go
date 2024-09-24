package api

import (
	"github.com/dcwk/gophkeeper/internal/repository"
	"github.com/dcwk/gophkeeper/internal/service/user"
	"github.com/dcwk/gophkeeper/proto"
)

type Controller struct {
	proto.UnimplementedGophkeeperServer
	userService *user.Service
}

func NewController(userRepository repository.UserRepository) *Controller {
	return &Controller{
		userService: user.NewService(userRepository),
	}
}
