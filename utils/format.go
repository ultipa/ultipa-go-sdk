package utils

import (
	"ultipa-go-sdk/rpc"
)

type Node map[string]string
type Edge map[string]string

type Value struct {
	Key   string
	Value string
}

type RpcNode struct {
	Id     string
	Values []Value
}

type RpcEdge struct {
	Id     string
	FromId string
	ToId   string
	Values []Value
}

type Path struct {
	Nodes []Node
	Edges []Edge
}

type Paths []Path

func FormatPaths(paths []*ultipa.ABPath) Paths {
	var ps Paths

	for _, v := range paths {

		var newPath Path

		for _, nv := range v.Nodes {
			newNode := make(Node)
			newPath.Nodes = append(newPath.Nodes, newNode)
			for _, nvv := range nv.Values {
				newNode[nvv.Key] = nvv.Value
			}
		}

		for _, ev := range v.Edges {
			newEdge := make(Edge)
			newPath.Edges = append(newPath.Edges, newEdge)
			for _, evv := range ev.Values {
				newEdge[evv.Key] = evv.Value
			}
		}

		ps = append(ps, newPath)
		// fmt.Printf("%v,%v\n", i, newPath)
	}

	return ps
}

// FormatNodes return beautiful nodes array instead of rpc
func FormatNodes(nodes []*ultipa.SearchNode) []Node {
	newNodes := []Node{}

	for _, n := range nodes {

		newNode := Node{}
		newNode["_id"] = n.XId

		for _, v := range n.Values {
			newNode[v.Key] = v.Value
		}

		newNodes = append(newNodes, newNode)
	}

	return newNodes
}

func FormatEdges(edges []*ultipa.SearchEdge) []Edge {
	newEdges := []Edge{}

	for _, n := range edges {

		newEdge := Edge{}
		newEdge["_id"] = n.XId
		newEdge["_from_id"] = n.XFromId
		newEdge["_to_id"] = n.XToId

		for _, v := range n.Values {
			newEdge[v.Key] = v.Value
		}

		newEdges = append(newEdges, newEdge)
	}

	return newEdges
}

func ToRpcNodes(nodes []Node) []RpcNode {

	var newNodes []RpcNode

	for _, n := range nodes {
		var newNode RpcNode
		for k, v := range n {
			if k == "_id" {
				newNode.Id = v
			} else {
				newNode.Values = append(newNode.Values, Value{Key: k, Value: v})
			}
		}
		newNodes = append(newNodes, newNode)
	}

	return newNodes
}

func ToRpcEdges(edges []Edge) []RpcEdge {

	var newEdges []RpcEdge

	for _, n := range edges {
		var newEdge RpcEdge
		for k, v := range n {
			if k == "_id" {
				newEdge.Id = v
			} else if k == "_from_id" {
				newEdge.FromId = v
			} else if k == "_to_id" {
				newEdge.ToId = v
			} else {
				newEdge.Values = append(newEdge.Values, Value{Key: k, Value: v})
			}
		}
		newEdges = append(newEdges, newEdge)
	}

	return newEdges
}
