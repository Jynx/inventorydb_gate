package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"playerInventory/inventorydb"
	"time"

	pb "github.com/Jynx/inventoryProtos/inventory"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	listenAddr := ":50051"
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://hotdegs:.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := inventorydb.NewInventoryDbClient(ctx, opts)
	server := grpc.NewServer()
	pb.RegisterInventoryServiceServer(server, client)

	log.Printf("gRPC server listening on %s", listenAddr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	resp, createErr := client.CreateInventory(ctx, &pb.CreateInventoryRequest{})
	if createErr != nil {
		log.Fatalf("failed to create inventory: %v", createErr)
	}
	fmt.Print(resp)
}
