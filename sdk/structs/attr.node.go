package structs

import ultipa "ultipa-go-sdk/rpc"

//AttrNodes represents an Attr with Rows that is List<List<Node>>
type AttrNodes struct {
	Name       string
	ResultType ultipa.ResultType
	NodesList  [][]*Node
}

func NewAttrNodes() *AttrNodes {
	return &AttrNodes{}
}
