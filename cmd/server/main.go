package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"context"

	pb "github.com/carlosalg/todo-grpc-go/api"
	todorepo "github.com/carlosalg/todo-grpc-go/internal/database/todo_repo"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type todoServer struct {
	pb.UnimplementedTodoServiceServer
	todoRepo *todorepo.TodoRepository
}

func (s *todoServer) CreateTodo(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	todos := convertToRepoTodo(todo)
	createdTodo, err := s.todoRepo.CreateTodo(todos)
	if err != nil {
		return nil, err
	}
	finalTodo := convertToGRPC(createdTodo)
	return finalTodo, nil
}

func (s *todoServer) GetTodoList(empty *empty.Empty, stream pb.TodoService_GetTodoListServer) error {
	todos, err := s.todoRepo.GetAllTodos()
	if err != nil {
		return err
	}

	for _, todo := range todos {
		grpcTodo := convertToGRPC(todo)
		if err := stream.Send(grpcTodo); err != nil {
			return err
		}
	}
	return nil
}

func (s *todoServer) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	todoID := req.Id

	err := s.todoRepo.DeleteTodo(todoID)
	if err != nil {
		return &pb.DeleteTodoResponse{Success: false}, err
	}
	return &pb.DeleteTodoResponse{Success: true}, nil
}

func (s *todoServer) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	log.Printf("Received UpdateTodo request: ID %d, Completed %t", req.Id, req.Completed)
	todoID := req.Id
	todoState := req.Completed
	err := s.todoRepo.UpdateTodo(todoID, todoState)
	if err != nil {
		return &pb.UpdateTodoResponse{Success: false}, err
	}
	return &pb.UpdateTodoResponse{Success: true}, nil
}

func convertToRepoTodo(grpcTodo *pb.Todo) todorepo.Todo {
	return todorepo.Todo{
		ID:        grpcTodo.Id,
		Title:     grpcTodo.Title,
		Completed: grpcTodo.Completed,
	}
}

func convertToGRPC(repoTodo todorepo.Todo) *pb.Todo {
	return &pb.Todo{
		Id:        repoTodo.ID,
		Title:     repoTodo.Title,
		Completed: repoTodo.Completed,
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo, err := todorepo.NewTodoRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoServer{
		todoRepo: repo,
	})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
