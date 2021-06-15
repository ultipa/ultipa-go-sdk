package http

import (
	"errors"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
)

type DataItem struct {
	Alias string
	Type ultipa.ResultType
	Data interface{}
}

func (di *DataItem) AsNodes() ([]*structs.Node, error){
	if di.Type != ultipa.ResultType_RESULT_TYPE_NODE {
		return nil, errors.New("DataItem " + di.Alias + " is not Type Node")
	}

	return di.Data.([]*structs.Node), nil
}

