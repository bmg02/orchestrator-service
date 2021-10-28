package main

import (
	"context"
	"errors"
	"net"

	"github.com/bmg02/orchestrator-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	serverInstance := grpc.NewServer()

	proto.RegisterUserServiceServer(serverInstance, &server{})
	reflection.Register(serverInstance)

	if e := serverInstance.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) GetUserByName(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	return nil, errors.New("Not implemented yet. Bhuvan will implement me.")
}
