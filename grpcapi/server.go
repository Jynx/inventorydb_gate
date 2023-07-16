package grpcapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jynx/inventorydb-gate/pb"
)

type DbClient struct {
	client *mongo.Client
	pb.UnimplementedInventoryDBGateServiceServer
}

type Item struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string
	Type   string
	Damage int
}

type Inventory struct {
	Username string
	PlayerId string `bson:"player_id,omitempty"`
	Items    []Item `bson:"items,omitempty"`
}

func NewInventoryDbClient(ctx context.Context, opts *options.ClientOptions) (*DbClient, error) {
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &DbClient{client: client}, nil
}
