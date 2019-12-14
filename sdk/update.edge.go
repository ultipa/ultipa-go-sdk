package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

// UpdateEdges update Edge data to db
func UpdateEdges(client ultipa.UltipaRpcsClient, edges []utils.Edge) (reply *ultipa.ModifyReply, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	newEdges := utils.ToRPCEdges(edges)

	var Edges []*ultipa.Edge

	for _, n := range newEdges {
		var Edge ultipa.Edge

		Edge.Id = n.Id
		Edge.FromId = n.FromId
		Edge.ToId = n.ToId
		for _, v := range n.Values {
			var value ultipa.Value
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
		log.Printf("update edge error %v", err)
	}

	return msg, err
}
