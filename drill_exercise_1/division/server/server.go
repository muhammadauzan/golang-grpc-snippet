package main

import (
	"context"
	"log"
	"net"

	divpb "github.com/golang-grpc-snippet/drill_exercise_1/division/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Division(c context.Context, req *divpb.DivRequest) (*divpb.DivResponse, error) {
	first := req.GetNumber().GetFirst()
	second := req.GetNumber().GetSecond()
	result := first / second
	res := &divpb.DivResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	s := grpc.NewServer()

	divpb.RegisterDivServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
}
