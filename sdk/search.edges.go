package sdk

import (
	"context"
	// "google.golang.org/grpc"
	// "fmt"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

// string begin_id = 1;
// int32 limit  = 2;
// Filter Edge_filter = 3;
// repeated string select_columns = 4;

type searchEdgesRequest struct {
	ID                   string
	EdgeFilter           ultipa.Filter
	Limit                int32
	SelectEdgeProperties []string
}

type searchEdgesResponse struct {
	TotalCost int32
	Count     int32
	Edges     []*utils.Edge
}

func NewSearchEdgesRequest() searchEdgesRequest {
	return searchEdgesRequest{"", ultipa.Filter{}, 10, []string{"name"}}
}

// SearchEdges returns edges by query
func SearchEdges(client ultipa.UltipaRpcsClient, request searchEdgesRequest) searchEdgesResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchEdges(ctx, &ultipa.SearchEdgesRequest{
		BeginId:          request.ID,
		Limit:            request.Limit,
		EdgeFilter:       &request.EdgeFilter,
		SelectProperties: request.SelectEdgeProperties,
	})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	// paths := utils.FormatPaths(msg.Paths)
	Edges := utils.FormatEdges(msg.Edges)

	return searchEdgesResponse{
		TotalCost: msg.TimeCost,
		Count:     msg.TotalCounts,
		Edges:     Edges,
	}
}
