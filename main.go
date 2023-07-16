package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jynx/inventorydb-gate/config"
	"github.com/jynx/inventorydb-gate/grpcapi"
	"github.com/jynx/inventorydb-gate/pb"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file with error: %s", err)
	}
}

func main() {
	config := config.NewConfig()

	go func() {
		err := StartGrpcGatewayServer(":"+config.Server.HTTPPort, config)
		if err != nil {
			log.Fatalf("failed to create and start inventory http server: %s", err)
		}
	}()

	err := StartGrpcServer(":"+config.Server.GRPCPort, config)
	if err != nil {
		log.Fatalf("failed to create and start inventory grpc server: %s", err)
	}
}

func StartGrpcServer(listenAddr string, config *config.Config) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connStr := "mongodb+srv://" + config.MongoDb.Username + ":" + config.MongoDb.Password + "@cluster0.c1xy1s5.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := grpcapi.NewInventoryDbServer(ctx, opts)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterInventoryDBGateServiceServer(server, client)
	reflection.Register(server)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on %s", listenAddr)
	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func StartGrpcGatewayServer(listenAddr string, config *config.Config) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connStr := "mongodb+srv://" + config.MongoDb.Username + ":" + config.MongoDb.Password + "@cluster0.c1xy1s5.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	server, err := grpcapi.NewInventoryDbServer(ctx, opts)
	if err != nil {
		return err
	}

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterInventoryDBGateServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	log.Printf("http gateway server listening on %s", listenAddr)
	if err := http.Serve(lis, mux); err != nil {
		return err
	}

	return nil
}
