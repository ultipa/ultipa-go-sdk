package api

import (
	"io"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ExportAsNodesEdges(schema *structs.Schema, limit int, config *configuration.RequestConfig, cb func(nodes []*structs.Node, edges []*structs.Edge) error) error {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return err
	}
	defer cancel()

	properties := []string{}

	for _, prop := range schema.Properties {
		properties = append(properties, prop.Name)
	}

	resp, err := client.Export(ctx, &ultipa.ExportRequest{
		Schema:           schema.Name,
		Limit:            int32(limit),
		SelectProperties: properties,
		DbType:           schema.DBType,
	})

	if err != nil {
		return err
	}

	for {
		record, err := resp.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		wg := sync.WaitGroup{}
		if record.NodeTable != nil {
			wg.Add(len(record.NodeTable.NodeRows))

			//record.NodeTable
			nodeSchemaMap := structs.NewSchemaMapFromProtoSchema(record.NodeTable.Schemas, ultipa.DBType_DBNODE)
			nodes := make([]*structs.Node, len(record.NodeTable.NodeRows))
			for index, nodeRow := range record.NodeTable.NodeRows {

				go func(index int, row *ultipa.NodeRow) {
					defer wg.Done()
					node := structs.NewNodeFromNodeRow(nodeSchemaMap[schema.Name], row)
					nodes[index] = node
				}(index, nodeRow)

			}

			wg.Wait()
			err = cb(nodes, nil)
		}

		if record.EdgeTable != nil {
			wg.Add(len(record.EdgeTable.EdgeRows))
			//record.EdgeTable
			edgeSchemaMap := structs.NewSchemaMapFromProtoSchema(record.EdgeTable.Schemas, ultipa.DBType_DBEDGE)
			edges := make([]*structs.Edge, len(record.EdgeTable.EdgeRows))
			for index, edgeRow := range record.EdgeTable.EdgeRows {

				go func(index int, row *ultipa.EdgeRow) {
					defer wg.Done()
					edge := structs.NewEdgeFromEdgeRow(edgeSchemaMap[schema.Name], row)
					edges[index] = edge
				}(index, edgeRow)

			}

			wg.Wait()
			err = cb(nil, edges)
		}

		if err != nil {
			return err
		}
	}

	return err
}
