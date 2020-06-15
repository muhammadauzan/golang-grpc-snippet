package main

import (
	"context"
	"fmt"
	"log"
	"net"

	substractpb "github.com/golang-grpc-snippet/drill_exercise_1/substraction/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Substract(ctx context.Context, req *substractpb.SubstractRequest) (*substractpb.SubstractResponse, error) {
	first := req.GetSubstraction().GetFirst()
	second := req.GetSubstraction().GetSecond()
	result := first - second
	res := &substractpb.SubstractResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Server start..")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	s := grpc.NewServer()

	substractpb.RegisterSubstractServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
}
