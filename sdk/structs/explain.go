package structs

type PlanNode struct {
	//DBType        ultipa.PlanNodeType
	Alias       string
	ChildrenNum uint32
	Uql         string
	Infos       string
}
