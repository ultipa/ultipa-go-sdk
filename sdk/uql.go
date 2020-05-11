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
	// parse paths
	uqlReply := utils.UqlReply{}
	res := utils.Res{}
	for {
		c, err := msg.Recv()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Printf("Failed %v \n", err)
				break
			}
		}
		//_json, _ := utils.StructToJSONString(c)
		//log.Print(_json)
		// append Paths
		paths := utils.FormatPaths(c.Paths)
		// log.Printf("%#v", paths)
		for _, path := range paths {
			uqlReply.Paths = append(uqlReply.Paths, path)
		}
		// append Nodes
		for _, nodes := range c.Nodes {
			ns := utils.FormatNodes(nodes.Nodes)
			group := utils.NodeGroup{
				Nodes: ns,
				Alias: nodes.Alias,
			}
			uqlReply.Nodes = append(uqlReply.Nodes, &group)
		}
		// append Edges
		for _, edges := range c.Edges {
			es := utils.FormatEdges(edges.Edges)
			group := utils.EdgeGroup{
				Edges: es,
				Alias: edges.Alias,
			}
			uqlReply.Edges = append(uqlReply.Edges, &group)
		}

		// append Attrs
		for _, attrs := range c.Attrs {
			at := utils.AttrGroup{
				Values: attrs.Values,
				Alias:  attrs.Alias,
			}
			uqlReply.Attrs = append(uqlReply.Attrs, &at)
		}

		for _, table := range c.Tables {
			tb := utils.Table{
				TableName: table.TableName,
				Headers:   table.Headers,
			}

			for _, row := range table.TableRows {
				tb.TableRows = append(tb.TableRows, row.Values)
			}

			uqlReply.Tables = append(uqlReply.Tables, &tb)
		}
		uqlReply.Values = utils.FormatValues(c.Values)

		if res.EngineCost == 0 {
			res.EngineCost = c.EngineTimeCost
		}

		if res.TotalCost == 0 {
			res.TotalCost = c.TotalTimeCost
		}

		if c.Status != nil {
			res.Status = &utils.Status{
				Code: 			c.Status.ErrorCode,
				Message:       	c.Status.Msg,
			}
		}

	}
	res.Data = uqlReply
	return res


}
