package main

import (
	"context"
	"fmt"
	"log"

	multipb "github.com/golang-grpc-snippet/exercise/multiplication/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	c := multipb.NewMultiServiceClient(conn)
	doUnary(c)
}

func doUnary(c multipb.MultiServiceClient) {
	nilai := &multipb.MultiRequest{
		Num: &multipb.Number{
			First:  30,
			Second: 3,
		},
	}

	res, err := c.Multiplication(context.Background(), nilai)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	fmt.Println("Hasil : ", res.GetResult())
}
