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
	client, _ := GetClient([]string{"192.168.1.87:50051"}, "default")

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
	client, _ := GetClient([]string{"192.168.1.85:61090"}, "listPropertyGraphTest")
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

func TestInsertBoolProperty(t *testing.T) {
	client, _ := GetClient([]string{"192.168.1.85:61099"}, "sdk_test")
	schema := structs.NewSchema("People")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name: "bool_prop",
		Type: ultipa.PropertyType_BOOL,
	})

	var nodes []*structs.Node
	node1 := structs.NewNode()
	node1.ID = "1"
	node1.Set("bool_prop", true)

	node2 := structs.NewNode()
	node2.ID = "2"
	node2.Set("bool_prop", "true")

	node3 := structs.NewNode()
	node3.ID = "3"
	node3.Set("bool_prop", "false")

	node4 := structs.NewNode()
	node4.ID = "4"
	node4.Set("bool_prop", false)

	nodes = append(nodes, node1, node2, node3, node4)

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
