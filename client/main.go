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
	reqCreateTodo := &pb.CreateTodoRequest{
		Title:     "Sleeping",
		Completed: false,
	}

	resCreateTodo, err := client.CreateTodo(context.Background(), reqCreateTodo)
	if err != nil {
		log.Fatal("Error to create todo", err.Error())
	}

	fmt.Printf("Created Todo: ID=%d, Title=%s, Completed=%t\n", resCreateTodo.Id, resCreateTodo.Title, resCreateTodo.Completed)

	// Call GetTodos service
	reqGetTodos := &pb.GetTodosRequest{}

	resGetTodos, err := client.GetTodos(context.Background(), reqGetTodos)
	if err != nil {
		log.Fatal("Error to get todos", err.Error())
	}

	for _, todo := range resGetTodos.TodoList {
		fmt.Printf("Todo: ID=%d, Title=%s, Completed=%t\n", todo.Id, todo.Title, todo.Completed)
	}
}
