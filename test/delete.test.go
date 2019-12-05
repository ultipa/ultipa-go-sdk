package main

import (
	"fmt"
	// "ultipa-go-sdk/pkg"
	"ultipa-go-sdk/sdk"
)

func main() {
	fmt.Println("test delete sdk")

	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60062")

	defer conn.Close()

	nodeDelMsg := sdk.DeleteNodes(client, []string{"999"})
	fmt.Printf("Delete Node 999, TimeCost: %v [%v] \n", nodeDelMsg.TimeCost, nodeDelMsg.Status)

	edgeDelMsg := sdk.DeleteEdges(client, []string{"999"})
	fmt.Printf("Delete Edge 999, TimeCost: %v [%v] \n", edgeDelMsg.TimeCost, edgeDelMsg.Status)

}
