package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bmg02/orchestrator-service/getUser/proto9001"
	"github.com/bmg02/orchestrator-service/getUserByName/proto9000"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Running on :9000")

	serverInstance := grpc.NewServer()

	proto9000.RegisterUserService9000Server(serverInstance, &server{})

	if e := serverInstance.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) GetUserByName(ctx context.Context, request *proto9000.RequestFor9000) (*proto9000.ResponseFor9000, error) {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	c := proto9001.NewUserService9001Client(conn)

	fmt.Println("Sending request to 9001")

	response, err := c.GetUser(context.Background(), &proto9001.RequestFor9001{
		Name: request.Name,
	})

	if err != nil {
		log.Println(err)
	}

	log.Printf("Response: %s", response.String())

	return &proto9000.ResponseFor9000{
		Name:  response.Name,
		Class: response.Class,
		Roll:  response.Roll,
	}, nil
}
