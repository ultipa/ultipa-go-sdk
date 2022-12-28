package configuration

import ultipa "ultipa-go-sdk/rpc"

type RequestType = int32

const (
	RequestType_Write  RequestType = 1 // insert drop delete create truncate
	RequestType_Task   RequestType = 2 // exec task
	RequestType_Normal RequestType = 3 // search
)

type RequestConfig struct {
	GraphName      string      // Graphset Name
	Timeout        uint32      // timeout (Seconds)
	ClusterId      string      // Name Server Only
	Host           string      // set for force host test
	UseMaster      bool        // Use Master( graphSet master )
	UseControl     bool        // Use Control Node( global master )
	RequestType    RequestType // choose connection by request type, write => master, task > algo, normal => random
	Uql            string      // for Go Only, used for inner program
	Timezone       string      // name of time zone , e.g. Aisa/Shanghai
	TimezoneOffset int64         // seconds that elapse from UTC, prior to TimeZone
}

type InsertRequestConfig struct {
	*RequestConfig
	InsertType           ultipa.InsertType // used for insertBulkNodes/Edges
	CreateNodeIfNotExist bool              // used for insertBulkEdges
	Silent               bool              // if returns new ids
}

func (rc *RequestConfig) SetRequestTypeByUql(uql string) {

}
