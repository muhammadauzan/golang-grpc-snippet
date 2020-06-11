package main

import (
	"context"
	"fmt"
	"log"

	substractpb "github.com/golang-grpc-snippet/exercise/substraction/protobuf"
	"google.golang.org/grpc"
)

func main() {
	fmt.Print("Client start..")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close()

	c := substractpb.NewSubstractServiceClient(conn)
	doUnary(c)
}

func doUnary(c substractpb.SubstractServiceClient) {
	req := &substractpb.SubstractRequest{
		Substraction: &substractpb.Substraction{
			First:  30,
			Second: 15,
		},
	}

	res, err := c.Substract(context.Background(), req)

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	log.Printf("Result : %v", res.GetResult())
}
