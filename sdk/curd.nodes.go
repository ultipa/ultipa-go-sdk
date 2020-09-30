package sdk

import (
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ResponseInsertNodes struct {
	*types.ResWithoutData
	Data []int64
}

func (t *Connection) InsertNodes(nodes []*map[string]interface{}, checkO bool, commonReq *types.Request_Common) *ResponseInsertNodes {

	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_insertNode)

	uql.SetCommandParams(nodes)

	res := t.UQL(uql.ToString(), commonReq)

	ids := []int64{}
	if len(*res.Data.Nodes) > 0 {
		for _, node := range *(*res.Data.Nodes)[0].Nodes {
			ids = append(ids, node.ID)
		}
	}

	return &ResponseInsertNodes{
		res.ResWithoutData,
		ids,
	}
}

type ResponseDeleteNodes struct {
	*types.ResWithoutData
}

type ResponseUpdateNodes struct {
	*types.ResWithoutData
}

func (t *Connection) UpdateNodes(filter interface{}, values map[string]interface{}, commonReq *types.Request_Common) *ResponseUpdateNodes {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_updateNodes)
	uql.SetCommandParams(filter)
	uql.AddParam("set", values, true)

	res := t.UQLListSample(uql.ToString(), commonReq)

	log.Println(uql.ToString(), res.Status.Message)

	return &ResponseUpdateNodes{
		res.ResWithoutData,
	}
}

func (t *Connection) DeleteNodes(filter interface{}, commonReq *types.Request_Common) *ResponseDeleteNodes {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_deleteNodes)
	uql.SetCommandParams(filter)

	res := t.UQLListSample(uql.ToString(), commonReq)

	log.Println(uql.ToString(), res.Status.Message)

	return &ResponseDeleteNodes{
		res.ResWithoutData,
	}
}

// InsertHugeNodes by one time
func (t *Connection) InsertHugeNodes(headers []string, rows [][]interface{}, silent bool, checkO bool, commonReq *types.Request_Common) *types.ResInsertHugeNodesReply {
	if commonReq == nil {
		commonReq = &types.Request_Common{}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		UseHost:        commonReq.UseHost,
		TimeoutSeconds: commonReq.TimeoutSeconds,
	})

	nodeTable := ultipa.NodeTable{}
	nodeRows := []*ultipa.NodeRow{}

	for _, header := range headers {
		if header == "_id" {
			// ignore node
		} else {
			nodeTable.Headers = append(nodeTable.Headers, &ultipa.Header{
				PropertyName: header,
			})
		}
	}

	for _, row := range rows {
		nodeRow := ultipa.NodeRow{}
		for index, header := range headers {
			if header == "_id" {
				nodeRow.Id = row[index].(int64)
			} else {
				nodeRow.Values = append(nodeRow.Values, row[index].([]byte))
			}
		}

		nodeRows = append(nodeRows, &nodeRow)
	}

	nodesRequest := ultipa.InsertNodesRequest{
		NodeTable: &nodeTable,
		Silent:    silent,
		CheckO:    checkO,
	}

	msg, err := clientInfo.ClientInfo.Client.InsertNodes(clientInfo.Context, &nodesRequest)

	res := &types.ResInsertHugeNodesReply{
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
