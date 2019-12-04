package sdk

import (
	"context"
	// "google.golang.org/grpc"
	// "fmt"
	"log"
	"time"
	// "ultipa-go-sdk/utils"
	ultipa "ultipa-go-sdk/rpc"
)

type khopRequest struct {
	Src   string
	Depth int32
	Limit int32
	turbo bool
}

func NewKhopRequest(src string) khopRequest {
	return khopRequest{src, 1, 50, false}
}

type node map[string]string

type KHopResponse struct {
	EngineCost int32
	TotalCost  int32
	Count      int32
	Nodes      []node
}

// Khop returns sample nodes and total counts of the src n hop neighbor
func SearchKhop(client ultipa.UltipaRpcsClient, request khopRequest) KHopResponse {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.SearchKhop(ctx, &ultipa.SearchKhopRequest{Source: request.Src, Limit: request.Limit, Depth: request.Depth})

	if err != nil {
		log.Fatalf("ab search error %v", err)
	}

	// int32 total_time_cost = 1;
	// int32 engine_time_cost = 2;
	// int32 total_number = 3;
	// repeated KhopRow nodes = 4;

	var newNodes []node
	for _, n := range msg.Nodes {
		newNode := make(node)
		for _, v := range n.Values {
			newNode[v.Key] = v.Value
		}
		newNodes = append(newNodes, newNode)
	}

	return KHopResponse{msg.EngineTimeCost, msg.TotalTimeCost, msg.TotalNumber, newNodes}
}
