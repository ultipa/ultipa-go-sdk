package test

import (
	"log"
	"testing"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/models"
	"ultipa-go-sdk/sdk/structs"
)

func TestCreateModel(t *testing.T) {

	var err error
	client, err = GetClient(hosts, graph)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[Test] Creating Graph Model")
	model := models.NewGraphModel(&structs.Graph{
		Name: "graph_by_model",
	})

	// create user schema
	model.AddSchema(&structs.Schema{
		Name:   "User",
		DBType: ultipa.DBType_DBNODE,
		Properties: []*structs.Property{
			{
				Name: "username",
				Desc: "user's name",
				Type: ultipa.PropertyType_STRING,
			},
			{
				Name: "password",
				Desc: "user's password",
				Type: ultipa.PropertyType_STRING,
			},
		},
	})

	// create indicator schema
	model.AddSchema(&structs.Schema{
		Name:   "Indicator",
		DBType: ultipa.DBType_DBNODE,
		Properties: []*structs.Property{
			{
				Name: "name",
				Desc: "indicator name",
				Type: ultipa.PropertyType_STRING,
			},
			{
				Name: "values",
				Desc: "json content to save indicator values",
				Type: ultipa.PropertyType_STRING,
			},
		},
	})

	model.AddSchema(&structs.Schema{
		Name:   "Privilege",
		DBType: ultipa.DBType_DBEDGE,
		Properties: []*structs.Property{
			{
				Name: "type",
				Desc: "type of privilege: r,w, rw",
				Type: ultipa.PropertyType_STRING,
			},
		},
	})

	log.Println("[TEST] Initial Model")
	err = client.InitModel(model, nil)

	if err != nil {
		t.Fatalf("Test Error %v \n", err)
	}
}
