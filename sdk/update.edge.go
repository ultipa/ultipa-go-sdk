package sdk

import (
	"context"
	"fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

type updateEdgeReturns struct {
}

// UpdateEdges update Edge data to db
func UpdateEdges(client ultipa.UltipaRpcsClient, edges []utils.Edge) updateEdgeReturns {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	newEdges := utils.ToRpcEdges(edges)

	var Edges []*ultipa.ModifyEdge

	for _, n := range newEdges {
		var Edge ultipa.ModifyEdge

		Edge.Id = n.Id
		Edge.FromId = n.FromId
		Edge.ToId = n.ToId
		for _, v := range n.Values {
			var value ultipa.ModifyValues
			value.Key = v.Key
			value.Value = v.Value
			Edge.Values = append(Edge.Values, &value)
		}
		Edges = append(Edges, &Edge)
	}

	// fmt.Printf("updateInf : %v : %v \n", newEdges, Edges)

	msg, err := client.Modify(ctx, &ultipa.ModifyRequest{
		Edges: Edges,
	})

	if err != nil {
		log.Fatalf("update edge error %v", err)
	}

	fmt.Printf("%v", msg)

	return updateEdgeReturns{}
}
