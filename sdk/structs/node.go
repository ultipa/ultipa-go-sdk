package structs

import (
	"ultipa-go-sdk/sdk/types"
)

type Node struct {
	ID types.ID
	UUID types.UUID
	Schema types.Schema
	Values types.Values
}

func NewNodeFromMetaData(md *MetaData) *Node {
	return &Node{
		ID: md.ID,
		UUID: md.UUID,
		Schema: md.Schema,
		Values: md.Values,
	}
}

func (node *Node) GetID() types.ID {
	return node.ID
}

func (node *Node) GetUUID() types.UUID {
	return node.UUID
}

func (node *Node) GetSchema() types.Schema {
	return node.Schema
}

// get a value by key
func (node *Node) Get(key string) interface{} {
	return (*node.Values)[key]
}

// set a value by key
func (node *Node) Set(key string, value interface{}) error {

	//todo: check value type
	(*node.Values)[key] = value
	return nil
}