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

// SearchABRes is the SearchAB value struct
type SearchABRes struct {
	nodeCount int32
	edgeCount int32
}

// SearchABRequest is the struct
type SearchABRequest struct {
	Src                  string
	Dest                 string
	Limit                int32
	Depth                int32
	Shortest             bool
	SelectNodeProperties []string
	SelectEdgeProperties []string
	NodeFilter           utils.Filter
	EdgeFilter           utils.Filter
	Turbo                bool
}

func NewABRequest(src string, dest string) SearchABRequest {
	return SearchABRequest{src, dest, 5, 1, false, []string{"name"}, []string{"name"}, utils.Filter{}, utils.Filter{}, false}
}

// SearchABResponse is the struct
type SearchABResponse struct {
	TotalCost  int32
	EngineCost int32
	Paths      utils.Paths
}

// SearchAB returns paths of two points
func SearchAB(client ultipa.UltipaRpcsClient, request SearchABRequest) SearchABResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchAB(ctx, &ultipa.SearchABRequest{
		Source:               request.Src,
		Dest:                 request.Dest,
		Limit:                request.Limit,
		Depth:                request.Depth,
		SelectNodeProperties: request.SelectNodeProperties,
		SelectEdgeProperties: request.SelectEdgeProperties,
		NodeFilter:           &request.NodeFilter,
		EdgeFilter:           &request.EdgeFilter,
		ByShortSearch:        request.Shortest,
		PassSuperNode:        request.Turbo,
	})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	paths := utils.FormatPaths(msg.Paths)

	return SearchABResponse{
		msg.TotalTimeCost,
		msg.EngineTimeCost,
		paths,
	}
}
