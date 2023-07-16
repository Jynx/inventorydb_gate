package inventorydb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func (db DbClient) CreateInventory(ctx context.Context, request *pb.CreateInventoryRequest) (*pb.CreateInventoryResponse, error) {
	collection := db.client.Database("player_inventory").Collection("inventory")

	// input validation
	newInventory := Inventory{
		Username: "steve2",
		PlayerId: "12345678",
	}

	_, err := collection.InsertOne(ctx, newInventory)
	if err != nil {
		fmt.Printf("error from insert: %s", err)
	}

	response := &pb.CreateInventoryResponse{
		Inventory: &pb.Inventory{
			PlayerId: "1",
			Items:    []*pb.Item{{Id: "1", Name: "Axe"}},
		},
	}
	return response, nil
}

func (db DbClient) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetInventoryResponse, error) {
	collection := db.client.Database("player_inventory").Collection("inventory")
	filter := bson.D{{"PlayerId", req.PlayerId}}

	var result Inventory
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// return special rpc error?
		}
		return nil, err
	}

	items := make([]*pb.Item, len(result.Items))
	for i, item := range result.Items {
		items[i] = &pb.Item{
			Id:   item.Id.String(),
			Name: item.Name,
		}
	}

	response := &pb.GetInventoryResponse{
		Inventory: &pb.Inventory{
			PlayerId: result.PlayerId,
			Items:    items,
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
