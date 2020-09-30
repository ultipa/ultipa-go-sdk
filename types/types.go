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

type PropertyTypeString string

const (
	PROPERTY_TYPE_INT32_STRING  PropertyTypeString = "int32"
	PROPERTY_TYPE_STRING_STRING PropertyTypeString = "string"
	PROPERTY_TYPE_FLOAT_STRING  PropertyTypeString = "float"
	PROPERTY_TYPE_DOUBLE_STRING PropertyTypeString = "double"
	PROPERTY_TYPE_UINT32_STRING PropertyTypeString = "uint32"
	PROPERTY_TYPE_INT64_STRING  PropertyTypeString = "int64"
	PROPERTY_TYPE_UINT64_STRING PropertyTypeString = "uint64"
	PROPERTY_TYPE_BLOB_STRING   PropertyTypeString = "blob"
)

type TASK_STATUS = ultipa.TASK_STATUS

const (
	TASK_STATUS_TASK_PENDING   TASK_STATUS = ultipa.TASK_STATUS_TASK_PENDING
	TASK_STATUS_TASK_COMPUTING TASK_STATUS = ultipa.TASK_STATUS_TASK_COMPUTING
	TASK_STATUS_TASK_WRITING   TASK_STATUS = ultipa.TASK_STATUS_TASK_WRITING
	TASK_STATUS_TASK_DONE      TASK_STATUS = ultipa.TASK_STATUS_TASK_DONE
	TASK_STATUS_TASK_FAILED    TASK_STATUS = ultipa.TASK_STATUS_TASK_FAILED
	TASK_STATUS_TASK_STOPPED   TASK_STATUS = ultipa.TASK_STATUS_TASK_STOPPED
)

var TASK_STATUS_MAP = map[TASK_STATUS]string{
	TASK_STATUS_TASK_PENDING:   "TASK_PENDING",
	TASK_STATUS_TASK_COMPUTING: "TASK_COMPUTING",
	TASK_STATUS_TASK_WRITING:   "TASK_WRITING",
	TASK_STATUS_TASK_DONE:      "TASK_DONE",
	TASK_STATUS_TASK_FAILED:    "TASK_FAILED",
	TASK_STATUS_TASK_STOPPED:   "TASK_STOPPED",
}

type DBType = ultipa.DBType

const (
	DBType_DBNODE DBType = ultipa.DBType_DBNODE
	DBType_DBEDGE DBType = ultipa.DBType_DBEDGE
)

type ErrorCode = ultipa.ErrorCode

const (
	ErrorCode_SUCCESS                        ErrorCode = ultipa.ErrorCode_SUCCESS
	ErrorCode_FAILED                         ErrorCode = ultipa.ErrorCode_FAILED
	ErrorCode_PARAM_ERROR                    ErrorCode = ultipa.ErrorCode_PARAM_ERROR
	ErrorCode_BASE_DB_ERROR                  ErrorCode = ultipa.ErrorCode_BASE_DB_ERROR
	ErrorCode_ENGINE_ERROR                   ErrorCode = ultipa.ErrorCode_ENGINE_ERROR
	ErrorCode_SYSTEM_ERROR                   ErrorCode = ultipa.ErrorCode_SYSTEM_ERROR
	ErrorCode_RAFT_REDIRECT                  ErrorCode = ultipa.ErrorCode_RAFT_REDIRECT
	ErrorCode_RAFT_LEADER_NOT_YET_ELECTED    ErrorCode = ultipa.ErrorCode_RAFT_LEADER_NOT_YET_ELECTED
	ErrorCode_RAFT_LOG_ERROR                 ErrorCode = ultipa.ErrorCode_RAFT_LOG_ERROR
	ErrorCode_UQL_ERROR                      ErrorCode = ultipa.ErrorCode_UQL_ERROR
	ErrorCode_NOT_RAFT_MODE                  ErrorCode = ultipa.ErrorCode_NOT_RAFT_MODE
	ErrorCode_RAFT_NO_AVAILABLE_FOLLOWERS    ErrorCode = ultipa.ErrorCode_RAFT_NO_AVAILABLE_FOLLOWERS
	ErrorCode_RAFT_NO_AVAILABLE_ALGO_SERVERS ErrorCode = ultipa.ErrorCode_RAFT_NO_AVAILABLE_ALGO_SERVERS
	ErrorCode_PERMISSION_DENIED              ErrorCode = ultipa.ErrorCode_PERMISSION_DENIED

	ErrorCode_UNKNOW ErrorCode = 1000
)

type RAFT_FOLLOWER_ROLE = ultipa.FollowerRole

const (
	RAFT_FOLLOWER_ROLE_UNSET           RAFT_FOLLOWER_ROLE = ultipa.FollowerRole_ROLE_UNSET
	RAFT_FOLLOWER_ROLE_READABLE        RAFT_FOLLOWER_ROLE = ultipa.FollowerRole_ROLE_READABLE
	RAFT_FOLLOWER_ROLE_ALGO_EXECUTABLE RAFT_FOLLOWER_ROLE = ultipa.FollowerRole_ROLE_ALGO_EXECUTABLE
)

type RaftPeerInfo struct {
	Host               string
	Status             bool
	IsLeader           bool
	IsFollowerReadable bool
	IsAlgoExecutable   bool
	IsUnset            bool
}

type ClusterInfo struct {
	Redirect  string
	RaftPeers []*RaftPeerInfo
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
	Alias string
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
	Alias string
}
type Path = struct {
	Nodes *NodeTable
	Edges *EdgeTable
}
type Paths = []*Path

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
type EdgeAliases []*EdgeAlias
type UqlReply struct {
	TotalCost  int32
	EngineCost int32

	Paths  *Paths
	Nodes  *NodeAliases
	Edges  *EdgeAliases
	Attrs  *Attrs
	Tables *Tables
	Values *map[string]interface{}
}

type ResWithoutData = struct {
	Status     *Status
	TotalCost  int32
	EngineCost int32
	Req        *map[string]interface{}
}

type ResUqlReply = struct {
	*ResWithoutData
	Data *UqlReply
}

type ResInsertHugeNodesReply = struct {
	*ResWithoutData
	Data struct {
		Ids           []int64
		IgnoreIndexes []int32
	}
}

type ResInsertHugeEdgesReply = struct {
	*ResWithoutData
	Data struct {
		Ids           []int64
		IgnoreIndexes []int32
	}
}
