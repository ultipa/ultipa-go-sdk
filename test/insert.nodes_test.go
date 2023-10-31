package test

import (
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
	"log"
	"testing"
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

func TestInsertBlobProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	schemaName := "node_schema"
	schema := structs.NewSchema(schemaName)
	schema.Properties = append(schema.Properties, &structs.Property{
		Name:     "name",
		Type:     ultipa.PropertyType_TEXT,
		SubTypes: nil,
	}, &structs.Property{
		Name:     "blob_prop",
		Type:     ultipa.PropertyType_BLOB,
		SubTypes: nil,
	})

	var nodes []*structs.Node
	node1 := structs.NewNode()
	node1.Set("name", "go_sdk")
	node1.Set("blob_prop", []byte{97, 98, 99})

	node2 := structs.NewNode()
	node2.Set("name", "test")
	node2.Set("blob_prop", "def")

	nodes = append(nodes, node1, node2)

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

	uql := fmt.Sprintf("find().nodes({@%s}) as nodes return nodes{*}", schemaName)
	response, err := client.UQL(uql, nil)

	//断言响应码
	if response.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(response.Status.Message)
		t.Fatal(response.Status.Message)
	}
	nodes, schemas, err := response.Alias("nodes").AsNodes()
	printers.PrintNodes(nodes, schemas)
}

func TestInsertDecimalProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	schemaName := "default"
	schema := structs.NewSchema(schemaName)
	schema.Properties = append(schema.Properties, &structs.Property{
		Name:     "name",
		Type:     ultipa.PropertyType_STRING,
		SubTypes: nil,
	}, &structs.Property{
		Name:     "salary",
		Type:     ultipa.PropertyType_DECIMAL,
		SubTypes: nil,
	})

	var nodes []*structs.Node
	node1 := structs.NewNode()
	node1.Set("name", "go_sdk")
	node1.Set("salary", "6.1")

	node2 := structs.NewNode()
	node2.Set("name", "test")
	node2.Set("salary", 6.1)

	nodes = append(nodes, node1, node2)

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

	uql := fmt.Sprintf("find().nodes({@%s}) as nodes return nodes{*}", schemaName)
	response, err := client.UQL(uql, nil)

	//断言响应码
	if response.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(response.Status.Message)
		t.Fatal(response.Status.Message)
	}
	nodes, schemas, err := response.Alias("nodes").AsNodes()
	printers.PrintNodes(nodes, schemas)
}

func TestInsertNodeWithSetProperty(t *testing.T) {
	client, _ := GetClient(hosts, graph)

	schema := structs.NewSchema("default")
	schema.Properties = append(schema.Properties, &structs.Property{
		Name:     "string_set",
		Type:     ultipa.PropertyType_SET,
		SubTypes: []ultipa.PropertyType{ultipa.PropertyType_STRING},
	})

	var nodes []*structs.Node
	node := structs.NewNode()

	node.Set("string_set", []string{"list", "set", "map"})

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
