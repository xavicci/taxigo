// main.go
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/xavicci/taxisgo/proto/auth"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Aquí iría la lógica de autenticación
	if req.Username == "user" && req.Password == "pass" {
		return &pb.LoginResponse{
			Success: true,
			Message: "Login successful",
			Token:   "someauthtoken",
		}, nil
	}
	return &pb.LoginResponse{
		Success: false,
		Message: "Invalid credentials",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
