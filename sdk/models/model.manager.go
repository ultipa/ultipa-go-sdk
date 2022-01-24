package models

import "ultipa-go-sdk/sdk/structs"

type GraphModel struct {
	GraphName string
	Schemas   []*structs.Schema
}

//TODO: finish Graph model
func NewGraphModel(graphName string) *GraphModel {

	gm := &GraphModel{
		GraphName: graphName,
	}

	return gm
}

func (gm *GraphModel) AddSchema(schema *structs.Schema) error {
	gm.Schemas = append(gm.Schemas, schema)
	return nil
}

func (gm *GraphModel) DeleteSchema() error {
	return nil
}
