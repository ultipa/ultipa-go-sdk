package sdk

import (
	"context"
	// "fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

// message InsertRequest {
//   repeated InsertEdge Edges = 1;
//   repeated InsertEdge edges = 2;
// }

// message InsertReply {
//   enum STATUS{
//     OK = 0;
//     FAILED = 1;
//   }
//   STATUS status = 1;
//   int32 time_cost = 2;
//   repeated int32 ids = 3;
// }

// CreateEdges update Edge data to db
func CreateEdges(client ultipa.UltipaRpcsClient, edges []utils.Edge) *ultipa.InsertReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	newEdges := utils.ToRpcEdges(edges)

	var Edges []*ultipa.InsertEdge

	for _, e := range newEdges {
		var Edge ultipa.InsertEdge

		Edge.FromId = e.FromId
		Edge.ToId = e.ToId

		if Edge.FromId == "" || Edge.ToId == "" {
			log.Fatalf("[Error] create Edge error: %v, fromId or toId is missing ", Edge)
			continue
		}

		for _, v := range e.Values {
			var value ultipa.InsertValues
			value.Key = v.Key
			value.Value = v.Value
			Edge.Values = append(Edge.Values, &value)
		}
		Edges = append(Edges, &Edge)
	}

	msg, err := client.Insert(ctx, &ultipa.InsertRequest{
		Edges: Edges,
	})

	if err != nil {
		log.Fatalf("[Error] create Edge error: %v", err)
	}

	return msg

}
