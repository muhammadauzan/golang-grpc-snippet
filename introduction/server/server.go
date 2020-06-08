package main

import (
	"log"
	"net"

	intropb "github.com/golang-grpc-snippet/introduction/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	intropb.RegisterIntroServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error to serve: %v", err)
	}
}
