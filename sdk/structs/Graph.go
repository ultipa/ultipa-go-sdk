package structs

import "github.com/ultipa/ultipa-go-sdk/sdk/types"

type Graph struct {
	Paths []*Path
	//Name        string
	Nodes map[types.UUID]*Node
	Edges map[types.UUID]*Edge
	//NodeSchemas map[string]*Schema
	//EdgeSchemas map[string]*Schema
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: map[types.UUID]*Node{},
		Edges: map[types.UUID]*Edge{},
	}
}

func (p *Graph) GetNodes(uuid types.UUID) *Node {
	return p.Nodes[uuid]
}

//func (p *Graph) GetLastNode() *Node {
//    return p.Nodes[p.GetLength()]
//}

func (p *Graph) GetEdges(uuid types.UUID) *Edge {
	return p.Edges[uuid]
}

//func (p *Graph) GetLength() int {
//    return len(p.Edges)
//}
