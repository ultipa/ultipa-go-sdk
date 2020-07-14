package types_response

type Property struct {
	Lte string
	PropertyName string
	PropertyType string
}

type Stat struct {
	CpuUsage string
	MemUsage string
}
type ClusterInfo struct {
	*Stat
	Host string
	Status bool
}