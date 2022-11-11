/**
 * Returns UQL Results by one time
 */

package http

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/types"
)

type InsertResponse struct {
	Status    *Status
	Statistic *Statistic
	Data      struct {
		UUIDs     []types.UUID
		IDs       []types.ID
		ErrorItem map[int]int // index : error code
	}
}

var InsertErrorCodeMsgMap = map[int]string{
	10001: "uuid and id not match",
	10002: "uuid/id and schema not match",
	10003: "from_uuid and from_id not match, or not exist",
	10004: "to_uuid and to_id not match, or not exist",
	10005: "id length exceed max length(128 bytes)",
}

func NewNodesInsertResponse(reply *ultipa.InsertNodesReply) (response *InsertResponse, err error) {

	response = &InsertResponse{
		Status: &Status{
			Message: reply.Status.Msg,
			Code:    reply.Status.ErrorCode,
		},
		Statistic: &Statistic{
			TotalCost:  int(reply.TimeCost),
			EngineCost: int(reply.EngineTimeCost),
		},
		Data: struct {
			UUIDs     []types.UUID
			IDs       []types.ID
			ErrorItem map[int]int
		}{
			UUIDs:     reply.Uuids,
			IDs:       reply.Ids,
			ErrorItem: map[int]int{},
		},
	}

	for index := range reply.IgnoreIndexes {
		code := reply.IgnoreErrorCode[index]
		response.Data.ErrorItem[index] = int(code)
	}

	return response, nil
}

func NewEdgesInsertResponse(reply *ultipa.InsertEdgesReply) (response *InsertResponse, err error) {

	response = &InsertResponse{
		Status: &Status{
			Message: reply.Status.Msg,
			Code:    reply.Status.ErrorCode,
		},
		Statistic: &Statistic{
			TotalCost:  int(reply.TimeCost),
			EngineCost: int(reply.EngineTimeCost),
		},
		Data: struct {
			UUIDs     []types.UUID
			IDs       []types.ID
			ErrorItem map[int]int
		}{
			UUIDs:     reply.Uuids,
			ErrorItem: map[int]int{},
		},
	}

	for index := range reply.IgnoreIndexes {
		code := reply.IgnoreErrorCode[index]
		response.Data.ErrorItem[index] = int(code)
	}

	return response, nil
}
