package sdk

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func (t *Connection) UQLListSample(uql string, commonReq *SdkRequest_Common) *types.ResListSample  {
	res := t.UQL(uql, commonReq)
	uqlReply := res.Data
	var data *[]*map[string]interface{}
	if uqlReply != nil && uqlReply.Tables != nil && len(*uqlReply.Tables) > 0 {
		data = utils.TableToArray((*uqlReply.Tables)[0])
	}
	return &types.ResListSample{
		res.ResWithoutData,
		data,
	}
}
func (t *Connection) UQL(uql string, commonReq *SdkRequest_Common) *types.ResUqlReply {
	if commonReq == nil {
		commonReq = &SdkRequest_Common{}
	}
	clientInfo := t.getClientInfo(&GetClientInfoParams{
		Uql: uql,
		UseHost: commonReq.UseHost,
	})
	defer clientInfo.CancelFunc()
	retry := commonReq.Retry
	if retry == nil {
		retry = &Retry{
			Current: 0,
			Max: 3,
		}
	}
	isUseUqlExtra := UqlIsExtra(uql)
	//msg, err := clientInfo.ClientInfo.Client.Uql(clientInfo.Context, &ultipa.UqlRequest{
	//	Uql: uql,
	//	Timeout: t.DefaultConfig.TimeoutWithSeconds,
	//})
	var msg ultipa.UltipaRpcs_UqlClient
	var err error
	if isUseUqlExtra {
		msg, err = clientInfo.ClientInfo.Client.UqlEx(clientInfo.Context, &ultipa.UqlRequest{
			Uql: uql,
			Timeout: t.DefaultConfig.TimeoutWithSeconds,
		})
	} else {
		msg, err = clientInfo.ClientInfo.Client.Uql(clientInfo.Context, &ultipa.UqlRequest{
			Uql: uql,
			Timeout: t.DefaultConfig.TimeoutWithSeconds,
		})
	}
	//log.Printf("❗️ UQL: %s, host: %s, graphSetName: %s", uql, clientInfo.Host, clientInfo.GraphSetName)

	res := &types.ResUqlReply{
		ResWithoutData:&types.ResWithoutData{},
	}
	if t.DefaultConfig.ResponseWithRequestInfo {
		res.Req = &map[string]interface{}{
			"host": clientInfo.Host,
			"graphSetName": clientInfo.GraphSetName,
			"retry": retry,
			"uql": uql,
			"isUseUqlExtra": isUseUqlExtra,
		}
	}

	if err != nil {
		//log.Printf("uql error %v", err)
		res.Status = utils.FormatStatus(nil, err)
		return res
	}

	for {
		c, err := msg.Recv()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				//log.Printf("Failed %v \n", err)
				res.Status = utils.FormatStatus(nil, err)
				return res
				break
			}
		}
		//_json, _ := utils.StructToJSONString(c)
		//log.Printf("--uql原始response--\n %v \n %v \n", c, _json)
		if res.Status == nil {
			res.Status = utils.FormatStatus(c.Status, nil)
			res.EngineCost = c.GetEngineTimeCost()
			res.TotalCost = c.GetTotalTimeCost()
		}
		newUqlReply := &types.UqlReply{}
		newUqlReply.EngineCost = c.GetEngineTimeCost()
		newUqlReply.TotalCost = c.GetTotalTimeCost()
		newUqlReply.Paths = utils.FormatPaths(c.GetPaths())
		newUqlReply.Nodes = utils.FormatNodeAliases(c.GetNodes())
		newUqlReply.Edges = utils.FormatEdgeAliases(c.GetEdges())
		newUqlReply.Attrs = utils.FormatAttrs( c.GetAttrs())
		newUqlReply.Tables = utils.FormatTables(c.GetTables())
		newUqlReply.Values = utils.FormatKeyValues(c.GetKeyValues())

		if res.Data != nil {
			// append
			uqlReply := res.Data
			utils.UqlResponseAppend(uqlReply, newUqlReply)
			newUqlReply = uqlReply
		}
		res.Data = newUqlReply

	}
	return res

}

