package main

import (
	"context"
	"log"
	"net"

	modupb "github.com/golang-grpc-snippet/drill_exercise_1/modulation/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Modulation(c context.Context, req *modupb.ModRequest) (*modupb.ModResponse, error) {
	first := req.GetConstant().GetFirst()
	second := req.GetConstant().GetSecond()
	result := first % second

	res := modupb.ModResponse{
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
	modupb.RegisterModServiceServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
}
