package sdk

import (
	ultipa "ultipa-go-sdk/rpc"

	"google.golang.org/grpc"
)

// Client keep the connection to ultipa db host
type Client = ultipa.UltipaRpcsClient

// ClientConn is the connection , you can close it
type ClientConn = grpc.ClientConn

type Property struct {
	Name string
	Type PropertyType
}

type PropertyType = ultipa.UltipaPropertyType

const (
	PROPERTY_TYPE_INT     PropertyType = ultipa.UltipaPropertyType_PROPERTY_INT
	PROPERTY_TYPE_STRING  PropertyType = ultipa.UltipaPropertyType_PROPERTY_STRING
	PROPERTY_TYPE_TEXT    PropertyType = ultipa.UltipaPropertyType_PROPERTY_TEXT
	PROPERTY_TYPE_BOOLEAN PropertyType = ultipa.UltipaPropertyType_PROPERTY_BOOLEAN
	PROPERTY_TYPE_UNKNOWN PropertyType = ultipa.UltipaPropertyType_PROPERTY_UNKNOWN
)

type DBType = ultipa.DBType

const (
	DBType_DBNODE DBType = ultipa.DBType_DBNODE
	DBType_DBEDGE DBType = ultipa.DBType_DBEDGE
)

type Status struct {
	ErrorCode ErrorCode
	Msg       string
}

type ErrorCode = ultipa.ErrorCode

const (
	ErrorCodeSuccess ErrorCode = ultipa.ErrorCode_SUCCESS
	ErrorCodeFailed  ErrorCode = ultipa.ErrorCode_FAILED
	ErrorCodeSystem  ErrorCode = ultipa.ErrorCode_SYSTEM_ERROR
	ErrorCodeParam   ErrorCode = ultipa.ErrorCode_PARAM_ERROR
	ErrorCodeEngine  ErrorCode = ultipa.ErrorCode_ENGINE_ERROR
	ErrorCodeBaseDB  ErrorCode = ultipa.ErrorCode_BASE_DB_ERROR
)
