package types_response

import "ultipa-go-sdk/types"

type Property struct {
	Lte string
	PropertyName string
	PropertyType string
}

type Stat struct {
	CpuUsage string
	MemUsage string
	ExpiredDate string
}
type ClusterInfo struct {
	*types.RaftPeerInfo
	*Stat
}

type ResListProperty = struct {
	*types.ResWithoutData
	Data []*Property
}
type ResListClusterInfo = struct {
	*types.ResWithoutData
	Data []*ClusterInfo
}
type ResStat = struct {
	*types.ResWithoutData
	Data *Stat
}
type ResListSample = struct {
	*types.ResWithoutData
	Data *[]*map[string]interface{}
}