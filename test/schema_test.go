package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
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
				Type: ultipa.UltipaPropertyType_INT32,
			},
			{
				Name: "prop2",
				Type: ultipa.UltipaPropertyType_UINT64,
			},
		},
	}

	s2 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.UltipaPropertyType_INT32,
			},
			{
				Name: "prop2",
				Type: ultipa.UltipaPropertyType_UINT64,
			},
		},
	}

	s3 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.UltipaPropertyType_INT32,
			},
		},
	}

	s4 := &structs.Schema{
		Name: "s1",
		Properties: []*structs.Property{
			{
				Name: "prop1",
				Type: ultipa.UltipaPropertyType_UINT64,
			},
			{
				Name: "prop2",
				Type: ultipa.UltipaPropertyType_UINT64,
			},
		},
	}

	log.Println(
		structs.CompareSchemas(s1, s2, false), // true
		structs.CompareSchemas(s1, s3, false), // false
		structs.CompareSchemas(s1, s3, true), // true
		structs.CompareSchemas(s1, s4, false), // false
		structs.CompareSchemas(s1, s4, true), // false
		structs.CompareSchemas(s1, nil, false), // false
		structs.CompareSchemas(s1, nil, true), // true
		structs.CompareSchemas(nil, s1, false), // false
		structs.CompareSchemas(nil, s1, true)) // false
}
