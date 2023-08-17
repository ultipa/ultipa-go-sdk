package test

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils/logger"
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"testing"
)

func TestListGraph(t *testing.T) {
	InitCases()
	client, _ := GetClient(hosts, graph)
	res, err := client.ListGraph(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}

func TestCreateGraph(t *testing.T) {

	client, err := GetClient(hosts, graph)

	if err != nil {
		log.Println(err)
		return
	}

	client.DropGraph(graph, nil)

	client.CreateGraph(&structs.Graph{
		Name: graph,
	}, nil)

	client.SetCurrentGraph(graph)

	resp, err := client.UQL("insert().nodes({}).into(@default)", nil)

	if err != nil {
		logger.PrintError(err.Error())
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		logger.PrintError(resp.Status.Message)
	}

}

func TestDeleteGraph(t *testing.T) {
	client.DropGraph("test_creation", nil)
}

func TestAsGraph(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	resp, _ := client.UQL("show().graph()", nil)
	graphs, err := resp.Alias(http.RESP_GRAPH_KEY).AsGraphs()

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintGraph(graphs)
}

func TestCreateGraphIfNotExist(t *testing.T) {

	client, err := GetClient(hosts, graph)

	if err != nil {
		t.Fatalf("failed to connect to server %v", err)
	}

	client.DropGraph(graph, nil)

	_, _, err = client.CreateGraphIfNotExit(&structs.Graph{
		Name: graph,
	}, nil)
	if err != nil {
		t.Fatalf("failed to create graph %v", err)
	}
}
