package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"context"

	pb "github.com/carlosalg/todo-grpc-go/api"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type todoServer struct {
	pb.UnimplementedTodoServiceServer
	todos []*pb.Todo
}

func (s *todoServer) CreateTodo(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	s.todos = append(s.todos, todo)
	return todo, nil
}

func (s *todoServer) GetTodoList(empty *empty.Empty, stream pb.TodoService_GetTodoListServer) error {
	for _, todo := range s.todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
