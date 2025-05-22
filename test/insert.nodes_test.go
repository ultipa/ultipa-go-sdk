package test

import (
	"fmt"
	"log"
	"testing"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
)

func TestInsertNodeWithListProperty(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

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
		t.Fatal(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
}

func TestInsertPointProperty(t *testing.T) {
	//client, _ := GetClient(hosts, graph)
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
		t.Fatal(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

}

func TestInsertBlobProperty(t *testing.T) {
	//client, _ := GetClient(hosts, graph)
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
		t.Fatal(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

	uql := fmt.Sprintf("find().nodes({@%s}) as nodes return nodes{*}", schemaName)
	response, err := client.Uql(uql, nil)

	if err != nil {
		t.Fatal(err)
	}

	nodes, schemas, err := response.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}
	printers.PrintNodes(nodes, schemas)
}

func TestInsertDecimalProperty(t *testing.T) {
	//client, _ := GetClient(hosts, graph)
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
		t.Fatal(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

	uql := fmt.Sprintf("find().nodes({@%s}) as nodes return nodes{*}", schemaName)
	response, err := client.Uql(uql, nil)
	if err != nil {
		t.Fatal(err)
	}

	nodes, schemas, err := response.Alias("nodes").AsNodes()
	if err != nil {
		t.Fatal(err)
	}

	printers.PrintNodes(nodes, schemas)
}

func TestInsertNodeWithSetProperty(t *testing.T) {
	//client, _ := GetClient(hosts, graph)

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
		t.Fatal(err)
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)
}

func TestInsertNodes(t *testing.T) {
	//client.SetCurrentGraph("go_sdk_test")
	schemaName := "default"

	ty := ultipa.DBType_DBNODE
	client.Truncate(&structs.TruncateParams{
		GraphName: graph,
		DBType:    &ty,
		Schema:    schemaName,
	}, nil)

	prop := &structs.Property{
		Schema: schemaName,
		Name:   "name",
		Type:   ultipa.PropertyType_STRING,
	}
	client.CreateNodeProperty(prop, nil)
	prop = &structs.Property{
		Schema: schemaName,
		Name:   "salary",
		Type:   ultipa.PropertyType_DOUBLE,
	}
	client.CreateNodeProperty(prop, nil)

	var nodes []*structs.Node
	node1 := structs.NewNode()
	node1.ID = "11131"
	node1.Set("name", "go_sdk")
	node1.Set("salary", "6.1")

	node2 := structs.NewNode()
	node2.ID = "223222"
	node2.Set("name", "test")
	node2.Set("salary", 6.1)

	node3 := structs.NewNode()
	node3.ID = "333433"
	//node3.Set("name", "test2")

	nodes = append(nodes, node1, node2, node3)

	uql := structs.NodesToInsertUql(nodes)
	t.Log(uql)

	config := &configuration.InsertRequestConfig{
		RequestConfig: &configuration.RequestConfig{},
		Silent:        true,
	}

	response, err := client.InsertNodes(schemaName, nodes, config)
	if err != nil {
		t.Error(err)
	}

	t.Log(response)

	// overwrite
	node1.Set("name", "go_sdk2")

	uql = structs.NodesToInsertUql(nodes)
	t.Log(uql)

	config = &configuration.InsertRequestConfig{
		InsertType: ultipa.InsertType_OVERWRITE,
		Silent:     true,
	}

	response, err = client.InsertNodes(schemaName, nodes, config)
	if err != nil {
		t.Error(err)
	}

	t.Log(response)

	// upsert
	node2.Set("salary", 10.1)

	uql = structs.NodesToInsertUql(nodes)
	log.Println(uql)

	config = &configuration.InsertRequestConfig{
		InsertType: ultipa.InsertType_UPSERT,
		Silent:     true,
	}

	response, err = client.InsertNodes(schemaName, nodes, config)
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
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

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		log.Println(resp.Status.Message)
		t.Log(resp.Status.Message)
	}
	log.Println(resp.Statistic.EngineCost, "|", resp.Statistic.TotalCost)

}
