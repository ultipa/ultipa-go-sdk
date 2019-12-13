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

type khopRequest struct {
	Src                  string
	Limit                int32
	Depth                int32
	SelectNodeProperties []string
	NodeFilter           utils.Filter
	EdgeFilter           utils.Filter
}

func NewKhopRequest(src string) khopRequest {
	return khopRequest{src, 10, 1, []string{"name"}, utils.Filter{}, utils.Filter{}}
}

type KHopResponse struct {
	EngineCost int32
	TotalCost  int32
	Count      int32
	Nodes      []*utils.Node
}

// SearchKhop returns sample nodes and total counts of the src n hop neighbor
func SearchKhop(client ultipa.UltipaRpcsClient, request khopRequest) KHopResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchKhop(ctx, &ultipa.SearchKhopRequest{
		Source:               request.Src,
		Limit:                request.Limit,
		Depth:                request.Depth,
		SelectNodeProperties: request.SelectNodeProperties,
		NodeFilter:           &request.NodeFilter,
		EdgeFilter:           &request.EdgeFilter,
	})

	if err != nil {
		log.Fatalf("khop search error %v", err)
	}

	var newNodes []*utils.Node

	newNodes = utils.FormatNodes(msg.Nodes)

	return KHopResponse{msg.EngineTimeCost, msg.TotalTimeCost, msg.TotalNumber, newNodes}
}
