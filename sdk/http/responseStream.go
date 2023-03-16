/**
 * Return a Stream to return UQL results
 */

package http

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
)

type UQLResponseStream struct {
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

func NewUQLResponseStream(resp ultipa.UltipaRpcs_UqlClient) (response *UQLResponseStream, err error) {

	response = &UQLResponseStream{
		Resp:   resp,
		Status: &Status{},
		DataItemMap: map[string]struct {
			DataItem *DataItem
			Index    int
		}{},
	}

	return response, nil
}

func (r *UQLResponseStream) Recv(fetch bool) (response *UQLResponse, err error) {
	if !fetch {
		return nil, r.Resp.CloseSend()
	}
	response = &UQLResponse{
		Status: &Status{},
		DataItemMap: map[string]struct {
			DataItem *DataItem
			Index    int
		}{},
	}

	record, err := r.Resp.Recv()

	if err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	response.Reply = record
	response.Status.Code = record.Status.ErrorCode
	response.Status.Message = record.Status.Msg

	if response.Status.Code != ultipa.ErrorCode_SUCCESS {
		return response, nil
	}

	var aliasList []string

	for _, alias := range response.Reply.Alias {
		aliasList = append(aliasList, alias.GetAlias())
	}
	response.AliasList = aliasList

	return response, nil
}

func (r *UQLResponseStream) NeedRedirect() bool {
	return r.Status.Code == ultipa.ErrorCode_RAFT_REDIRECT
}

func (r *UQLResponseStream) Close() error {
	return r.Resp.CloseSend()
}
