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

	msg := sdk.Statistic(client)

	fmt.Printf("%+v\n", msg)
	fmt.Print("\n=================AB Search=====================\n\n")

	abReq := sdk.NewABRequest("123", "321")
	abReq.Depth = 7
	abReq.Limit = 1
	abMsg := sdk.SearchAB(client, abReq)
	for i, path := range abMsg.Paths {
		fmt.Printf("Path[%v] : ", i)
		for i := 0; i < len(path.Nodes); i++ {
			node := *path.Nodes[i]
			fmt.Printf("%v", node["name"])

			if i < len(path.Nodes)-1 {
				edge := *path.Edges[i]
				fmt.Printf(" - [%v] - ", edge["name"])
			}
		}
		fmt.Print("\n\n--------------------------------------\n\n")
	}

	fmt.Printf("engine cost %v ms, total cost %v ms \n", abMsg.EngineCost, abMsg.TotalCost)

	fmt.Print("\n======================================\n")

	fmt.Print("\n=================Search Khop=====================\n\n")

	khopReq := sdk.NewKhopRequest("123")
	khopMsg := sdk.SearchKhop(client, khopReq)

	for _, n := range khopMsg.Nodes {
		node := *n
		fmt.Printf(" [%v] ", node["name"])
	}

	fmt.Print("\n\n--------------------------------------\n\n")
	fmt.Printf("engine cost %v ms, total cost %v ms , total num : %v ms\n ", khopMsg.EngineCost, khopMsg.TotalCost, khopMsg.Count)

	fmt.Print("\n=================Search Nodes=====================\n\n")

	nodeReq := sdk.NewSearchNodesRequest()
	// nodeReq.ID = "123"

	filterConds := utils.NewFilterCondition("age", ">", []string{"20"})

	fmt.Printf("%v \n", filterConds)

	filter := utils.NewFilter("AND", filterConds)

	nodeReq.NodeFilter = filter

	nodeReq.SelectNodeProperties = []string{"name", "age"}

	nodeReq.Limit = 100

	nodeMsg := sdk.SearchNodes(client, nodeReq)

	for _, v := range nodeMsg.Nodes {
		node := *v
		fmt.Printf(" [%v | %v]", node["name"], node["age"])
	}

	fmt.Print("\n\n--------------------------------------\n\n")
	fmt.Printf("total cost %v ms , total num : %v \n", nodeMsg.TotalCost, nodeMsg.Count)

	fmt.Print("\n=================Search Edges=====================\n\n")

	edgeReq := sdk.NewSearchEdgesRequest()
	// edgeReq.ID = "123"

	filterConds = utils.NewFilterCondition("name", "=", []string{"Like"})

	fmt.Printf("%v \n", filterConds)

	filter = utils.NewFilter("AND", filterConds)

	edgeReq.EdgeFilter = filter

	edgeReq.SelectEdgeProperties = []string{"name"}

	edgeReq.Limit = 100

	edgeMsg := sdk.SearchEdges(client, edgeReq)

	for _, v := range edgeMsg.Edges {
		edge := *v
		fmt.Printf(" [ id: %v | type: %v] ", edge["_id"], edge["name"])
	}

	fmt.Print("\n\n--------------------------------------\n\n")
	fmt.Printf("total cost %v ms , total num : %v \n", edgeMsg.TotalCost, edgeMsg.Count)

	fmt.Print("\n=================Spread Node=====================\n\n")

	spreadRequest := sdk.NewSpreadRequest("123")
	spreadRequest.Limit = 20
	spreadRequest.Depth = 3
	spreadMsg := sdk.Spread(client, spreadRequest)

	for i, path := range spreadMsg.Paths {
		fmt.Printf("Path[%v] : ", i)
		for i := 0; i < len(path.Nodes); i++ {
			node := *path.Nodes[i]
			fmt.Printf("%v", node["name"])

			if i < len(path.Nodes)-1 {
				edge := *path.Edges[i]
				fmt.Printf(" - [%v] - ", edge["name"])
			}
		}
		fmt.Print("\n\n--------------------------------------\n\n")
	}

	fmt.Printf("total cost %v ms , engine cost : %v \n", spreadMsg.TotalCost, spreadMsg.EngineCost)

	fmt.Print("\n======================================\n\n")

}
