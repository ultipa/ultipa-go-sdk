/**
 * Returns Uql Results by one time
 */

package http

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/types"
)

type InsertResponse struct {
	UUIDs     []types.UUID
	IDs       []types.ID
	ErrorItem map[int]int // index : error code
	Statistic *Statistic
	Status    *Status
}

var InsertErrorCodeMsgMap = map[int]string{
	//10001: "uuid and id not match",
	//10002: "id and schema not match",
	//10003: "from_id not exist",
	//10004: "to_id not exist",
	//10005: "id length exceed max length(128 bytes)",
	//10007: "duplicate ids in the data or the ids already exist in the database",
	// update in 2024-11-12
	//10001: "uuid and id not match",
	//10002: "id and schema not match",
	//10003: "from_id not exist",
	//10004: "to_id not exist",
	//10005: "id length exceed max length (128 bytes)",
	//10006: "some required field is null",
	//10007: "duplicate ids in the data or the ids already exist in the database",
	//10008: "id is empty",
	//10009: "from_id is empty",
	//10010: "to_id is empty",
	//10011: "duplicate id found in the data",
	//11001: "operation succeeded but the id already existed",
	//19999: "other error",
	// update in 2025-05-06
	10001: "ID_NOT_MATCH_UUID: id and uuid do not match",
	10002: "ID_UUID_NOT_MATCH_SCHEMA: id and schema do not match",
	10003: "FROM_ID_NOT_EXISTED: fromId does not exist",
	10004: "TO_ID_NOT_EXISTED: toId does not exist",
	10005: "ID_LEN: length of id exceeds the maximum length (128 bytes)",
	10006: "NOT_NULL: violation of the NOT NULL constraint",
	10007: "UNIQUCHECK: violation of the UNIQUE constraint",
	10008: "ID_EMPTY: id cannot be empty",
	10009: "FROM_ID_EMPTY: fromId cannot be empty",
	10010: "TO_ID_EMPTY: toId cannot be empty",
	10011: "DUPLICATE_ID: duplicated id",
	10012: "KEY_CONSTRAINT_VIOLATED: violation of the EDGE KEY constraint",
	11001: "OK_BUT_ID_EXISTED: id already exists",
	19999: "OTHERS: other error",
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
		UUIDs: reply.Uuids,
		IDs:   reply.Ids,
	}

	if len(reply.IgnoreIndexes) != len(reply.IgnoreErrorCode) {
		return response, nil
	}

	for index := range reply.IgnoreIndexes {
		v := int(reply.IgnoreIndexes[index])
		code := reply.IgnoreErrorCode[index]
		response.ErrorItem[v] = int(code)
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
		UUIDs:     reply.Uuids,
		ErrorItem: map[int]int{},
	}

	for index := range reply.IgnoreIndexes {
		v := int(reply.IgnoreIndexes[index])
		code := reply.IgnoreErrorCode[index]
		response.ErrorItem[v] = int(code)
	}

	return response, nil
}
