package main

import (
	"context"
	"fmt"
	"log"

	addpb "github.com/golang-grpc-snippet/addition/protobuf"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client start..")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close()

	c := addpb.NewAddServiceClient(conn)
	doUnary(c)

}

func doUnary(c addpb.AddServiceClient) {
	req := addpb.AddRequest{
		Addition: &addpb.Addition{
			First:  27,
			Second: 34,
		},
	}

	res, err := c.Add(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	fmt.Println("The result is ", res)
}
