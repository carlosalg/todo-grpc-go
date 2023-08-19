package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	pb "github.com/carlosalg/todo-grpc-go/api"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Variable to indicate the IP address to connect to the server
var (
	addr         = flag.String("addr", "localhost:50051", "the address to connect to")
	deleteID     = flag.Int("d", 0, "ID of the todo to delete")
	updateID     = flag.Int("u", 0, "ID of the todo to update")
	updateStatus = flag.Bool("status", false, "New completion status (true/false)")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Establish  a gRPC connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection
	c := pb.NewTodoServiceClient(conn)

	// Create a context with a timeout of one second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Check for arguments in the command-line and creates a new Todo using the CreateTodo method
	if len(flag.Args()) >= 2 {
		idInt, _ := strconv.Atoi(flag.Args()[0])
		id := int32(idInt)
		title := flag.Args()[1]
		completed := false

		if len(flag.Args()) >= 3 && flag.Args()[2] == "true" {
			completed = true
		}
		r, err := c.CreateTodo(ctx, &pb.Todo{Id: id, Title: title, Completed: completed})
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		log.Printf("Task: [%d] %s, completed: %t", r.GetId(), r.GetTitle(), r.GetCompleted())
	}

	// Initiate a stream to receive a list of Todo items
	stream, err := c.GetTodoList(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error calling GetTodoList: %v", err)
	}

	// Receive and process the stream of Todo messages
	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			// stream ended
			break
		}

		if err != nil {
			log.Fatalf("Error  receiving todo: %v", err)
		}

		// Print the received Todo item
		fmt.Printf("Received todo: %v\n", todo)
	}

	if *deleteID != 0 {
		deleteCtx, deleteCancel := context.WithTimeout(context.Background(), time.Second)
		defer deleteCancel()

		deleteResp, deleteErr := c.DeleteTodo(deleteCtx, &pb.DeleteTodoRequest{Id: int32(*deleteID)})
		if deleteErr != nil {
			log.Fatalf("Error deleting todo: %v", deleteErr)
		}

		if deleteResp.Success {
			log.Printf("Todo deleted successfully")
		} else {
			log.Printf("Todo deletion was unsuccessfully")
		}

	}

	if *updateID != 0 {
		updateCtx, updateCancel := context.WithTimeout(context.Background(), time.Second)
		defer updateCancel()

		updateResp, updateErr := c.UpdateTodo(updateCtx, &pb.UpdateTodoRequest{
			Id:        int32(*updateID),
			Completed: *updateStatus})
		if updateErr != nil {
			log.Fatalf("Error updating todo: %v", updateErr)
		}

		if updateResp.Success {
			log.Printf("Todo updated successfully")
		} else {
			log.Printf("Todo update was unsuccessfully")
		}
	}

}
