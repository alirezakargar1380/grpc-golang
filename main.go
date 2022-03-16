package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alirezakargar1380/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("helo 1")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("%v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("helo")
}
