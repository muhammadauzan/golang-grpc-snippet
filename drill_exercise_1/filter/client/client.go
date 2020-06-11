package main

import (
	"context"
	"fmt"
	"log"

	filterpb "github.com/golang-grpc-snippet/drill_exercise_1/filter/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	c := filterpb.NewFilterServiceClient(conn)
	doUnary(c)

}

func doUnary(c filterpb.FilterServiceClient) {
	req := filterpb.FilterRequest{
		Data: &filterpb.Number{
			Num: 15,
		},
	}
	res, err := c.FilterMethod(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	fmt.Println("Hasil : ", res.GetResult())
}
