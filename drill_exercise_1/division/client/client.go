package main

import (
	"context"
	"fmt"
	"log"

	divpb "github.com/golang-grpc-snippet/drill_exercise_1/division/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	c := divpb.NewDivClient(conn)
	doUnary(c)
}

func doUnary(c divpb.DivClient) {
	req := &divpb.DivRequest{
		Number: &divpb.Number{
			First:  30,
			Second: 15,
		},
	}
	res, err := c.Division(context.Background(), req)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	fmt.Printf("Hasil : %d", res.GetResult())
}
