package structs

type Explain struct {
	//DBType        ultipa.PlanNodeType
	Alias       string
	ChildrenNum uint32
	Uql         string
	Infos       string
}
