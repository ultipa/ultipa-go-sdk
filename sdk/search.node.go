package sdk

import (
	"context"
	// "google.golang.org/grpc"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/utils"
	ultipa "ultipa-go-sdk/rpc"
)

// string begin_id = 1;
// int32 limit  = 2;
// Filter node_filter = 3;
// repeated string select_columns = 4;

type searchNodesRequest struct {
	ID         string
	NodeFilter ultipa.Filter
	Limit      int32
	Select     []string
}

type searchNodesResponse struct {
	TotalCost int32
	Count     int32
	Nodes     []utils.Node
}

func NewSearchNodesRequest() searchNodesRequest {
	return searchNodesRequest{"", ultipa.Filter{}, 10, []string{"name"}}
}

func SearchNodes(client ultipa.UltipaRpcsClient, request searchNodesRequest) searchNodesResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchNodes(ctx, &ultipa.SearchNodesRequest{
		BeginId:       request.ID,
		Limit:         request.Limit,
		NodeFilter:    &request.NodeFilter,
		SelectColumns: request.Select,
	})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	// paths := utils.FormatPaths(msg.Paths)
	nodes := utils.FormatNodes(msg.Nodes)

	return searchNodesResponse{
		msg.TimeCost,
		msg.TotalCounts,
		nodes,
	}
}
