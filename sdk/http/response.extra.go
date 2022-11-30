package http

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
)

type ResponseGraphs struct {
	Status *Status
	Graphs []*ResponseGraph
}

func (graphs *ResponseGraphs) Find(name string) *ResponseGraph {
	for _, g := range graphs.Graphs {
		if name == g.Name {
			return g
		}
	}

	return nil
}

type ResponseGraph struct {
	Id         int64
	ClusterId  string
	Name       string
	TotalNodes int64
	TotalEdges int64
	Status     string
	Description string
}

type ResponseProperty struct {
}

type ResponseSchema struct {
	Name        string
	Description string
	Properties  []*ResponseProperty
	TotalNodes  int64
	TotalEdges  int64
}

type ResponseNodeSchemas struct {
	Status  *Status
	Schemas []*ResponseSchema
}
type ResponseIndex struct {
	Type    ultipa.DBType
	Indexes []*structs.Index
}
