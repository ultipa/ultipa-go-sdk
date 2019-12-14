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
	NodeCount int32
	EdgeCount int32
}

// Statistic returns the node and edge count number from server
func Statistic(client ultipa.UltipaRpcsClient) *StatisticRes {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	msg, err := client.DbInformation(ctx, &ultipa.DbInformationRequest{})
	if err != nil {
		log.Printf("db info error %v", err)
	}

	return &StatisticRes{
		NodeCount: msg.TotalNodes,
		EdgeCount: msg.TotalEdges,
	}
}
