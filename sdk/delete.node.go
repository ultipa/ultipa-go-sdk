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
	Status   string
}

// DeleteNodes update node data to db
func DeleteNodes(client ultipa.UltipaRpcsClient, ids []string) deleteNodesResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.Delete(ctx, &ultipa.DeleteRequest{
		Type:      ultipa.DeleteRequest_DBNODE,
		DeleteIds: ids,
	})

	if err != nil {
		log.Fatalf("[Error] delete node error: %v", err)
	}

	status := "ok"

	if msg.Status == ultipa.DeleteReply_FAILED {
		status = "failed"
	}
	return deleteNodesResponse{
		TimeCost: msg.TimeCost,
		Status:   status,
	}
}
