package structs

type Projection struct {
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	FilterType  string `json:"filter_type"`
	GraphName   string `json:"graph_name"`
	Status      string `json:"status"`
	Stats       string `json:"stats"`
	Config      string `json:"config"`
}
