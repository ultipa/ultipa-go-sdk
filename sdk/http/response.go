/**
 * Returns UQL Results by one time
 */

package http

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

type UQLResponse struct {
	DataItemMap map[string]struct {
		DataItem *DataItem
		Index    int
	}
	Reply     *ultipa.UqlReply
	Status    *Status
	Statistic *Statistic
	AliasList []string
	Resp      ultipa.UltipaRpcs_UqlClient
}

func NewUQLResponse(resp ultipa.UltipaRpcs_UqlClient) (response *UQLResponse, err error) {

	response = &UQLResponse{
		Resp:   resp,
		Status: &Status{},
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

func (r *UQLResponse) Get(index int) (di *DataItem) {
	if len(r.AliasList) > index {
		return r.Alias(r.AliasList[index])
	}
	return nil
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
