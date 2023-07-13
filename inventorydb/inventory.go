package inventorydb

import (
	"context"
	"fmt"

	pb "github.com/Jynx/inventoryProtos/inventory"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbClient struct {
	client *mongo.Client
	pb.UnimplementedInventoryServiceServer
}

type Inventory struct {
	Username string
	PlayerId string        `bson:"player_id,omitempty"`
	Items    []interface{} `bson:"items,omitempty"`
}

func (db DbClient) CreateInventory(ctx context.Context, request *pb.CreateInventoryRequest) (*pb.GetInventoryResponse, error) {
	collection := db.client.Database("player_inventory").Collection("inventory")

	// input validation
	newInventory := Inventory{
		Username: "steve2",
		PlayerId: "12345678",
	}

	result, err := collection.InsertOne(ctx, newInventory)
	if err != nil {
		fmt.Printf("error from insert: %s", err)
	}

	response := &pb.GetInventoryResponse{
		Inventory: &pb.Inventory{
			PlayerId: "1",
			Items:    []*pb.Item{{Id: "1", Name: "Axe"}},
		},
	}
	// get rid of this
	fmt.Println(result)
	return response, nil
}

func (db DbClient) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetInventoryResponse, error) {
	response := &pb.GetInventoryResponse{
		Inventory: &pb.Inventory{
			PlayerId: "1",
			Items:    []*pb.Item{{Id: "1", Name: "Axe"}},
		},
	}
	return response, nil
}

func NewInventoryDbClient(ctx context.Context, opts *options.ClientOptions) (*DbClient, error) {
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &DbClient{client: client}, nil
}
