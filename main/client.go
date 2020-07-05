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

	c := api.NewCalcClient(conn)

	response, err := c.AddNumber(context.Background(), &api.ParameterMessage{
		Num1: 2,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Error when calling AddNumber: %s", err)
	}
	log.Printf("%v", response.ResNum)

	response, err = c.SubNumber(context.Background(), &api.ParameterMessage{
		Num1: 2,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Error when calling SubNumber: %s", err)
	}
	log.Printf("%v", response.ResNum)

	response, err = c.MulNumber(context.Background(), &api.ParameterMessage{
		Num1: 2,
		Num2: 3,
	})
	if err != nil {
		log.Fatalf("Error when calling MulNumber: %s", err)
	}
	log.Printf("%v", response.ResNum)

	response, err = c.DivNumber(context.Background(), &api.ParameterMessage{
		Num1: 10,
		Num2: 5,
	})
	if err != nil {
		log.Fatalf("Error when calling DivNumber: %s", err)
	}
	log.Printf("%v", response.ResNum)

	response, err = c.DivNumber(context.Background(), &api.ParameterMessage{
		Num1: 2,
		Num2: 0,
	})
	if err != nil {
		log.Fatalf("Error when calling DivNumber: %s", err)
	}
	log.Printf("%v", response.ResNum)
}
