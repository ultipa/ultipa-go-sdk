package structs

type Graph struct {
	Name string
	Nodes []*Node
	Edges []*Edge
	NodeSchemas map[string]*Schema
	EdgeSchemas map[string]*Schema
}


func NewGraph() *Graph {
	return &Graph{
		NodeSchemas: map[string]*Schema{},
		EdgeSchemas: map[string]*Schema{},
	}
}

func (p *Graph) GetNodes() []*Node {
	return p.Nodes
}

func (p *Graph) GetLastNode() *Node {
	return p.Nodes[p.GetLength()]
}

func (p *Graph) GetEdges() []*Edge {
	return p.Edges
}

func (p *Graph) GetLength() int {
	return len(p.Edges)
}