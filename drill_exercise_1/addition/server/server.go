package main

import (
	"context"
	"fmt"
	"log"
	"net"

	addpb "github.com/golang-grpc-snippet/drill_exercise_1/addition/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *addpb.AddRequest) (*addpb.AddResponse, error) {
	fmt.Println("Service start..")
	first := req.GetAddition().GetFirst()
	second := req.GetAddition().GetSecond()
	result := first + second

	res := &addpb.AddResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Server starting..")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	errCheck(err)

	s := grpc.NewServer()
	addpb.RegisterAddServiceServer(s, &server{})

	err = s.Serve(lis)

	errCheck(err)
}

func errCheck(e error) {
	if e != nil {
		log.Printf("Error : %v", e)
	}
}
