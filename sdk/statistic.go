package sdk

import (
	"context"
	// "google.golang.org/grpc"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

// StatisticRes is the Statistic value struct
type StatisticRes struct {
	nodeCount int32
	edgeCount int32
}

// Statistic returns the node and edge count number from server
func Statistic(client ultipa.UltipaRpcsClient) StatisticRes {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	msg, err := client.DbInformation(ctx, &ultipa.DbInformationRequest{})
	if err != nil {
		log.Fatalf("db info error %v", err)
	}

	return StatisticRes{nodeCount: msg.TotalNodes, edgeCount: msg.TotalEdges}
}
