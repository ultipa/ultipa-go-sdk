package structs

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
