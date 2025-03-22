package server

import (
	"context"
	"log"
	"net"

	"taxiya/internal/auth"
	pb "taxiya/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	authService *auth.AuthService
}

func NewServer() *Server {
	return &Server{
		authService: auth.NewAuthService(),
	}
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.authService.Login(req.Email, req.Password)
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return s.authService.Register(req.Email, req.Password, req.Name, req.Phone)
}

func StartServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, NewServer())
	reflection.Register(s)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
