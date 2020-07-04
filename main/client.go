package main

import (
	"context"
	"github.com/edwardsuwirya/simpleCalcRpcClient/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewAdditionClient(conn)

	response, err := c.AddNumber(context.Background(), &api.AdditionMessage{
		Num1: 2,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("%v", response.ResNum)
}
