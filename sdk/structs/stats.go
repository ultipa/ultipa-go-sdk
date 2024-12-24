package structs

type Stats struct {
	CPUUsage    string `json:"cpuUsage"`
	MemUsage    string `json:"memUsage"`
	ExpiredDate string `json:"expiredDate"`
	CPUCores    string `json:"cpuCores"`
	Company     string `json:"company"`
	ServerType  string `json:"serverType"`
	Version     string `json:"version"`
}
