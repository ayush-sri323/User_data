package main

import (
	"context"
	"fmt"
	"log"
	pb "user/user"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewUserServiceClient(conn)

	// Example: Fetch user by ID
	user, err := client.GetUserById(context.Background(), &pb.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error calling GetUserById: %v", err)
	}
	fmt.Printf("User fetched by ID: %v\n", user)

	// Example: Fetch multiple users by IDs
	ids := []int32{1, 2}
	usersResponse, err := client.GetUsersByIds(context.Background(), &pb.UsersRequest{Ids: ids})
	if err != nil {
		log.Fatalf("Error calling GetUsersByIds: %v", err)
	}
	fmt.Printf("Users fetched by IDs: %v\n", usersResponse.Users)
}
