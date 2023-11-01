package main

import (
	"context"
	"log"
	"net"

	pb "user/user" // Update with the correct import path for userpb

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	users []pb.User
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserById(ctx context.Context, request *pb.UserRequest) (*pb.User, error) {
	for _, usr := range s.users {
		if usr.Id == request.Id {
			return &usr, nil
		}
	}
	return nil, status.Error(codes.NotFound, "User not found") // Import "google.golang.org/grpc/status" and "google.golang.org/grpc/codes"
}

func (s *server) GetUsersByIds(ctx context.Context, request *pb.UsersRequest) (*pb.UsersResponse, error) {
	var usersResponse pb.UsersResponse
	for _, id := range request.Ids {
		for _, user := range s.users {
			if user.Id == id {
				usersResponse.Users = append(usersResponse.Users, &user)
			}
		}
	}
	return &usersResponse, nil
}

func main() {
	user := []pb.User{
		{
			Id:      1,
			Fname:   "Steve",
			City:    "LA",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
		// Add more users as needed
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: user}) // Remove the "+" before "server"

	log.Println("Server is listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
