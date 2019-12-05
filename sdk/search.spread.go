package sdk

import (
	"context"
	// "google.golang.org/grpc"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

// SearchABRes is the SearchAB value struct

// SearchABRequest is the struct
type spreadRequest struct {
	Src     string
	Limit   int32
	Depth   int32
	Selects []string
	Type    string
	Turbo   bool //BFS or DFS
}

func NewSpreadRequest(src string) spreadRequest {
	return spreadRequest{src, 3, 2, []string{"name"}, "BFS", false}
}

// SpreadResponse is the struct
type SpreadResponse struct {
	TotalCost  int32
	EngineCost int32
	Paths      utils.Paths
}

// Spread returns paths of spread one point
func Spread(client ultipa.UltipaRpcsClient, request spreadRequest) SpreadResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.NodeSpread(ctx, &ultipa.NodeSpreadRequest{
		Source:        request.Src,
		Limit:         request.Limit,
		Depth:         request.Depth,
		SelectColumns: request.Selects,
	})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	paths := utils.FormatPathsFromSpread(msg.Paths)

	return SpreadResponse{
		msg.TotalTimeCost,
		msg.EngineTimeCost,
		paths,
	}
}
