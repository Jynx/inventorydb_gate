package grpcapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jynx/inventorydb-gate/pb"
)

type DbClient struct {
	client *mongo.Client
	pb.UnimplementedInventoryDBGateServiceServer
}

func NewInventoryDbServer(ctx context.Context, opts *options.ClientOptions) (*DbClient, error) {
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &DbClient{client: client}, nil
}
