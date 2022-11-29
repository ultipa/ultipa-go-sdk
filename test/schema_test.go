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

func TestCreateSchema(t *testing.T) {
	// create schema with properties
	newSchemaWithProperties := &structs.Schema{
		Name: "text_schema",
		Desc: "A Schema with 2 properties",
		Properties: []*structs.Property{
			{
				Name: "username",
				Type: ultipa.PropertyType_STRING,
			},
			{
				Name: "password",
				Type: ultipa.PropertyType_TEXT,
			},
		},
	}

	resp2, _ := client.CreateSchema(newSchemaWithProperties, true, nil)
	log.Println(resp2)
}