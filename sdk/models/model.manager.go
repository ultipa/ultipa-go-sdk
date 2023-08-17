package models

import "github.com/ultipa/ultipa-go-sdk/sdk/structs"

type GraphModel struct {
	Graph   *structs.Graph
	Schemas []*structs.Schema
}

func NewGraphModel(graph *structs.Graph) *GraphModel {

	gm := &GraphModel{
		Graph: graph,
	}

	return gm
}

func (gm *GraphModel) AddSchema(schema *structs.Schema) {
	gm.Schemas = append(gm.Schemas, schema)
}

//TODO:
func (gm *GraphModel) NewGraphModelFromYAML(path string) {
}
