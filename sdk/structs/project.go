package structs

//type HDCGraph struct {
//	Name string `json:"project_name"`
//	ProjectType string `json:"project_type"`
//	FilterType  string `json:"filter_type"`
//	Graph   string `json:"graph_name"`
//	Status      string `json:"status"`
//	Stats       string `json:"stats"`
//	Config      string `json:"config"`
//}

// HDCGragh
type HDCGraph struct {
	Name            string `json:"name"`
	GraphName       string `json:"graph_name"`
	Status          string `json:"status"`
	Stats           string `json:"stats"`
	IsDefault       string `json:"is_default"`
	HDCServerName   string `json:"hdc_server_name"`
	HDCServerStatus string `json:"hdc_server_status"`
	Config          string `json:"config"`
}

type Projection struct {
	Name      string `json:"name"`
	GraphName string `json:"graph_name"`
	Status    string `json:"status"`
	Stats     string `json:"stats"`
	Config    string `json:"config"`
}
