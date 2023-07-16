package grpcapi

import (
	"context"
	"fmt"

	"github.com/jynx/inventorydb-gate/pb"
)

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
