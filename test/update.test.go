package main

import (
	"fmt"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/utils"
)

// import "encoding/json"

// import ultipa "ultipa-go-sdk/rpc"

func main() {
	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60064")

	defer conn.Close()

	fmt.Println("Starting test update module")

	// nodeReq := sdk.NewSearchNodesRequest()
	// nodeReq.ID = "123"
	// nodeRes := sdk.SearchNodes(client, nodeReq)
	// node := *nodeRes.Nodes[0]
	// fmt.Printf("%v\n", node["fullName"])
	// updateNode := utils.Node{"_id": "123", "name": "test"}

	// uNMsg := sdk.UpdateNodes(client, []utils.Node{updateNode})
	// fmt.Printf("%v, %v\n", uNMsg.Status.ErrorCode, uNMsg.Status.Msg)

	edgeReg := sdk.NewSearchEdgesRequest()
	edgeReg.ID = "123"
	edgeMsg := sdk.SearchEdges(client, edgeReg)
	updateEdge := utils.Edge{"_id": "123", "name": "test123"}
	uEMsg := sdk.UpdateEdges(client, []utils.Edge{updateEdge})
	fmt.Printf("%#v \n %+v \n", edgeMsg.Edges, uEMsg)
}
