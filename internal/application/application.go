package application

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/dcwk/gophkeeper/internal/closer"
	"github.com/dcwk/gophkeeper/internal/config"
	"github.com/dcwk/gophkeeper/proto"

	"google.golang.org/grpc"
)

type Application struct {
	config     *config.ServerConf
	container  *Container
	grpcServer *grpc.Server
}

func NewApp(ctx context.Context) (*Application, error) {
	app := &Application{}

	err := app.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *Application) initDependencies(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		app.initConfig,
		app.initContainer,
		app.initGrpcServer,
	}

	for _, initFunc := range inits {
		err := initFunc(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *Application) initConfig(ctx context.Context) error {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	app.config = conf
	return nil
}

func (app *Application) initContainer(ctx context.Context) error {
	app.container = newContainer()

	return nil
}

func (app *Application) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(app.grpcServer)

	proto.RegisterGophkeeperServer(
		app.grpcServer,
		app.container.Controller(),
	)

	return nil
}

func (app *Application) Start() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return app.runGRPCServer()
}

func (app *Application) runGRPCServer() error {
	log.Println("starting gRPC server on %s", app.config.RunAddress)

	list, err := net.Listen("tcp", app.config.RunAddress)
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
