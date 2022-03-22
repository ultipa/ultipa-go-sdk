package api_test

import (
	"log"
	"ultipa-go-sdk/sdk"
	"ultipa-go-sdk/sdk/api"
	"ultipa-go-sdk/sdk/configuration"
)

func ExampleNewUltipaAPI() {

	config := configuration.NewUltipaConfig(&configuration.UltipaConfig{
		Hosts: []string{"10.0.0.1:60061","10.0.0.2:60061","10.0.0.3:60061"},
		Username: "root",
		Password: "root",
	})

	ultipa, err := sdk.NewUltipa(config)

	if err != nil {
		log.Fatalln(err)
	}

	graph, _ := ultipa.ListGraph(nil)

	log.Println(graph.Graphs)
}

func ExampleUltipaAPI_UQL_Nodes_Edges() {

	var ultipa *api.UltipaAPI

	rConfig := &configuration.RequestConfig{
		Timeout: 20,
	}
	resp, err := ultipa.UQL("find().nodes() return nodes limit 1", rConfig)

	nodes, schemas, err := resp.Alias("nodes").AsNodes()

	log.Println(nodes, schemas, err)



	respEdges, err := ultipa.UQL("find().edges() return edges limit 1", nil)

	edges, edgeSchemas, err := respEdges.Alias("edges").AsEdges()

	log.Println(edges, edgeSchemas, err)
}
