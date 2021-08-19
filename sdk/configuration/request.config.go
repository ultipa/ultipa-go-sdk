package configuration

import ultipa "ultipa-go-sdk/rpc"

type RequestType = int32

const (
	RequestType_Write  RequestType = 1 // insert drop delete create truncate
	RequestType_Task   RequestType = 2 // exec task
	RequestType_Normal RequestType = 3 // search
)

type RequestConfig struct {
	GraphName string
	Timeout   uint32
	ClusterId string
	Host      string // set for force host test
	RequestType RequestType
	InsertType ultipa.InsertType
	CreateNodeIfNotExist bool
	Uql string
}

func (rc * RequestConfig) SetRequestTypeByUql(uql string) {

}
