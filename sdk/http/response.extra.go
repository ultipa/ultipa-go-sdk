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
type ResponseProperty struct {


}
type ResponseSchema struct {
	Name string
	Description string
	Properties []*ResponseProperty
	TotalNodes int64
	TotalEdges int64
}
type ResponseNodeSchemas struct {
	Status *Status
	Schemas []*ResponseSchema
}