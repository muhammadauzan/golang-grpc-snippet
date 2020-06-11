package main

import (
	"context"
	"log"
	"net"

	filterpb "github.com/golang-grpc-snippet/drill_exercise_1/filter/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (server) FilterMethod(c context.Context, req *filterpb.FilterRequest) (*filterpb.FilterResponse, error) {
	num := req.GetData().GetNum()
	result := false
	if num > 10 {
		result = true
	}

	res := filterpb.FilterResponse{
		Result: result,
	}

	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	s := grpc.NewServer()
	filterpb.RegisterFilterServiceServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
}
