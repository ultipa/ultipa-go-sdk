package types

type Response_Property struct {
	Lte string
	PropertyName string
	PropertyType string
}

type Response_Stat struct {
	CpuUsage string
	MemUsage string
	ExpiredDate string
}
type Response_ClusterInfo struct {
	*RaftPeerInfo
	*Response_Stat
}

type ResListProperty = struct {
	*ResWithoutData
	Data []*Response_Property
}
type ResListClusterInfo = struct {
	*ResWithoutData
	Data []*Response_ClusterInfo
}
type ResStat = struct {
	*ResWithoutData
	Data *Response_Stat
}
type ResListSample = struct {
	*ResWithoutData
	Data *[]*map[string]interface{}
}