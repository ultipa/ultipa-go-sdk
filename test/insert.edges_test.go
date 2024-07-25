package test

import (
	"log"
	"testing"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func TestInsertEdge(t *testing.T) {
	var edges []*structs.Edge
	edge1 := structs.NewEdge()
	edge1.UUID = 1
	edge1.FromUUID = 1
	edge1.ToUUID = 2
	edge1.Set("value", 1.11)

	edge2 := structs.NewEdge()
	edge2.From = "ULTIPA8000000000000002"
	edge2.To = "ULTIPA8000000000000003"
	edge2.Set("value", "1212")

	edge3 := structs.NewEdge()
	edge3.FromUUID = 1
	edge3.ToUUID = 2

	edges = append(edges, edge1, edge2, edge3)

	uql := structs.EdgesToInsertUql(edges)
	log.Println(uql)

	requestConfig := &configuration.InsertRequestConfig{
		RequestConfig: &configuration.RequestConfig{
			GraphName: "test",
		},
		Silent: true,
	}

	response, err := client.InsertEdges("default", edges, requestConfig)
	if err != nil {
		return
	}

	log.Println(response.Status.Message)
}
