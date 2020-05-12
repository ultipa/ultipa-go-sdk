package utils

import ultipa "ultipa-go-sdk/rpc"


type PropertyType = ultipa.UltipaPropertyType

const (
	PROPERTY_TYPE__INT32   PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT32
	PROPERTY_TYPE__STRING  PropertyType = ultipa.UltipaPropertyType_PROPERTY_STRING
	PROPERTY_TYPE__FLOAT   PropertyType = ultipa.UltipaPropertyType_PROPERTY_FLOAT
	PROPERTY_TYPE__DOUBLE  PropertyType = ultipa.UltipaPropertyType_PROPERTY_DOUBLE
	PROPERTY_TYPE__UINT32  PropertyType = ultipa.UltipaPropertyType_PROPERTY_UINT32
	PROPERTY_TYPE__INT64   PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT64
	PROPERTY_TYPE__UINT64  PropertyType = ultipa.UltipaPropertyType_PROPERTY_UINT64
	PROPERTY_TYPE__UNKNOWN PropertyType = ultipa.UltipaPropertyType_PROPERTY_UNKNOWN
)

type DBType = ultipa.DBType

const (
	DBType_DBNODE DBType = ultipa.DBType_DBNODE
	DBType_DBEDGE DBType = ultipa.DBType_DBEDGE
)

type ErrorCode = ultipa.ErrorCode

const (
	ErrorCode_SUCCESS        ErrorCode = ultipa.ErrorCode_SUCCESS
	ErrorCode_FAILED         ErrorCode = ultipa.ErrorCode_FAILED
	ErrorCode_PARAM_ERROR    ErrorCode = ultipa.ErrorCode_PARAM_ERROR
	ErrorCode_BASE_DB_ERROR  ErrorCode = ultipa.ErrorCode_BASE_DB_ERROR
	ErrorCode_ENGINE_ERROR   ErrorCode = ultipa.ErrorCode_ENGINE_ERROR
	ErrorCode_SYSTEM_ERROR   ErrorCode = ultipa.ErrorCode_SYSTEM_ERROR
	ErrorCode_RAFT_REDIRECT  ErrorCode = ultipa.ErrorCode_RAFT_REDIRECT
	ErrorCode_RAFT_LEADER_NOT_YET_ELECTED 	ErrorCode = ultipa.ErrorCode_RAFT_LEADER_NOT_YET_ELECTED
	ErrorCode_RAFT_LOG_ERROR 				ErrorCode = ultipa.ErrorCode_RAFT_LOG_ERROR
	ErrorCode_UQL_ERROR      				ErrorCode = ultipa.ErrorCode_UQL_ERROR
	ErrorCode_NOT_RAFT_MODE      		ErrorCode = ultipa.ErrorCode_NOT_RAFT_MODE
)


type ClusterInfo struct {
	Redirect string
	RaftPeers []string
}

type Status = struct {
	Code    ErrorCode
	Message string
	ClusterInfo *ClusterInfo
}

type Node = struct {
	ID string
	Values *map[string]string
}
type Nodes = []*Node
type Edge = struct {
	ID string
	From string
	To string
	Values *map[string]string
}
type Edges = []*Edge
type Path = struct {
	Nodes Nodes;
	Edges Edges;
}

type Paths = []*Path

type Res = struct {
	Status *Status
	Data interface{}
	TotalCost int32
	EngineCost int32
}
type Table = struct {
	TableName string
	Headers []string
	TableRows [][]string
}

type AttrGroup struct {
	Values []string
	Alias  string
}

type NodeGroup struct {
	Nodes Nodes
	Alias string
}

type EdgeGroup struct {
	Edges Edges
	Alias string
}

type UqlReply struct {
	Paths  Paths
	Nodes  []*NodeGroup
	Edges  []*EdgeGroup
	Attrs  []*AttrGroup
	Tables []*Table
	Values *map[string]string
}

type Property struct {
	PropertyName string
	PropertyType string
	Lte bool
	Index bool
}
