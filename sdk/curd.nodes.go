package sdk

import (
	"errors"
	"log"
	"strconv"
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

func (t *Connection) InsertHugeNodes(headers []ultipa.Header, rows [][]interface{}, silent bool, checkO bool, commonReq *types.Request_Common) (*types.ResInsertHugeNodesReply, error) {
	if commonReq == nil {
		commonReq = &types.Request_Common{
			GraphSetName: t.DefaultConfig.GraphSetName,
		}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		UseHost:        commonReq.UseHost,
		TimeoutSeconds: t.GetTimeOut(commonReq),
	})

	nodeTable := ultipa.NodeTable{}
	nodeRows := []*ultipa.NodeRow{}

	for index, _ := range headers {
		if headers[index].PropertyName == "_id" {
			// id header not add to header infos
		} else {
			nodeTable.Headers = append(nodeTable.Headers, &headers[index])
		}
	}

	for _, row := range rows {
		nodeRow := ultipa.NodeRow{}
		for index := range headers {
			h := headers[index]
			if h.PropertyName == "_id" {
				item := row[index]
				switch item.(type) {
				case int64:
					nodeRow.Id = item.(int64)
				case string:
					id, err := strconv.Atoi(item.(string))
					if err != nil {
						return nil, err
					}
					nodeRow.Id = int64(id)
				}

			} else {

				v := []byte{}
				item := row[index]
				// buff := new(bytes.Buffer)
				var err error
				// log.Println(item, h.PropertyName, h.PropertyType)
				v, err = utils.ConvertToBytes(item, h.PropertyType)

				// fmt.Println(nodeRow.Id, h.PropertyName, v)
				// if nodeRow.Id <= 0 && h.PropertyName == "_o" {
				// 	// nodeRow.Id = utils.Hash64(v)
				// }

				if err != nil {
					return nil, err
				} else {
					nodeRow.Values = append(nodeRow.Values, v)
				}

			}
		}
		nodeRows = append(nodeRows, &nodeRow)
	}

	nodeTable.NodeRows = nodeRows

	nodesRequest := ultipa.InsertNodesRequest{
		GraphName: t.DefaultConfig.GraphSetName,
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
