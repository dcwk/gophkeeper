package main

import (
	"github.com/golang/protobuf/ptypes/wrappers"

	pb "practicum/gophkeeper/proto"
)

type GophkeeperServer struct {
	pb.UnimplementedGophkeeperServer
}

func main() {
	//conf, err := config.NewServerConf()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//application.Run(conf)
}

func (s *GophkeeperServer) SearchOrders(searchQuery *wrappers.StringValue, stream pb.Gophkeeper_SearchOrdersServer) error {

}
