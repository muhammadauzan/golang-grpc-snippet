package main

import (
	"context"
	"fmt"
	"log"
	"net"

	intropb "github.com/golang-grpc-snippet/introduction/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Intro(ctx context.Context, req *intropb.IntroRequest) (*intropb.IntroResponse, error) {
	fmt.Println("Incoming introduction")
	firstname := req.GetIntroduce().GetFirstName()
	result := "Hello " + firstname
	res := &intropb.IntroResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Printf("Server starting..")

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
