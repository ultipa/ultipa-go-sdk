package sdk

import (
	"context"
	"fmt"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

type updateNodeReturns struct {
}

// UpdateNodes update node data to db
func UpdateNodes(client ultipa.UltipaRpcsClient, nodes []utils.Node) updateNodeReturns {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	newNodes := utils.ToRpcNodes(nodes)

	var Nodes []*ultipa.ModifyNode

	for _, n := range newNodes {
		var Node ultipa.ModifyNode

		Node.Id = n.Id
		for _, v := range n.Values {
			var value ultipa.ModifyValues
			value.Key = v.Key
			value.Value = v.Value
			Node.Values = append(Node.Values, &value)
		}
		Nodes = append(Nodes, &Node)
	}

	// fmt.Printf("updateInf : %v : %v \n", newNodes, Nodes)

	msg, err := client.Modify(ctx, &ultipa.ModifyRequest{
		Nodes: Nodes,
	})

	if err != nil {
		log.Fatalf("update node error %v", err)
	}

	fmt.Printf("%v", msg)

	return updateNodeReturns{}
}
