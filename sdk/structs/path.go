package structs

type Path struct {
	Name string
	Nodes []*Node
	Edges []*Edge
	NodeSchemas map[string]*Schema
	EdgeSchemas map[string]*Schema
}


func NewPath() *Path {
	return &Path{
		NodeSchemas: map[string]*Schema{},
		EdgeSchemas: map[string]*Schema{},
	}
}

func (p *Path) GetNodes() []*Node {
	return p.Nodes
}

func (p *Path) GetLastNode() *Node {
	return p.Nodes[p.GetLength()]
}

func (p *Path) GetEdges() []*Edge {
	return p.Edges
}

func (p *Path) GetLength() int {
	return len(p.Edges)
}