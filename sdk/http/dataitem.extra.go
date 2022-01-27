package http

import (
	"ultipa-go-sdk/sdk/structs"
)

func (di *DataItem) AsFirstNode() (node *structs.Node, err error) {
	nodes, _, err := di.AsNodes()

	if len(nodes) < 1 {
		return nil, err
	}

	return nodes[0], err
}

func (di *DataItem) AsFirstEdge() (node *structs.Edge, err error) {
	edges, _, err := di.AsEdges()

	if len(edges) < 1 {
		return nil ,err
	}

	return edges[0], err
}