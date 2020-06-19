package sdk

import (
	"io"
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func (t *Connection) UQL(uql string) types.ResAny {
	clientInfo, ctx, cancel := t.chooseClient(TIMEOUT_DEFAUL)
	defer cancel()
	msg, err := clientInfo.Client.Uql(ctx, &ultipa.UqlRequest{
		Uql: uql,
		Timeout: t.DefaultConfig.TimeoutWithSeconds,
	})

	if err != nil {
		log.Printf("uql error %v", err)
	}

	res := types.ResAny{}
	for {
		c, err := msg.Recv()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Printf("Failed %v \n", err)
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
		newUqlReply := types.UqlReply{}
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
			uqlReply := res.Data.(types.UqlReply)
			utils.UqlResponseAppend(&uqlReply, &newUqlReply)
			newUqlReply = uqlReply
		}
		res.Data = newUqlReply

	}
	return res

}

