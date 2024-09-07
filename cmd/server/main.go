package main

import (
	"log"

	"practicum/gophkeeper/internal/application"
	"practicum/gophkeeper/internal/config"
	pb "practicum/gophkeeper/proto"
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
