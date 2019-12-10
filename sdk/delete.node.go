package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type deleteNodesResponse struct {
	TimeCost int32
	Status   *ultipa.Status
}

// DeleteNodes update node data to db
func DeleteNodes(client ultipa.UltipaRpcsClient, ids []string) deleteNodesResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.Delete(ctx, &ultipa.DeleteRequest{
		Type:      ultipa.DBType_DBNODE,
		DeleteIds: ids,
	})

	if err != nil {
		log.Fatalf("[Error] delete node error: %v", err)
	}

	return deleteNodesResponse{
		TimeCost: msg.TimeCost,
		Status:   msg.Status,
	}
}
