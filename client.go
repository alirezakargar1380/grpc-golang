package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/alirezakargar1380/grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "hello from client!",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		panic(err)
	}

	log.Printf("response from server: %s", response.Body)

}
