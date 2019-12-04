package main

import (
	"fmt"
	// "ultipa-go-sdk/pkg"
	"ultipa-go-sdk/sdk"
)

// import "encoding/json"

// import ultipa "ultipa-go-sdk/rpc"

func main() {
	client, conn := sdk.Connect("root", "password", "poc02.ultipa.com:60062")

	defer conn.Close()

	msg := sdk.Statistic(client)

	fmt.Printf("%+v\n", msg)
	// fmt.Print("\n=================AB Search=====================\n\n")

	// abReq := sdk.NewABRequest("123", "321")
	// abReq.Depth = 7
	// abReq.Limit = 1
	// abMsg := sdk.SearchAB(client, abReq)
	// for i, path := range abMsg.Paths {
	// 	fmt.Printf("Path[%v] : ", i)
	// 	for i := 0; i < len(path.Nodes); i++ {
	// 		fmt.Printf("%v", path.Nodes[i]["name"])

	// 		if i < len(path.Nodes)-1 {
	// 			fmt.Printf(" - [%v] - ", path.Edges[i]["name"])
	// 		}
	// 	}
	// 	fmt.Print("\n\n--------------------------------------\n\n")
	// }

	// fmt.Printf("engine cost %v ms, total cost %v ms \n", abMsg.EngineCost, abMsg.TotalCost)

	// fmt.Print("\n======================================\n")

	// fmt.Print("\n=================Search Khop=====================\n\n")

	// khopReq := sdk.NewKhopRequest("123")
	// khopMsg := sdk.SearchKhop(client, khopReq)

	// for _, n := range khopMsg.Nodes {
	// 	fmt.Printf(" [%v] ", n["name"])
	// }

	// fmt.Print("\n\n--------------------------------------\n\n")
	// fmt.Printf("engine cost %v ms, total cost %v ms , total num : %v \n", khopMsg.EngineCost, khopMsg.TotalCost, khopMsg.Count)

	// fmt.Print("\n=================Search Nodes=====================\n\n")

	// nodeReq := sdk.NewSearchNodesRequest()
	// // nodeReq.ID = "123"

	// filterConds := pkg.NewFilterCondition("age", ">", []string{"20"})

	// // fmt.Printf("%v \n", filterConds)

	// filter := pkg.NewFilter("AND", filterConds)

	// nodeReq.NodeFilter = filter

	// nodeReq.Select = []string{"name", "age"}

	// nodeReq.Limit = 100

	// nodeMsg := sdk.SearchNodes(client, nodeReq)

	// for _, v := range nodeMsg.Nodes {
	// 	fmt.Printf(" [%v | %v]", v["name"], v["age"])
	// }

	// fmt.Print("\n\n--------------------------------------\n\n")
	// fmt.Printf("total cost %v ms , total num : %v \n", nodeMsg.TotalCost, nodeMsg.Count)

	// fmt.Print("\n=================Search Edges=====================\n\n")

	// edgeReq := sdk.NewSearchEdgesRequest()
	// // edgeReq.ID = "123"

	// filterConds := pkg.NewFilterCondition("name", "=", []string{"Like"})

	// // fmt.Printf("%v \n", filterConds)

	// filter := pkg.NewFilter("AND", filterConds)

	// edgeReq.EdgeFilter = filter

	// edgeReq.Select = []string{"name", "age"}

	// edgeReq.Limit = 100

	// edgeMsg := sdk.SearchEdges(client, edgeReq)

	// for _, v := range edgeMsg.Edges {
	// 	fmt.Printf(" [ id: %v | type: %v] ", v["_id"], v["name"])
	// }

	// fmt.Print("\n\n--------------------------------------\n\n")
	// fmt.Printf("total cost %v ms , total num : %v \n", edgeMsg.TotalCost, edgeMsg.Count)

}
