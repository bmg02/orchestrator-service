package main

import (
	"context"
	"log"

	"github.com/bmg02/orchestrator-service/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewUserServiceClient(conn)

	response, err := c.GetUserByName(context.Background(), &proto.Request{})

	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", response.String())
}
