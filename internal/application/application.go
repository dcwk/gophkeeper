package application

import (
	"context"
	"fmt"
	"log"
	"net"

	"practicum/gophkeeper/internal/config"
	"practicum/gophkeeper/proto"

	"google.golang.org/grpc"
)

type Application struct {
}

type GophkeeperServer struct {
	Conf *config.ServerConf
	proto.UnimplementedGophkeeperServer
}

func (s *GophkeeperServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println(fmt.Sprintf("Success get data with login %s and password %s", req.Login, req.Password))

	return &proto.RegisterResponse{UserId: "10"}, nil
}

func Run(conf *config.ServerConf) {
	listen, err := net.Listen("tcp", conf.RunAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	server := GophkeeperServer{Conf: conf}
	proto.RegisterGophkeeperServer(s, &server)

	fmt.Println(fmt.Sprintf("gRPC server started at %s...", conf.RunAddress))
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
