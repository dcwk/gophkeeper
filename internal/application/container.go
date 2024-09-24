package application

import (
	"context"
	"log"

	"github.com/dcwk/gophkeeper/internal/api"
	"github.com/dcwk/gophkeeper/internal/closer"
	"github.com/dcwk/gophkeeper/internal/config"
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/infra/db/pg"
	"github.com/dcwk/gophkeeper/internal/infra/db/transaction"
	"github.com/dcwk/gophkeeper/internal/repository"
	"github.com/dcwk/gophkeeper/internal/repository/user"
)

type Container struct {
	config     *config.ServerConf
	controller *api.Controller
	dbClient   db.Client
	txManager  db.TxManager
	userRepo   repository.UserRepository
}

func newContainer() *Container {
	return &Container{}
}

func (container *Container) DBClient(ctx context.Context) db.Client {
	if container.dbClient == nil {
		cl, err := pg.New(ctx, container.config.DatabaseDSN)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %container", err.Error())
		}
		closer.Add(cl.Close)

		container.dbClient = cl
	}

	return container.dbClient
}

func (container *Container) TxManager(ctx context.Context) db.TxManager {
	if container.txManager == nil {
		container.txManager = transaction.NewTransactionManager(container.DBClient(ctx).DB())
	}

	return container.txManager
}

func (container *Container) userRepository() repository.UserRepository {
	if container.userRepo == nil {
		container.userRepo = user.NewRepository(container.dbClient)
	}

	return container.userRepo
}

func (container *Container) Controller() *api.Controller {
	if container.controller == nil {
		container.controller = api.NewController(container.userRepository())
	}

	return container.controller
}
