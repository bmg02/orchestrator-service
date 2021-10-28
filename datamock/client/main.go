package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bmg02/orchestrator-service/datamock/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewUserServiceClient(conn)

	userInput := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name of user: ")

	input, err := userInput.ReadString('\n')

	if err != nil {
		panic(err)
	}

	response, err := c.GetUserMockData(context.Background(), &proto.Request{Name: strings.TrimSpace(input)})

	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", response.String())
}
