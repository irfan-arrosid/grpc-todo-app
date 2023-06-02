package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/irfan-arrosid/grpc-todo-app/proto"
)

func main() {
	// Set up a connection to the server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":2020", opts...)
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	// Create a new TodoService client
	client := pb.NewTodoServiceClient(conn)

	// Call CreateTodo service
	req := &pb.CreateTodoRequest{
		Title:     "Sample Todo",
		Completed: false,
	}

	res, err := client.CreateTodo(context.Background(), req)
	if err != nil {
		log.Fatal("Error get todo by id", err.Error())
	}

	fmt.Printf("Created Todo: ID=%d, Title=%s, Completed=%t\n", res.Id, res.Title, res.Completed)
}
