package grpcapi

import (
	"context"

	"github.com/jynx/inventorydb-gate/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

type GetInventoryFilter struct {
	PlayerId string `bson:"player_id"`
}

func (db DbClient) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetInventoryResponse, error) {
	collection := db.client.Database("player_inventory").Collection("inventory")
	filter := GetInventoryFilter{PlayerId: req.PlayerId}

	filterDoc, err := bson.Marshal(filter)
	if err != nil {
		//
	}

	var result Inventory
	err = collection.FindOne(ctx, filterDoc).Decode(&result)
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
