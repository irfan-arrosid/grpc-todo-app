package main

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	pb "github.com/irfan-arrosid/grpc-todo-app/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_server_CreateTodo(t *testing.T) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type fields struct {
		db                             *gorm.DB
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateTodoResponse
		wantErr bool
	}{
		{
			name: "CreateTodo_Success",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.CreateTodoRequest{
					Title:     "Test Todo",
					Completed: false,
				},
			},
			want: &pb.CreateTodoResponse{
				Id:        1,
				Title:     "Test Todo",
				Completed: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				db:                             tt.fields.db,
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.CreateTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetTodos(t *testing.T) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type fields struct {
		db                             *gorm.DB
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetTodosRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetTodosResponse
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db:                             db,
				UnimplementedTodoServiceServer: pb.UnimplementedTodoServiceServer{},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.GetTodosRequest{},
			},
			want: &pb.GetTodosResponse{
				TodoList: []*pb.GetTodosResponse_Todo{
					{
						Id:        1,
						Title:     "Test Todo",
						Completed: false,
					},
					{
						Id:        2,
						Title:     "Test Todo",
						Completed: false,
					},
					{
						Id:        3,
						Title:     "Test Todo",
						Completed: false,
					},
					{
						Id:        4,
						Title:     "Test Todo",
						Completed: false,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				db:                             tt.fields.db,
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.GetTodos(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_UpdateTodo(t *testing.T) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type fields struct {
		db                             *gorm.DB
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateTodoResponse
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db:                             db,
				UnimplementedTodoServiceServer: pb.UnimplementedTodoServiceServer{},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.UpdateTodoRequest{
					Id:        1,
					Title:     "Updated Todo",
					Completed: true,
				},
			},
			want: &pb.UpdateTodoResponse{
				Id:        1,
				Title:     "Updated Todo",
				Completed: true,
			},
			wantErr: false,
		},
		{
			name: "Todo not found",
			fields: fields{
				db:                             db,
				UnimplementedTodoServiceServer: pb.UnimplementedTodoServiceServer{},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.UpdateTodoRequest{
					Id:        999,
					Title:     "Updated Todo",
					Completed: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				db:                             tt.fields.db,
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.UpdateTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_DeleteTodo(t *testing.T) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	type fields struct {
		db                             *gorm.DB
		UnimplementedTodoServiceServer pb.UnimplementedTodoServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteTodoResponse
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db:                             db,
				UnimplementedTodoServiceServer: pb.UnimplementedTodoServiceServer{},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.DeleteTodoRequest{
					Id: 1,
				},
			},
			want:    &pb.DeleteTodoResponse{},
			wantErr: false,
		},
		{
			name: "Todo not found",
			fields: fields{
				db:                             db,
				UnimplementedTodoServiceServer: pb.UnimplementedTodoServiceServer{},
			},
			args: args{
				ctx: context.TODO(),
				req: &pb.DeleteTodoRequest{
					Id: 999,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				db:                             tt.fields.db,
				UnimplementedTodoServiceServer: tt.fields.UnimplementedTodoServiceServer,
			}
			got, err := s.DeleteTodo(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
