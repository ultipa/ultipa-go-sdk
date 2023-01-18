package structs

import ultipa "ultipa-go-sdk/rpc"

type Attr struct {
	Name         string
	PropertyType ultipa.PropertyType
	Rows         Row
}

func NewAttr() *Attr {
	return &Attr{
		Rows: Row{},
	}
}

type AttrListData struct {
	ResultType ultipa.ResultType
	Nodes      []*Node
	Edges      []*Edge
	Paths      []*Path
	Attrs      []*Attr
}

func NewAttrListData() *AttrListData {
	return &AttrListData{}
}

type AttrMapData struct {
	Key   *Attr
	Value *Attr
}

func NewAttrMapData() *AttrMapData {
	return &AttrMapData{}
}
