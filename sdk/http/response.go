package http

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/utils"
)

type UQLResponse struct {
	DataItemMap map[string]struct{
		DataItem *DataItem
		Index int
	}
	Reply *ultipa.UqlReply
	Status *Status
	Statistic *Statistic
	AliasList []string
	Resp ultipa.UltipaRpcs_UqlClient
}

func NewUQLResponse(resp ultipa.UltipaRpcs_UqlClient) (response *UQLResponse, err error) {

	response = &UQLResponse{
		Resp: resp,
		DataItemMap : map[string]struct {
			DataItem *DataItem
			Index    int
		}{},
	}

	for {
		record, err := resp.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil ,err
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

	return response, nil
}


func (r *UQLResponse) Init() {

}

func (r *UQLResponse) Get(index int) (di *DataItem) {
	return r.Alias(r.AliasList[index])
}

func (r *UQLResponse) Alias(alias string) (*DataItem) {

	data, t := utils.FindAliasDataInReply(r.Reply, alias)

	return &DataItem{
		Data: data,
		Type: t,
	}
}