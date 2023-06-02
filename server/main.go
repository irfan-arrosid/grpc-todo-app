package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "github.com/irfan-arrosid/grpc-todo-app/proto"
)

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

type server struct {
	db *gorm.DB
	pb.UnimplementedTodoServiceServer
}

func (s *server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	todo := &Todo{
		Title:     req.Title,
		Completed: req.Completed,
	}

	s.db.Create(todo)

	return &pb.CreateTodoResponse{
		Id:        int64(todo.ID),
		Title:     todo.Title,
		Completed: todo.Completed,
	}, nil
}

func (s *server) GetTodos(ctx context.Context, req *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	var todos []*Todo

	s.db.Find(&todos)

	var resTodos []*pb.GetTodosResponse_Todo
	for _, todo := range todos {
		resTodo := &pb.GetTodosResponse_Todo{
			Id:        int64(todo.ID),
			Title:     todo.Title,
			Completed: todo.Completed,
		}
		resTodos = append(resTodos, resTodo)
	}

	return &pb.GetTodosResponse{
		TodoList: resTodos,
	}, nil
}

func (s *server) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	var todo Todo

	s.db.First(&todo, req.Id)
	if todo.ID == 0 {
		return nil, fmt.Errorf("Todo not found")
	}

	todo.Title = req.Title
	todo.Completed = req.Completed

	s.db.Save(&todo)

	return &pb.UpdateTodoResponse{
		Id:        int64(todo.ID),
		Title:     todo.Title,
		Completed: todo.Completed,
	}, nil
}

func (s *server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	var todo Todo

	s.db.First(&todo, req.Id)
	if todo.ID == 0 {
		return nil, fmt.Errorf("Todo not found")
	}

	s.db.Delete(&todo)

	return &pb.DeleteTodoResponse{}, nil
}

func main() {
	// Connect to the PostgreSQL database
	dsn := "host=localhost user=irfanarrosid password=at19ir97ar dbname=grpc-todo-app port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("Database connecting is success")

	// Automigrate the Todo struct to the PostgreSQL database
	db.AutoMigrate(&Todo{})

	// Create a new server
	grpcServer := grpc.NewServer()

	// Register TodoService server
	pb.RegisterTodoServiceServer(grpcServer, &server{db: db})

	// Start listening on port
	listen, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatalln("error to listen", err.Error())
	}
	log.Println("Server listening on port 2020")

	// Serve gRPC requests
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve grpc", err.Error())
	}
}
