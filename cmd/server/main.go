package main

import (
	"context"
	"log"

	"github.com/dcwk/gophkeeper/internal/application"
	pb "github.com/dcwk/gophkeeper/pkg/gophkeeper"
)

type GophkeeperServer struct {
	pb.UnimplementedGophkeeperServer
}

func main() {
	ctx := context.Background()

	a, err := application.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %s", err)
	}

	err = a.Start()
	if err != nil {
		log.Fatalf("failed to start application: %s", err)
	}

}
