package sdk

import (
	"context"
	// "google.golang.org/grpc"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/pkg"
	ultipa "ultipa-go-sdk/rpc"
)

// SearchABRes is the SearchAB value struct
type SearchABRes struct {
	nodeCount int32
	edgeCount int32
}

// SearchABRequest is the struct
type SearchABRequest struct {
	Src      string
	Dest     string
	Limit    int32
	Depth    int32
	shortest bool
	Selects []string
	Turbo    bool
}

func NewABRequest(src string, dest string) SearchABRequest {
	return SearchABRequest{src, dest, 3, 2, false, []string{"name"}, false}
}

// SearchABResponse is the struct
type SearchABResponse struct {
	TotalCost  int32
	EngineCost int32
	Paths      pkg.Paths
}

// SearchAB returns paths of two points
func SearchAB(client ultipa.UltipaRpcsClient, request SearchABRequest) SearchABResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchAB(ctx, &ultipa.SearchABRequest{
		Source: request.Src, 
		Dest: request.Dest, 
		Limit: request.Limit, 
		Depth: request.Depth,
		SelectColumns: request.Selects,
	})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	paths := pkg.FormatPaths(msg.Paths)

	return SearchABResponse{
		msg.TotalTimeCost,
		msg.EngineTimeCost,
		paths,
	}
}
