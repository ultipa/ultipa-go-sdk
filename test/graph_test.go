package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/utils"
)

func TestListGraph(t *testing.T) {
	InitCases()
	client, _ := GetClient([]string{"192.168.1.85:61095"}, "miniCircle")
	res, err := client.ListGraph(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}

func TestCreateGraph(t *testing.T) {

	graphName := "test_creation"
	hosts := []string{
		"192.168.1.85:61095",
	}
	client, err := GetClient(hosts, "default")

	if err != nil {
		log.Println(err)
		return
	}

	client.DropGraph(graphName, nil)

	client.CreateGraph(&structs.Graph{
		Name: graphName,
	}, nil)

	client.SetCurrentGraph(graphName)

	resp, err := client.UQL("insert().nodes({}).into(@default)", nil)

	if err != nil {
		printers.PrintError(err.Error())
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		printers.PrintError(resp.Status.Message)
	}

}

func TestDeleteGraph(t *testing.T) {
	client.DropGraph("test_creation", nil)
}

func TestAsGraph(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:60701"}, "miniCircle")
	resp, _ := client.UQL("show().graph()", nil)
	graphs, err := resp.Alias(http.RESP_GRAPH_KEY).AsGraphs()

	if err != nil {
		log.Fatalln(err)
	}
	printers.PrintGraph(graphs)
}
