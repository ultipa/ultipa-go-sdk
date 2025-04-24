/**
 * Returns Uql Results by one time
 */

package http

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
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
	//10001: "uuid and id not match",
	//10002: "id and schema not match",
	//10003: "from_id not exist",
	//10004: "to_id not exist",
	//10005: "id length exceed max length(128 bytes)",
	//10007: "duplicate ids in the data or the ids already exist in the database",
	// update in 2024-11-12
	10001: "uuid and id not match",
	10002: "id and schema not match",
	10003: "from_id not exist",
	10004: "to_id not exist",
	10005: "id length exceed max length (128 bytes)",
	10006: "some required field is null",
	10007: "duplicate ids in the data or the ids already exist in the database",
	10008: "id is empty",
	10009: "from_id is empty",
	10010: "to_id is empty",
	10011: "duplicate id found in the data",
	11001: "operation succeeded but the id already existed",
	19999: "other error",
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

	if len(reply.IgnoreIndexes) != len(reply.IgnoreErrorCode) {
		return response, nil
	}

	for index := range reply.IgnoreIndexes {
		v := int(reply.IgnoreIndexes[index])
		code := reply.IgnoreErrorCode[index]
		response.Data.ErrorItem[v] = int(code)
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
		v := int(reply.IgnoreIndexes[index])
		code := reply.IgnoreErrorCode[index]
		response.Data.ErrorItem[v] = int(code)
	}

	return response, nil
}
