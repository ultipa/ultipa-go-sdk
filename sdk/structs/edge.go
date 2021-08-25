package structs

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/types"
	"ultipa-go-sdk/sdk/utils"
)

type Edge struct {
	Name     string
	From     types.ID
	To       types.ID
	FromUUID types.UUID
	ToUUID   types.UUID
	UUID     types.UUID
	Schema   string
	Values   *Values
}

func NewEdge() *Edge {
	return &Edge{
		Values: NewValues(),
	}
}

func NewEdgeFromMetaData(md *MetaData) *Edge {
	return &Edge{
		From:   md.From,
		To:     md.To,
		UUID:   md.UUID,
		Schema: md.Schema,
		Values: md.Values,
	}
}

func NewEdgeFromEdgeRow(schema *Schema, edgeRow *ultipa.EdgeRow) *Edge {
	newEdge := NewEdge()

	newEdge.UUID = edgeRow.Uuid
	newEdge.From = edgeRow.FromId
	newEdge.To = edgeRow.ToId
	newEdge.FromUUID = edgeRow.FromUuid
	newEdge.ToUUID = edgeRow.ToUuid
	newEdge.Name = edgeRow.SchemaName

	for index, v := range edgeRow.GetValues() {
		prop := schema.Properties[index]
		newEdge.Values.Set(prop.Name, utils.ConvertBytesToInterface(v, prop.Type))
	}

	return newEdge
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

func (edge *Edge) GetSchema() string {
	return edge.Schema
}

func (edge *Edge) GetValues() *Values {
	return edge.Values
}

// get a value by key
func (edge *Edge) Get(key string) interface{} {
	return edge.Values.Get(key)
}

// get a value by key
func (edge *Edge) GetBytes(key string) ([]byte, error) {
	v := edge.Values.Get(key)
	return utils.ConvertInterfaceToBytes(v)
}

// set a value by key
func (edge *Edge) Set(key string, value interface{}) error {

	edge.Values.Set(key, value)
	return nil
}
