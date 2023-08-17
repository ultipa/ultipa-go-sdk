package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/types"
)

func TestInsertNodeWithListProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	schema := structs.NewSchema("default")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name:     "name",
		Type:     ultipa.PropertyType_LIST,
		SubTypes: []ultipa.PropertyType{ultipa.PropertyType_STRING},
	})

	var nodes []*structs.Node
	node := structs.NewNode()

	node.Set("name", []string{"lzq", "list", "set", "map"})

	nodes = append(nodes, node)

	resp, err := client.InsertNodesBatchBySchema(schema, nodes, &configuration.InsertRequestConfig{
		InsertType: ultipa.InsertType_OVERWRITE,
	})

	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
}

func TestInsertPointProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	schema := structs.NewSchema("nodeSchemaList")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name:     "typePoint",
		Type:     ultipa.PropertyType_POINT,
		SubTypes: nil,
	})

	var nodes []*structs.Node
	nodeWithPointer := structs.NewNode()
	nodeWithPointer.Set("typePoint", types.NewPoint(1.01, -2.01))

	nodeWithValue := structs.NewNode()
	nodeWithValue.Set("typePoint", types.Point{
		Latitude:  100.90,
		Longitude: 80.11,
	})

	nodeWithString := structs.NewNode()
	nodeWithString.Set("typePoint", "point(1.03 26.05)")

	nodes = append(nodes, nodeWithPointer, nodeWithValue, nodeWithString)

	resp, err := client.InsertNodesBatchBySchema(schema, nodes, &configuration.InsertRequestConfig{
		InsertType: ultipa.InsertType_OVERWRITE,
	})

	if err != nil {
		log.Fatalln(err)
	}
	//断言响应码
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

}
