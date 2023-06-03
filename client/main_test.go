package main

import (
	"log"
	"testing"

	pb "github.com/irfan-arrosid/grpc-todo-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateTodo(t *testing.T) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	type args struct {
		client pb.TodoServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test CreateTodo",
			args: args{
				client: client,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateTodo(tt.args.client)
		})
	}
}

func TestGetTodos(t *testing.T) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	type args struct {
		client pb.TodoServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test GetTodos",
			args: args{
				client: client,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTodos(tt.args.client)
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	type args struct {
		client pb.TodoServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test UpdateTodo",
			args: args{
				client: client,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateTodo(tt.args.client)
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error in dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)

	type args struct {
		client pb.TodoServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test DeleteTodo",
			args: args{
				client: client,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteTodo(tt.args.client)
		})
	}
}
