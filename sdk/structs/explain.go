package structs

import ultipa "ultipa-go-sdk/rpc"

type Explain struct {
	Type        ultipa.PlanNodeType
	Alias       string
	ChildrenNum uint32
	Uql         string
	Infos       string
}
