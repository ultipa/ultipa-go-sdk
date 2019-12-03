package pkg

import (
	"ultipa-go-sdk/rpc"
)

type Path struct {
	Nodes []map[string]string
	Edges []map[string]string
}

type Paths []Path

func FormatPaths(paths []*ultipa.ABPath) Paths {
	var ps Paths

	for _, v := range paths {

		var newPath Path

		for _, nv := range v.Nodes {
			newNode := make(map[string]string)
			newPath.Nodes = append(newPath.Nodes, newNode)
			for _, nvv := range nv.Values {
				newNode[nvv.Key] = nvv.Value
			}
		}

		for _, ev := range v.Edges {
			newEdge := make(map[string]string)
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
