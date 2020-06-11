package main

import (
	"context"
	"fmt"
	"log"

	modupb "github.com/golang-grpc-snippet/drill_exercise_1/modulation/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	c := modupb.NewModServiceClient(conn)
	doUnary(c)
}

func doUnary(c modupb.ModServiceClient) {
	req := modupb.ModRequest{
		Constant: &modupb.Variable{
			First:  30,
			Second: 2,
		},
	}
	res, err := c.Modulation(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	fmt.Println("Hasil :", res.GetResult())
}
