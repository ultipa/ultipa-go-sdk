package structs

import ultipa "ultipa-go-sdk/rpc"

type Attr struct {
	Name string
	PropertyType ultipa.PropertyType
	Rows Row
}

func NewAttr() *Attr {
	return &Attr{
		Rows: Row{},
	}
}
