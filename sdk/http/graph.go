package http

type ResponseGraphs struct {
	Status *Status
	Graphs []*ResponseGraph
}
type ResponseGraph struct {
	Id         int64
	ClusterId  string
	Name       string
	TotalNodes int64
	TotalEdges int64
}