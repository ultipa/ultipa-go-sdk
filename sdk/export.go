package sdk

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
)

// ExportNodes  to export nodes
func (t *Connection) ExportNodes(properties []string, limit int32, cb func(rows []*ultipa.NodeRow, headers []*ultipa.Header) error, commonReq *types.Request_Common) (*ultipa.Status, error) {
	if commonReq == nil {
		commonReq = &types.Request_Common{
			GraphSetName: t.DefaultConfig.GraphSetName,
		}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		UseHost:        commonReq.UseHost,
		TimeoutSeconds: t.GetTimeOut(commonReq),
	})

	exportRequest := ultipa.ExportRequest{
		DbType:           ultipa.DBType_DBNODE,
		Limit:            limit,
		SelectProperties: properties,
	}

	msg, err := clientInfo.ClientInfo.Client.Export(clientInfo.Context, &exportRequest)

	if err != nil {
		return nil, err
	}

	index := 0
	headers := []*ultipa.Header{}

	for {
		reply, err := msg.Recv()

		if err == io.EOF {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		if len(headers) <= 0 && len(reply.NodeTable.Headers) > 0 {
			headers = reply.NodeTable.Headers
		}

		rows := reply.NodeTable.NodeRows

		err = cb(rows, headers)
		if err != nil {
			return nil, err
		}
		index++
	}
}

// ExportEdges  to export edges
func (t *Connection) ExportEdges(properties []string, limit int32, cb func(rows []*ultipa.EdgeRow, headers []*ultipa.Header) error, commonReq *types.Request_Common) (*ultipa.Status, error) {
	if commonReq == nil {
		commonReq = &types.Request_Common{
			GraphSetName: t.DefaultConfig.GraphSetName,
		}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		UseHost:        commonReq.UseHost,
		TimeoutSeconds: t.GetTimeOut(commonReq),
	})

	exportRequest := ultipa.ExportRequest{
		DbType:           ultipa.DBType_DBEDGE,
		Limit:            limit,
		SelectProperties: properties,
	}

	msg, err := clientInfo.ClientInfo.Client.Export(clientInfo.Context, &exportRequest)

	if err != nil {
		return nil, err
	}

	index := 0
	headers := []*ultipa.Header{}

	for {
		reply, err := msg.Recv()

		if err == io.EOF {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		if len(headers) <= 0 && len(reply.EdgeTable.Headers) > 0 {
			headers = reply.EdgeTable.Headers
		}

		rows := reply.EdgeTable.EdgeRows

		err = cb(rows, headers)
		if err != nil {
			return nil, err
		}
		index++
	}
}
