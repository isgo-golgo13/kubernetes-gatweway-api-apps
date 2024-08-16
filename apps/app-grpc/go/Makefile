.PHONY: proto build-server build-client build run

# Compile Protocol Buffers and gRPC files
proto:
	protoc --go_out=proto --go-grpc_out=proto proto/send.proto

# Build the server binary
build-server:
	go build -o bin/server ./server.go

# Build the client binary
build-client:
	go build -o bin/client ./client.go

# Build both server and client binaries
build: proto build-server build-client

# Run the application using Docker Compose
run:
	docker-compose up --build

# Clean the build directory
clean:
	rm -rf bin/*
