package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/bmg02/orchestrator-service/datamock/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}

	serverInstance := grpc.NewServer()

	proto.RegisterUserServiceServer(serverInstance, &server{})

	if e := serverInstance.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) GetUserMockData(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	log.Printf("Log: Name input: %s\n", request.Name)

	nameLength := len(request.Name)

	if nameLength < 6 {
		return nil, errors.New("Name length must be greater than 6.")
	} else {
		return &proto.Response{
			Name:  request.Name,
			Class: strconv.Itoa(nameLength),
			Roll:  int64(nameLength * 10),
		}, nil
	}
}
