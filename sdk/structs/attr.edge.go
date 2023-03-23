package structs

import ultipa "ultipa-go-sdk/rpc"

//AttrEdges represents an Attr with Rows that is List<List<Edge>>
type AttrEdges struct {
	Name       string
	ResultType ultipa.ResultType
	EdgesList  [][]*Edge
}

func NewAttrEdges() *AttrEdges {
	return &AttrEdges{}
}
