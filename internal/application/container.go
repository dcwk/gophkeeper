package application

import (
	"github.com/dcwk/gophkeeper/internal/api"
)

type Container struct {
	controller *api.Controller
}

func newContainer() *Container {
	return &Container{}
}

func (container *Container) Controller() *api.Controller {
	if container.controller == nil {
		container.controller = api.NewController()
	}

	return container.controller
}
