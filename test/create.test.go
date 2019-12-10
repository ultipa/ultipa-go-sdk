package main

import (
	"fmt"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/utils"
)

func main() {
	fmt.Println("test create sdk")

	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60064")

	defer conn.Close()

	// createNodeProMsg := sdk.CreateNodeProperty(client, "test", sdk.PROPERTY_TYPE_INT)
	// fmt.Printf("Create Node Property, %v \n", createNodeProMsg)

	// createEdgeProMsg := sdk.CreateEdgeProperty(client, "test", sdk.PROPERTY_TYPE_INT)
	// fmt.Printf("Create Edge Property, %v \n", createEdgeProMsg)

	// newNode := utils.Node{
	// 	"name": "hello all",
	// }
	// createNodeMsg := sdk.CreateNodes(client, []utils.Node{newNode})
	// fmt.Printf("Create Nodes, IDs: %v ms, time cost %v ms \n", createNodeMsg.Ids, createNodeMsg.TimeCost)

	// newEdge := utils.Edge{
	// 	"_from_id": "12332",
	// 	"_to_id":   "999",
	// 	"name":     "test",
	// }

	// createEdgeMsg := sdk.CreateEdges(client, []utils.Edge{newEdge})
	// fmt.Printf("Create Edges, IDs: %v ms, time cost %v ms\n", createEdgeMsg.Ids, createNodeMsg.TimeCost)

}
