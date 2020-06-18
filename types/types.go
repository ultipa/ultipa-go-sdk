package types

import (
	ultipa "ultipa-go-sdk/rpc"
)

type PropertyType = ultipa.UltipaPropertyType

const (
	PROPERTY_TYPE_INT32  PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT32
	PROPERTY_TYPE_STRING PropertyType = ultipa.UltipaPropertyType_PROPERTY_STRING
	PROPERTY_TYPE_FLOAT  PropertyType = ultipa.UltipaPropertyType_PROPERTY_FLOAT
	PROPERTY_TYPE_DOUBLE PropertyType = ultipa.UltipaPropertyType_PROPERTY_DOUBLE
	PROPERTY_TYPE_UINT32 PropertyType = ultipa.UltipaPropertyType_PROPERTY_UINT32
	PROPERTY_TYPE_INT64  PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT64
	PROPERTY_TYPE_UINT64 PropertyType = ultipa.UltipaPropertyType_PROPERTY_UINT64
	PROPERTY_TYPE_BLOB   PropertyType = ultipa.UltipaPropertyType_PROPERTY_BLOB
)

type DBType = ultipa.DBType

const (
	DBType_DBNODE DBType = ultipa.DBType_DBNODE
	DBType_DBEDGE DBType = ultipa.DBType_DBEDGE
)

type ErrorCode = ultipa.ErrorCode

const (
	ErrorCode_SUCCESS                     ErrorCode = ultipa.ErrorCode_SUCCESS
	ErrorCode_FAILED                      ErrorCode = ultipa.ErrorCode_FAILED
	ErrorCode_PARAM_ERROR                 ErrorCode = ultipa.ErrorCode_PARAM_ERROR
	ErrorCode_BASE_DB_ERROR               ErrorCode = ultipa.ErrorCode_BASE_DB_ERROR
	ErrorCode_ENGINE_ERROR                ErrorCode = ultipa.ErrorCode_ENGINE_ERROR
	ErrorCode_SYSTEM_ERROR                ErrorCode = ultipa.ErrorCode_SYSTEM_ERROR
	ErrorCode_RAFT_REDIRECT               ErrorCode = ultipa.ErrorCode_RAFT_REDIRECT
	ErrorCode_RAFT_LEADER_NOT_YET_ELECTED ErrorCode = ultipa.ErrorCode_RAFT_LEADER_NOT_YET_ELECTED
	ErrorCode_RAFT_LOG_ERROR              ErrorCode = ultipa.ErrorCode_RAFT_LOG_ERROR
	ErrorCode_UQL_ERROR                   ErrorCode = ultipa.ErrorCode_UQL_ERROR
	ErrorCode_NOT_RAFT_MODE               ErrorCode = ultipa.ErrorCode_NOT_RAFT_MODE
)

type ClusterInfo struct {
	Redirect  string
	RaftPeers []string
}

type Status = struct {
	Code        ErrorCode
	Message     string
	ClusterInfo *ClusterInfo
}

type NodeRow = struct {
	ID     int64
	Values *map[string]interface{}
}
type NodeTable = []*NodeRow
type NodeAlias = struct {
	Nodes *NodeTable
	Alias     string
}

type EdgeRow = struct {
	ID     int64
	From   int64
	To     int64
	Values *map[string]interface{}
}
type EdgeTable = []*EdgeRow
type EdgeAlias = struct {
	Edges *EdgeTable
	Alias     string
}
type Path = struct {
	NodeTable *NodeTable
	EdgeTable *EdgeTable
}
type Paths = []*Path

type ResAny = struct {
	Status *Status
	TotalCost  int32
	EngineCost int32
	Req interface{}
	Data   interface{}
}

type TableRows []*[]interface{}
type Table = struct {
	TableName string
	Headers   []string
	TableRows *TableRows
}

type AttrAlias struct {
	Values []interface{}
	Alias  string
}
type Attrs []*AttrAlias
type Tables []*Table
type NodeAliases []*NodeAlias
type EdgeAliases  []*EdgeAlias
type UqlReply struct {
	TotalCost  int32
	EngineCost int32

	Paths       *Paths
	Nodes *NodeAliases
	Edges *EdgeAliases
	Attrs       *Attrs
	Tables      *Tables
	Values     *map[string]interface{}
}

type Property struct {
	PropertyName string
	PropertyType string
	Lte          bool
	Index        bool
}

