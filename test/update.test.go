package main

import (
	"fmt"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/utils"
)

// import "encoding/json"

// import ultipa "ultipa-go-sdk/rpc"

func main() {
	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60062")

	defer conn.Close()

	fmt.Println("Starting test update module")

	// nodeReq := sdk.NewSearchNodesRequest()
	// nodeReq.ID = "123"
	// nodeRes := sdk.SearchNodes(client, nodeReq)
	// fmt.Printf("%v\n", nodeRes.Nodes[0]["fullName"])
	// updateNode := utils.Node{"_id": "123", "fullName": "test"}

	// uNMsg := sdk.UpdateNodes(client, []utils.Node{updateNode})
	// fmt.Printf("%v\n", uNMsg)

	edgeReg := sdk.NewSearchEdgesRequest()
	edgeReg.ID = "123"
	edgeMsg := sdk.SearchEdges(client, edgeReg)
	updateEdge := utils.Edge{"_id": "123", "phase": "test"}
	uEMsg := sdk.UpdateEdges(client, []utils.Edge{updateEdge})
	fmt.Printf("%v \n %v \n", edgeMsg, uEMsg)
}
