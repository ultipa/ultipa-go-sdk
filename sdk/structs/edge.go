package structs

import (
"ultipa-go-sdk/sdk/types"
)

type Edge struct {
	ID types.ID
	From types.ID
	To types.ID
	UUID types.UUID
	Schema types.Schema
	Values types.Values
}
func NewEdgeFromMetaData(md *MetaData) *Edge {
	return &Edge{
		ID: md.ID,
		From: md.From,
		To: md.To,
		UUID: md.UUID,
		Schema: md.Schema,
		Values: md.Values,
	}
}

func (edge *Edge) GetID() types.ID {
	return edge.ID
}

func (edge *Edge) GetUUID() types.UUID {
	return edge.UUID
}

func (edge *Edge) GetFrom() types.ID {
	return edge.From
}

func (edge *Edge) GetTo() types.ID {
	return edge.To
}

func (edge *Edge) GetSchema() types.Schema {
	return edge.Schema
}

// get a value by key
func (edge *Edge) Get(key string) interface{} {
	return (*edge.Values)[key]
}

// set a value by key
func (edge *Edge) Set(key string, value interface{}) error {

	//todo: check value type
	(*edge.Values)[key] = value
	return nil
}