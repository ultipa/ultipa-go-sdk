/**
 * Return a Stream to return Uql results
 */

package http

import (
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"io"
)

type UQLResponseStream struct {
	DataItemMap map[string]struct {
		DataItem *DataItem
		Index    int
	}
	Reply     *ultipa.QueryReply
	Status    *Status
	Statistic *Statistic
	AliasList []string
	Resp      ultipa.UltipaRpcs_QueryClient
}

func NewUQLResponseStream(resp ultipa.UltipaRpcs_QueryClient) (response *UQLResponseStream, err error) {

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

func (r *UQLResponseStream) Recv(cb func(*Response) error) (err error) {
	//if !fetch {
	//	return nil, r.Resp.CloseSend()
	//}
	defer func() {
		_ = r.Resp.CloseSend()
	}()

	for {
		response := &Response{
			Status: &Status{},
			DataItemMap: map[string]struct {
				DataItem *DataItem
				Index    int
			}{},
		}

		record, err := r.Resp.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if response.Statistic == nil {
			response.Statistic, err = ParseStatistic(record.Statistics)
			if err != nil {
				return err
			}
		}

		if response.ExplainPlan == nil {
			response.ExplainPlan, err = ParseExplainPlan(record.ExplainPlan)
			if err != nil {
				return err
			}
		}

		response.Reply = record
		if record.Status != nil {
			response.Status.Code = record.Status.ErrorCode
			response.Status.Message = record.Status.Msg
			if response.Status.Code != ultipa.ErrorCode_SUCCESS {
				return fmt.Errorf(response.Status.Message)
			}
		}

		var aliasList []string

		for _, alias := range response.Reply.Alias {
			aliasList = append(aliasList, alias.GetAlias())
		}
		response.AliasList = aliasList

		if err := cb(response); err != nil {
			return err
		}

	}

	return nil

}

func (r *UQLResponseStream) NeedRedirect() bool {
	return r.Status.Code == ultipa.ErrorCode_RAFT_REDIRECT
}

func (r *UQLResponseStream) Close() error {
	return r.Resp.CloseSend()
}
