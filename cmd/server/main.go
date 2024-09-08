package main

import (
	"log"

	"github.com/dcwk/gophkeeper/internal/application"
	"github.com/dcwk/gophkeeper/internal/config"
	pb "github.com/dcwk/gophkeeper/proto"
)

type GophkeeperServer struct {
	pb.UnimplementedGophkeeperServer
}

func main() {
	conf, err := config.NewServerConf()
	if err != nil {
		log.Fatal(err)
	}

	application.Run(conf)
}
