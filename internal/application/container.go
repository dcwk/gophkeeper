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
	userRepository "github.com/dcwk/gophkeeper/internal/repository/user"
	"github.com/dcwk/gophkeeper/internal/service"
	userService "github.com/dcwk/gophkeeper/internal/service/user"
)

type Container struct {
	config          *config.ServerConf
	controller      *api.Controller
	dbClient_       db.Client
	txManager_      db.TxManager
	userRepository_ repository.UserRepository
	userService_    service.UserService
}

func newContainer(conf *config.ServerConf) *Container {
	return &Container{config: conf}
}

func (container *Container) dbClient(ctx context.Context) db.Client {
	if container.dbClient_ == nil {
		client, err := pg.New(ctx, container.config.DatabaseDSN)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %container", err.Error())
		}
		closer.Add(client.Close)

		container.dbClient_ = client
	}

	return container.dbClient_
}

func (container *Container) txManager(ctx context.Context) db.TxManager {
	if container.txManager_ == nil {
		container.txManager_ = transaction.NewTransactionManager(container.dbClient(ctx).DB())
	}

	return container.txManager_
}

func (container *Container) userRepository(ctx context.Context) repository.UserRepository {
	if container.userRepository_ == nil {
		container.userRepository_ = userRepository.NewRepository(container.dbClient(ctx))
	}

	return container.userRepository_
}

func (container *Container) userService(ctx context.Context) service.UserService {
	if container.userService_ == nil {
		container.userService_ = userService.NewService(
			container.userRepository(ctx),
			container.txManager(ctx),
		)
	}

	return container.userService_
}

func (container *Container) Controller(ctx context.Context) *api.Controller {
	if container.controller == nil {
		container.controller = api.NewController(container.userService(ctx))
	}

	return container.controller
}
