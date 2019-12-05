package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

type CreateIndexResponse struct {
	TimeCost int32
	Status   ultipa.CreateIndexReply_STATUS
}

// DeleteNodes update node data to db
func createIndex(client ultipa.UltipaRpcsClient, dbType ultipa.CreateIndexRequest_DBType, propertyName string) CreateIndexResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.CreateIndex(ctx, &ultipa.CreateIndexRequest{
		Type:       dbType,
		ColumnName: propertyName,
	})

	if err != nil {
		log.Fatalf("[Error] delete node error: %v", err)
	}

	return CreateIndexResponse{
		TimeCost: msg.TimeCost,
		Status:   msg.Status,
	}
}

// CreateNodeIndex create index for node
func CreateNodeIndex(client ultipa.UltipaRpcsClient, propertyName string) CreateIndexResponse {
	return createIndex(client, ultipa.CreateIndexRequest_DBNODE, propertyName)
}

// CreateEdgeIndex create index for Edge
func CreateEdgeIndex(client ultipa.UltipaRpcsClient, propertyName string) CreateIndexResponse {
	return createIndex(client, ultipa.CreateIndexRequest_DBEDGE, propertyName)
}
