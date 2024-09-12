package structs

import "github.com/ultipa/ultipa-go-sdk/sdk/types"

type Path struct {
	//Name        string
	NodeUUIDs []types.UUID
	EdgeUUIDs []types.UUID
	//NodeSchemas map[string]*Schema
	//EdgeSchemas map[string]*Schema
}

//func NewPath() *Path {
//    return &Path{
//        NodeSchemas: map[string]*Schema{},
//        EdgeSchemas: map[string]*Schema{},
//    }
//}

func (p *Path) GetNodes() []types.UUID {
	return p.NodeUUIDs
}

func (p *Path) GetLastNode() types.UUID {
	return p.NodeUUIDs[p.GetLength()]
}

func (p *Path) GetEdges() []types.UUID {
	return p.EdgeUUIDs
}

func (p *Path) GetLength() int {
	return len(p.EdgeUUIDs)
}
