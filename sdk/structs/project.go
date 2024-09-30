package structs

//type Projection struct {
//	ProjectName string `json:"project_name"`
//	ProjectType string `json:"project_type"`
//	FilterType  string `json:"filter_type"`
//	GraphName   string `json:"graph_name"`
//	Status      string `json:"status"`
//	Stats       string `json:"stats"`
//	Config      string `json:"config"`
//}

type Projection struct {
	ProjectName string `json:"name"`
	ProjectType string `json:"type"`
	FilterType  string `json:"filterType"`
	IsDefault   string `json:"isDefault"`
	SourceGraph string `json:"sourceGraph"`
	Status      string `json:"status"`
	Stats       string `json:"stats"`
	HDCName     string `json:"HDCName"`
	HDCStatus   string `json:"HDCStatus"`
	Config      string `json:"config"`
}
