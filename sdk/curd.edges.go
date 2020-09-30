package sdk

import (
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
func (t *Connection) InsertHugeEdges(headers []string, rows [][]interface{}, silent bool, checkO bool, commonReq *types.Request_Common) *types.ResInsertHugeEdgesReply {
	if commonReq == nil {
		commonReq = &types.Request_Common{}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		UseHost:        commonReq.UseHost,
		TimeoutSeconds: commonReq.TimeoutSeconds,
	})

	edgeTable := ultipa.EdgeTable{}
	edgeRows := []*ultipa.EdgeRow{}

	for _, header := range headers {
		if header == "_id" || header == "_from_id" || header == "_to_id" {
		} else {
			edgeTable.Headers = append(edgeTable.Headers, &ultipa.Header{
				PropertyName: header,
			})
		}
	}

	for _, row := range rows {
		edgeRow := ultipa.EdgeRow{}
		for index, header := range headers {
			switch header {
			case "_id":
				edgeRow.Id = row[index].(int64)
			case "_from_id":
				edgeRow.FromId = row[index].(int64)
			case "_to_id":
				edgeRow.ToId = row[index].(int64)
			default:
				edgeRow.Values = append(edgeRow.Values, row[index].([]byte))
			}
		}

		edgeRows = append(edgeRows, &edgeRow)
	}

	EdgesRequest := ultipa.InsertEdgesRequest{
		EdgeTable: &edgeTable,
		Silent:    silent,
	}

	msg, err := clientInfo.ClientInfo.Client.InsertEdges(clientInfo.Context, &EdgesRequest)

	res := &types.ResInsertHugeEdgesReply{
		ResWithoutData: &types.ResWithoutData{},
	}

	if err != nil {
		//log.Printf("uql error %v", err)
		res.Status = utils.FormatStatus(nil, err)
		return res
	}

	if res.Status == nil {
		res.Status = utils.FormatStatus(msg.Status, nil)
		res.EngineCost = msg.GetEngineTimeCost()
		res.TotalCost = msg.GetTimeCost()
	}

	res.Data.Ids = msg.GetIds()
	res.Data.IgnoreIndexes = msg.GetIgnoreIndexes()

	return res
}
