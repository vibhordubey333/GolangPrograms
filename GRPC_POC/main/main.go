package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to server : %s", err.Error)
	}
}
