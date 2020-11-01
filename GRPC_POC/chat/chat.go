package chat

import (
	proto "GRPC_POC/proto_files"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &proto.Message{Body: "Hello From the Server!"}, nil
}
