package main

import (
	"GRPC_POC/chat"
	chat "GRPC_POC/chat"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Chat Server Initiating")
	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(2)
	}
	fmt.Println("Server Started.")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
