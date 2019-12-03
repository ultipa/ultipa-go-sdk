package pkg

import (
	"ultipa-go-sdk/rpc"
)

type Node map[string]string
type Edge map[string]string

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
