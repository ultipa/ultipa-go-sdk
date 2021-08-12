package structs

import (
"ultipa-go-sdk/sdk/types"
	"ultipa-go-sdk/sdk/utils"
)

type Edge struct {
	Name string
	From types.ID
	To types.ID
	FromUUID types.UUID
	ToUUID types.UUID
	UUID types.UUID
	Schema string
	Values *Values
}
func NewEdgeFromMetaData(md *MetaData) *Edge {
	return &Edge{
		From: md.From,
		To: md.To,
		UUID: md.UUID,
		Schema: md.Schema,
		Values: md.Values,
	}
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