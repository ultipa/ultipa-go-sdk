package sdk

import (
	"context"
	"io"
	"ultipa-go-sdk/utils"
	// "fmt"
	"log"
	"time"
	ultipa "ultipa-go-sdk/rpc"
)

type AttrGroup struct {
	Values []string
	Alias  string
}

type NodeGroup struct {
	Nodes []*utils.Node
	Alias string
}

type EdgeGroup struct {
	Edges []*utils.Edge
	Alias string
}

type TableGroup struct {
	TableName string
	Headers   []string
	Rows      [][]string
}

type UQLReply struct {
	Paths      []*utils.Path
	Nodes      []*NodeGroup
	Edges      []*EdgeGroup
	Attrs      []*AttrGroup
	Tables     []*TableGroup
	EngineCost int32
	TotalCost  int32
	Status     Status
}

func UQL(client ultipa.UltipaRpcsClient, uql string) UQLReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	msg, err := client.Uql(ctx, &ultipa.UqlRequest{
		Uql: uql,
	})

	if err != nil {
		log.Printf("uql error %v", err)
	}

	// parse paths
	res := UQLReply{}

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
		paths := utils.FormatPaths(c.Paths)

		// log.Printf("%#v", paths)

		for _, path := range paths {
			res.Paths = append(res.Paths, path)
		}

		// append Nodes
		for _, nodes := range c.Nodes {
			ns := utils.FormatNodes(nodes.Nodes)
			group := NodeGroup{
				Nodes: ns,
				Alias: nodes.Alias,
			}

			res.Nodes = append(res.Nodes, &group)
		}

		// append Edges
		for _, edges := range c.Edges {
			es := utils.FormatEdges(edges.Edges)
			group := EdgeGroup{
				Edges: es,
				Alias: edges.Alias,
			}

			res.Edges = append(res.Edges, &group)
		}

		// append Attrs
		for _, attrs := range c.Attrs {
			at := AttrGroup{
				Values: attrs.Values,
				Alias:  attrs.Alias,
			}
			res.Attrs = append(res.Attrs, &at)
		}

		for _, table := range c.Tables {
			tb := TableGroup{
				TableName: table.TableName,
				Headers:   table.Headers,
			}

			for _, row := range table.TableRows {
				tb.Rows = append(tb.Rows, row.Values)
			}

			res.Tables = append(res.Tables, &tb)
		}

		if res.EngineCost == 0 {
			res.EngineCost = c.EngineTimeCost
		}

		if res.TotalCost == 0 {
			res.TotalCost = c.TotalTimeCost
		}

		if c.Status != nil {
			res.Status = Status{
				ErrorCode: c.Status.ErrorCode,
				Msg:       c.Status.Msg,
			}
		}

	}

	return res

}
