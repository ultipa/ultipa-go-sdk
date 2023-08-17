/**
 * Returns UQL Results by one time
 */

package http

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"io"
)

type UQLResponse struct {
	DataItemMap map[string]struct {
		DataItem *DataItem
		Index    int
	}
	Reply       *ultipa.UqlReply
	Status      *Status
	Statistic   *Statistic
	ExplainPlan *ExplainPlan
	AliasList   []string
	Resp        ultipa.UltipaRpcs_UqlClient
}

func NewUQLResponse(resp ultipa.UltipaRpcs_UqlClient) (response *UQLResponse, err error) {

	response = &UQLResponse{
		Resp:   resp,
		Status: &Status{},
		//Reply:  &ultipa.UqlReply{},
		DataItemMap: map[string]struct {
			DataItem *DataItem
			Index    int
		}{},
	}
	for {

		record, err := resp.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if response.Statistic == nil {
			response.Statistic, err = ParseStatistic(record.Statistics)
			if err != nil {
				return nil, err
			}
		}

		if response.ExplainPlan == nil {
			response.ExplainPlan, err = ParseExplainPlan(record.ExplainPlan)
			if err != nil {
				return nil, err
			}
		}

		if response.Reply == nil {
			response.Reply = record
		} else {
			response.Reply = utils.MergeUQLReply(response.Reply, record)
		}

		response.Status.Code = record.Status.ErrorCode
		response.Status.Message = record.Status.Msg

		if response.Status.Code != ultipa.ErrorCode_SUCCESS {
			return response, nil
		}
	}

	var aliasList []string

	for _, alias := range response.Reply.Alias {
		aliasList = append(aliasList, alias.GetAlias())
	}

	response.AliasList = aliasList

	return response, nil
}

func (r *UQLResponse) NeedRedirect() bool {
	return r.Status.Code == ultipa.ErrorCode_RAFT_REDIRECT
}

func (r *UQLResponse) IsSuccess() bool {
	if r != nil && r.Status != nil {
		return r.Status.Code == ultipa.ErrorCode_SUCCESS
	}
	return false

}

func (r *UQLResponse) Get(index int) (di *DataItem) {
	if len(r.AliasList) > index {
		return r.Alias(r.AliasList[index])
	}
	return &DataItem{
		Data: nil,
		Type: ultipa.ResultType_RESULT_TYPE_UNSET,
	}
}

func (r *UQLResponse) Alias(alias string) *DataItem {

	data, t := utils.FindAliasDataInReply(r.Reply, alias)

	return &DataItem{
		Data: data,
		Type: t,
	}
}

func (r *UQLResponse) GetSingleTable() (*structs.Table, error) {
	di := r.Get(0)
	if di != nil {
		t, err := di.AsTable()
		return t, err
	}
	return nil, nil
}
