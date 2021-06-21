package structs

import (
	"ultipa-go-sdk/sdk/types"
)

type Node struct {
	Name string
	ID types.ID
	UUID types.UUID
	Schema string
	Values *Values
}


func (node *Node) GetID() types.ID {
	return node.ID
}

func (node *Node) GetUUID() types.UUID {
	return node.UUID
}

func (node *Node) GetSchema() string {
	return node.Schema
}

func (node *Node) GetValues() *Values {
	return node.Values
}

// get a value by key
func (node *Node) Get(key string) interface{} {
	return node.Values.Get(key)
}

// set a value by key
func (node *Node) Set(key string, value interface{}) error {

	//todo: check value type
	node.Values.Set(key, value)
	return nil
}