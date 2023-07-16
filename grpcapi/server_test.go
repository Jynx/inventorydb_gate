package inventorydb

import (
	"context"
	"testing"

	"github.com/jynx/inventorydb-gate/pb"
)

type MockDocumentDataStore struct {
	// some fake document DB stuff here, perhaps a map?
}

func (db MockDocumentDataStore) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetInventoryResponse, error) {
	response := &pb.GetInventoryResponse{
		Inventory: &pb.Inventory{
			PlayerId: "1",
			Items:    []*pb.Item{{Id: "1", Name: "Axe"}},
		},
	}
	return response, nil
}

func TestDbClient_GetInventory(t *testing.T) {

}
