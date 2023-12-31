package test

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/printers"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/utils"
	"log"
	"testing"
)

func TestListSchema(t *testing.T) {
	InitCases()
	res, err := client.ListNodeSchema(nil)
	if err != nil {
		log.Panic(err)
	}
	log.Printf(utils.JSONString(res))
}

func TestCompareSchema(t *testing.T) {

	s1 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.PropertyType_INT32,
			},
			{
				Name: "prop2",
				Type: ultipa.PropertyType_UINT64,
			},
		},
	}

	s2 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.PropertyType_INT32,
			},
			{
				Name: "prop2",
				Type: ultipa.PropertyType_UINT64,
			},
		},
	}

	s3 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.PropertyType_INT32,
			},
		},
	}

	s4 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.PropertyType_UINT64,
			},
			{
				Name: "prop2",
				Type: ultipa.PropertyType_UINT64,
			},
		},
	}

	schemaParis := []struct {
		First  *structs.Schema
		Second *structs.Schema
		Fit    bool
		Expect bool
	}{
		{s1, s2, false, true},
		{s1, s3, false, false},
		{s1, s3, true, true},
		{s1, s4, false, false},
		{s1, s4, true, false},
		{s1, nil, false, false},
		{s1, nil, true, true},
		{nil, s1, false, false},
		{nil, s1, true, false},
	}
	for index, pair := range schemaParis {
		err, _ := structs.CompareSchemas(pair.First, pair.Second, false)

		if err != nil {
			log.Fatalln("Test Compare schema failed ：", index, err)
		}
	}

}

func TestShowSchema(t *testing.T) {
	resp, _ := client.UQL("show().schema()", nil)

	nodeSchemas, err := resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		log.Fatalln(err)
	}
	edgeSchemas, err := resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		log.Fatalln(err)
	}

	printers.PrintSchema(nodeSchemas)
	printers.PrintSchema(edgeSchemas)
}

func TestCreateSchemaWithProperties(t *testing.T) {
	client, _ := GetClient(hosts, graph)
	// create schema with properties
	newSchemaWithProperties := &structs.Schema{
		Name: "_abc _acd",
		Desc: "A Schema with 2 properties",
		Properties: []*structs.Property{
			{
				Name: "username用户@",
				Type: ultipa.PropertyType_STRING,
			},
			{
				Name: "123passwor@d",
				Type: ultipa.PropertyType_TEXT,
			},
		},
	}

	resp2, err := client.CreateSchema(newSchemaWithProperties, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(resp2)
}

func TestCreateSchema(t *testing.T) {
	// create schema with properties
	newSchemaWithoutProperties := &structs.Schema{
		Name: "People",
		Desc: "People",
	}

	resp2, _ := client.CreateSchema(newSchemaWithoutProperties, false, nil)
	log.Println(resp2)
}
