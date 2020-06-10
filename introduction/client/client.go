package main

import (
	"context"
	"fmt"
	"log"

	intropb "github.com/golang-grpc-snippet/introduction/protobuf"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'am Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect to %v", err)
	}
	defer conn.Close()

	c := intropb.NewIntroServiceClient(conn)
	doUnary(c)

}

func doUnary(c intropb.IntroServiceClient) {
	req := &intropb.IntroRequest{
		Introduce: &intropb.Introduce{
			FirstName: "Muhammad",
			LastName:  "Auzan",
		},
	}
	res, err := c.Intro(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while introducing : %v", err)
	}

	log.Printf("Response : %v", res)
}
