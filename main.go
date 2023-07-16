package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jynx/inventorydb-gate/config"
	inventorydb "github.com/jynx/inventorydb-gate/grpcapi"
	"github.com/jynx/inventorydb-gate/pb"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file with error: %s", err)
	}
}

func main() {
	config := config.NewConfig()

	listenAddr := ":" + config.Server.Port
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	_, _, clientErr := NewInventoryDbClient(listenAddr, lis, config)
	if clientErr != nil {
		log.Fatalf("failed to create and start inventory server")
	}
}

// todo: Maybe decouple the data store choice here from the grpc server startup
func NewInventoryDbClient(listenAddr string, lis net.Listener, config *config.Config) (context.Context, *inventorydb.DbClient, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connStr := "mongodb+srv://" + config.MongoDb.Username + ":" + config.MongoDb.Password + "@cluster0.c1xy1s5.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := inventorydb.NewInventoryDbClient(ctx, opts)
	if err != nil {
		return nil, nil, err
	}

	server := grpc.NewServer()
	pb.RegisterInventoryDBGateServiceServer(server, client)
	reflection.Register(server)

	log.Printf("gRPC server listening on %s", listenAddr)
	if err := server.Serve(lis); err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}
