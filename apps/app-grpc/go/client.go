package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"go-grpc-app/proto"
)

func main() {
	godotenv.Load()

	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = "localhost:50051" // Default address
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewSendServiceClient(conn)

	sendData(client, 1, []byte("Hello"))
	sendDataWithTimeout(client, 2, []byte("Hello with timeout"), 1000)
	sendAllData(client, [][]byte{[]byte("Stream 1"), []byte("Stream 2")})
}

func sendData(client proto.SendServiceClient, id int32, data []byte) {
	res, err := client.Send(context.Background(), &proto.SendRequest{Id: id, Data: data})
	if err != nil {
		log.Fatalf("Error calling Send: %v", err)
	}
	log.Printf("Bytes sent: %d\n", res.BytesSent)
}

func sendDataWithTimeout(client proto.SendServiceClient, id int32, data []byte, timeout int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()

	res, err := client.SendWithTimeout(ctx, &proto.SendWithTimeoutRequest{Id: id, Data: data, Timeout: timeout})
	if err != nil {
		log.Printf("Error calling SendWithTimeout: %v", err)
		return
	}
	log.Printf("Bytes sent: %d\n", res.BytesSent)
}

func sendAllData(client proto.SendServiceClient, dataArray [][]byte) {
	stream, err := client.SendAll(context.Background())
	if err != nil {
		log.Fatalf("Error calling SendAll: %v", err)
	}

	for _, data := range dataArray {
		if err := stream.Send(&proto.SendRequest{Id: 0, Data: data}); err != nil {
			log.Fatalf("Error sending data: %v", err)
		}
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		log.Printf("Bytes sent: %d\n", res.BytesSent)
	}
}
