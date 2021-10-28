package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bmg02/orchestrator-service/getUserByName/proto9000"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	c := proto9000.NewUserService9000Client(conn)

	userInput := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name of user: ")

	input, err := userInput.ReadString('\n')

	if err != nil {
		log.Println(err)
	}

	response, err := c.GetUserByName(context.Background(), &proto9000.RequestFor9000{
		Name: strings.TrimSpace(input),
	})

	if err != nil {
		log.Println(err)
	}

	if response.Roll != -1 {
		log.Printf("Response: %s", response.String())
	} else {
		log.Printf("Response: %s", response.Name)
	}
}
