package main

import (
	"context"
	"fmt"
	"log"
	"net"

	multipb "github.com/golang-grpc-snippet/exercise/multiplication/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Multiplication(c context.Context, req *multipb.MultiRequest) (*multipb.MultiResponse, error) {
	first := req.GetNum().GetFirst()
	second := req.GetNum().GetSecond()
	fmt.Printf("Perkalian %v dengan %v", first, second)
	result := first * second
	res := &multipb.MultiResponse{
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

	multipb.RegisterMultiServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Error : %v", err)
	}
}
