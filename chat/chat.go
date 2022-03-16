package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("recieved message from body: %v", message.Body)
	return &Message{Body: "Hello From  the Server"}, nil
}
