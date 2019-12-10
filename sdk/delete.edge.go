package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type deleteEdgesResponse struct {
	TimeCost int32
	Status   *ultipa.Status
}

// DeleteEdges update node data to db
func DeleteEdges(client ultipa.UltipaRpcsClient, ids []string) deleteEdgesResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.Delete(ctx, &ultipa.DeleteRequest{
		Type:      ultipa.DBType_DBEDGE,
		DeleteIds: ids,
	})

	if err != nil {
		log.Fatalf("[Error] delete edge error: %v", err)
	}

	return deleteEdgesResponse{
		TimeCost: msg.TimeCost,
		Status:   msg.Status,
	}
}
