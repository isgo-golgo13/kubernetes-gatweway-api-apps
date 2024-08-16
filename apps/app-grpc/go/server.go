package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"go-grpc-app/proto"
)

type server struct {
	proto.UnimplementedSendServiceServer
}

func (s *server) Send(ctx context.Context, req *proto.SendRequest) (*proto.SendResponse, error) {
	return &proto.SendResponse{BytesSent: int32(len(req.Data))}, nil
}

func (s *server) SendWithTimeout(ctx context.Context, req *proto.SendWithTimeoutRequest) (*proto.SendResponse, error) {
	select {
	case <-time.After(time.Duration(req.Timeout) * time.Millisecond):
		return nil, fmt.Errorf("timeout exceeded")
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &proto.SendResponse{BytesSent: int32(len(req.Data))}, nil
	}
}

func (s *server) SendAll(stream proto.SendService_SendAllServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		stream.Send(&proto.SendResponse{BytesSent: int32(len(req.Data))})
	}
}

type serverOptions struct {
	port string
}

type ServerOption func(*serverOptions)

func WithPort(port string) ServerOption {
	return func(o *serverOptions) {
		o.port = port
	}
}

func NewServer(opts ...ServerOption) *grpc.Server {
	options := &serverOptions{
		port: "50051", // Default port
	}
	for _, opt := range opts {
		opt(options)
	}

	s := grpc.NewServer()
	srv := &server{}
	proto.RegisterSendServiceServer(s, srv)

	// Start the server
	go func() {
		lis, err := net.Listen("tcp", ":"+options.port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Printf("Server is running on port %s", options.port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received shutdown signal, gracefully stopping the server...")
		s.GracefulStop()
		os.Exit(0)
	}()

	return s
}

func main() {
	godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "50051" // default port
	}

	_ = NewServer(WithPort(port))
}
