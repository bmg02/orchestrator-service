package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/bmg02/orchestrator-service/datamock/proto10000"
	"github.com/bmg02/orchestrator-service/getUser/proto9001"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Running on :9001")

	serverInstance := grpc.NewServer()

	proto9001.RegisterUserService9001Server(serverInstance, &server{})

	if e := serverInstance.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) GetUser(ctx context.Context, request *proto9001.RequestFor9001) (*proto9001.ResponseFor9001, error) {

	log.Printf("Log: Name input: %s\n", request.Name)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	c := proto10000.NewUserServiceClient(conn)

	response, err := c.GetUserMockData(context.Background(), &proto10000.Request{
		Name: strings.TrimSpace(request.Name),
	})

	if err != nil {
		log.Println(err)
	}

	log.Printf("Response: %s", response.String())

	return &proto9001.ResponseFor9001{
		Name:  response.Name,
		Class: response.Class,
		Roll:  response.Roll,
	}, nil
}
