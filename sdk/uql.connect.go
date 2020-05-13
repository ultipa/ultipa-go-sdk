package sdk

import (
	"io"
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/utils"
)

func (t *Connection) UQL(uql string) utils.Res {
	clientInfo, ctx, cancel := t.chooseClient(TIMEOUT_DEFAUL)
	defer cancel()
	msg, err := clientInfo.Client.Uql(ctx, &ultipa.UqlRequest{
		Uql: uql,
	})

	if err != nil {
		log.Printf("uql error %v", err)
	}

	res := utils.Res{}
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
		var uqlReplys []*utils.UqlReply
		for _, uqlReply := range c.UqlData {
			newUqlReply := utils.UqlReply{}
			newUqlReply.SequenceId = uqlReply.GetSequenceId()
			newUqlReply.EngineCost = uqlReply.GetEngineTimeCost()
			newUqlReply.TotalCost = uqlReply.GetTotalTimeCost()
			newUqlReply.Paths = utils.FormatPaths(uqlReply.GetPaths())
			newUqlReply.NodeAliases = utils.FormatNodeAliases(uqlReply.GetNodes())
			newUqlReply.EdgeAliases = utils.FormatEdgeAliases(uqlReply.GetEdges())
			newUqlReply.Attrs = utils.FormatAttrs( uqlReply.GetAttrs())
			newUqlReply.Tables = utils.FormatTables(uqlReply.GetTables())
			newUqlReply.KeyValues = utils.FormatKeyValues(uqlReply.GetKeyValues())
			uqlReplys = append(uqlReplys, &newUqlReply)
		}
		res.Data = uqlReplys
		if c.Status != nil {
			res.Status = utils.FormatStatus(c.Status, nil)
		}

	}
	return res

}
