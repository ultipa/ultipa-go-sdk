package sdk

import (
	"errors"
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ResponseInsertEdges struct {
	*types.ResWithoutData
	Data []int64
}

func (t *Connection) InsertEdges(Edges []*map[string]interface{}, checkO bool, commonReq *types.Request_Common) *ResponseInsertEdges {

	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_insertEdge)

	uql.SetCommandParams(Edges)

	res := t.UQL(uql.ToString(), commonReq)

	ids := []int64{}
	if len(*res.Data.Edges) > 0 {
		for _, edge := range *(*res.Data.Edges)[0].Edges {
			ids = append(ids, edge.ID)
		}
	}

	return &ResponseInsertEdges{
		res.ResWithoutData,
		ids,
	}
}

type ResponseDeleteEdges struct {
	*types.ResWithoutData
}

type ResponseUpdateEdges struct {
	*types.ResWithoutData
}

func (t *Connection) UpdateEdges(filter interface{}, values map[string]interface{}, commonReq *types.Request_Common) *ResponseUpdateEdges {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_updateEdges)
	uql.SetCommandParams(filter)
	uql.AddParam("set", values, true)

	res := t.UQLListSample(uql.ToString(), commonReq)

	log.Println(uql.ToString(), res.Status.Message)

	return &ResponseUpdateEdges{
		res.ResWithoutData,
	}
}

func (t *Connection) DeleteEdges(filter interface{}, commonReq *types.Request_Common) *ResponseDeleteEdges {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_deleteEdges)
	uql.SetCommandParams(filter)

	res := t.UQLListSample(uql.ToString(), commonReq)

	log.Println(uql.ToString(), res.Status.Message)

	return &ResponseDeleteEdges{
		res.ResWithoutData,
	}
}

// InsertHugeEdges by one time
func (t *Connection) InsertHugeEdges(headers []ultipa.Header, rows [][]interface{}, silent bool, commonReq *types.Request_Common) (*types.ResInsertHugeEdgesReply, error) {
	if commonReq == nil {
		commonReq = &types.Request_Common{
			GraphSetName: t.DefaultConfig.GraphSetName,
		}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		GraphSetName:   commonReq.GraphSetName,
		ClientType:     ClientType_Update,
		TimeoutSeconds: t.GetTimeOut(commonReq),
	})

	edgeTable := ultipa.EdgeTable{}
	edgeRows := []*ultipa.EdgeRow{}

	for index, _ := range headers {
		header := headers[index]
		if header.PropertyName == "_id" || header.PropertyName == "_from_id" || header.PropertyName == "_to_id" {
		} else {
			edgeTable.Headers = append(edgeTable.Headers, &header)
		}
	}

	for _, row := range rows {
		edgeRow := ultipa.EdgeRow{}
		for index, _ := range headers {
			header := headers[index]
			value := row[index]

			switch header.PropertyName {
			case "_id":
				edgeRow.Id = utils.ConvertToID(value)
			case "_from_id":
				edgeRow.FromId = utils.ConvertToID(value)
			case "_to_id":
				edgeRow.ToId = utils.ConvertToID(value)
			default:
				v, err := utils.ConvertToBytes(value, header.PropertyType)

				if err != nil {
					return nil, err
				}

				// if edgeRow.FromId < 0 && header.PropertyName == "_from_o" {
				// 	edgeRow.FromId = utils.Hash64(v)
				// } else if edgeRow.ToId < 0 && header.PropertyName == "_to_o" {
				// 	edgeRow.ToId = utils.Hash64(v)
				// } else {
				// }
				edgeRow.Values = append(edgeRow.Values, v)
			}
		}

		edgeRows = append(edgeRows, &edgeRow)
	}

	edgeTable.EdgeRows = edgeRows
	edgesRequest := ultipa.InsertEdgesRequest{
		GraphName: commonReq.GraphSetName,
		EdgeTable: &edgeTable,
		Silent:    silent,
	}

	// utils.PrintJSON(edgesRequest)

	msg, err := clientInfo.ClientInfo.Client.InsertEdges(clientInfo.Context, &edgesRequest)

	res := &types.ResInsertHugeEdgesReply{
		ResWithoutData: &types.ResWithoutData{},
	}

	// utils.PrintJSON(msg)
	if err != nil {
		//log.Printf("uql error %v", err)
		res.Status = utils.FormatStatus(nil, err)
		return res, err
	}

	res.Status = utils.FormatStatus(msg.Status, nil)

	if res.Status.Code != ultipa.ErrorCode_SUCCESS {
		res.Status = utils.FormatStatus(msg.Status, nil)
		return res, errors.New(res.Status.Message)
	}

	res.EngineCost = msg.GetEngineTimeCost()
	res.TotalCost = msg.GetTimeCost()
	res.Data.Ids = msg.GetIds()
	res.Data.IgnoreIndexes = msg.GetIgnoreIndexes()

	return res, nil
}
