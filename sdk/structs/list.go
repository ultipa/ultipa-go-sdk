package structs

import ultipa "ultipa-go-sdk/rpc"

type List struct {
	BaseType ultipa.PropertyType
	Values   []*ListValue
}

type ListValue struct {
	Type  ultipa.PropertyType
	Value interface{}
}

type ListData struct {
	Values []interface{}
}
