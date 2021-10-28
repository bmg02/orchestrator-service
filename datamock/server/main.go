package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/bmg02/orchestrator-service/datamock/proto10000"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Running on :10000")

	serverInstance := grpc.NewServer()

	proto10000.RegisterUserServiceServer(serverInstance, &server{})

	if e := serverInstance.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) GetUserMockData(ctx context.Context, request *proto10000.Request) (*proto10000.Response, error) {

	log.Printf("Log: Name input: %s\n", request.Name)

	nameLength := len(request.Name)

	if nameLength < 6 {
		return &proto10000.Response{
			Name:  "Name length must be greater than 6.",
			Class: "-1",
			Roll:  int64(-1),
		}, nil
	} else {
		return &proto10000.Response{
			Name:  request.Name,
			Class: strconv.Itoa(nameLength),
			Roll:  int64(nameLength * 10),
		}, nil
	}
}
