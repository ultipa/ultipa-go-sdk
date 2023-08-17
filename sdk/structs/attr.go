package structs

import ultipa "github.com/ultipa/ultipa-go-sdk/rpc"

type Attr struct {
	Name         string
	PropertyType ultipa.PropertyType
	ResultType   ultipa.ResultType
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
	return &AttrMapData{
		Key:   NewAttr(),
		Value: NewAttr(),
	}
}
