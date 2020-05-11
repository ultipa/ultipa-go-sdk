package sdk

import (
	"io"
	"log"
	ultipa "ultipa-go-sdk/rpc"
)

func (t *Connection) UQL(uql string) Res {
	clientInfo, ctx, cancel := t.choiseClient(TIMEOUT_DEFAUL)
	defer cancel()
	msg, err := clientInfo.Client.Uql(ctx, &ultipa.UqlRequest{
		Uql: uql,
	})

	if err != nil {
		log.Printf("uql error %v", err)
	}
	// parse paths
	uqlReply := UqlReply{}
	res := Res{}
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
		// append Paths
		paths := FormatPaths(c.Paths)
		// log.Printf("%#v", paths)
		for _, path := range paths {
			uqlReply.Paths = append(uqlReply.Paths, path)
		}
		// append Nodes
		for _, nodes := range c.Nodes {
			ns := FormatNodes(nodes.Nodes)
			group := NodeGroup{
				Nodes: ns,
				Alias: nodes.Alias,
			}
			uqlReply.Nodes = append(uqlReply.Nodes, &group)
		}
		// append Edges
		for _, edges := range c.Edges {
			es := FormatEdges(edges.Edges)
			group := EdgeGroup{
				Edges: es,
				Alias: edges.Alias,
			}
			uqlReply.Edges = append(uqlReply.Edges, &group)
		}

		// append Attrs
		for _, attrs := range c.Attrs {
			at := AttrGroup{
				Values: attrs.Values,
				Alias:  attrs.Alias,
			}
			uqlReply.Attrs = append(uqlReply.Attrs, &at)
		}

		for _, table := range c.Tables {
			tb := Table{
				TableName: table.TableName,
				Headers:   table.Headers,
			}

			for _, row := range table.TableRows {
				tb.TableRows = append(tb.TableRows, row.Values)
			}

			uqlReply.Tables = append(uqlReply.Tables, &tb)
		}

		if res.EngineCost == 0 {
			res.EngineCost = c.EngineTimeCost
		}

		if res.TotalCost == 0 {
			res.TotalCost = c.TotalTimeCost
		}

		if c.Status != nil {
			res.Status = &Status{
				Code: 			c.Status.ErrorCode,
				Message:       	c.Status.Msg,
			}
		}

	}
	res.Data = uqlReply
	return res


}
