package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-chi-rest-app/repository"
	"go-chi-rest-app/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type Server struct {
	httpServer *http.Server
}

type ServerOption func(*Server)

func WithPort(port string) ServerOption {
	return func(s *Server) {
		s.httpServer.Addr = fmt.Sprintf(":%s", port)
	}
}

func NewServer(router chi.Router, opts ...ServerOption) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr:    ":3000", // default port
			Handler: router,
		},
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *Server) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", s.httpServer.Addr, err)
		}
	}()
	log.Printf("Server is ready to handle requests at %s\n", s.httpServer.Addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // default port if not set in .env
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Create an instance of OwnerRepository
	ownerRepo := repository.NewOwnerRepository(redisClient)

	// Initialize router and routes
	router := chi.NewRouter()
	routeHandler := routes.NewRouteHandler(ownerRepo)
	routeHandler.RegisterRoutes(router)

	// Initialize server with functional options
	server := NewServer(router, WithPort(port))

	// Start the server
	server.Start()

	// Graceful shutdown on SIGINT or SIGTERM
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}