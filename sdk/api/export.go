package api

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"google.golang.org/grpc"
	"io"
	"sync"
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

	var resp ultipa.UltipaControls_ExportClient
	var respErr error
	if config != nil && config.MaxPkgSize > 0 {
		resp, respErr = client.Export(ctx, &ultipa.ExportRequest{
			Schema:           schema.Name,
			Limit:            int32(limit),
			SelectProperties: properties,
			DbType:           schema.DBType,
		}, grpc.MaxCallRecvMsgSize(config.MaxPkgSize), grpc.MaxCallSendMsgSize(config.MaxPkgSize),
		)
	} else {
		resp, respErr = client.Export(ctx, &ultipa.ExportRequest{
			Schema:           schema.Name,
			Limit:            int32(limit),
			SelectProperties: properties,
			DbType:           schema.DBType,
		})
	}

	if respErr != nil {
		return respErr
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
			wg.Add(len(record.NodeTable.EntityRows))

			//record.NodeTable
			nodeSchemaMap := structs.NewSchemaMapFromProtoSchema(record.NodeTable.Schemas, ultipa.DBType_DBNODE)
			nodes := make([]*structs.Node, len(record.NodeTable.EntityRows))
			for index, nodeRow := range record.NodeTable.EntityRows {
				var parseErr error
				go func(index int, row *ultipa.EntityRow) {
					defer wg.Done()
					node, err := structs.NewNodeFromNodeRow(nodeSchemaMap[schema.Name], row)
					if err != nil {
						parseErr = err
					}
					nodes[index] = node
				}(index, nodeRow)
				if parseErr != nil {
					return parseErr
				}
			}

			wg.Wait()
			err = cb(nodes, nil)
		}

		if record.EdgeTable != nil {
			wg.Add(len(record.EdgeTable.EntityRows))
			//record.EdgeTable
			edgeSchemaMap := structs.NewSchemaMapFromProtoSchema(record.EdgeTable.Schemas, ultipa.DBType_DBEDGE)
			edges := make([]*structs.Edge, len(record.EdgeTable.EntityRows))
			for index, edgeRow := range record.EdgeTable.EntityRows {
				var parseErr error
				go func(index int, row *ultipa.EntityRow) {
					defer wg.Done()
					edge, err := structs.NewEdgeFromEdgeRow(edgeSchemaMap[schema.Name], row)
					if err != nil {
						parseErr = err
					}
					edges[index] = edge
				}(index, edgeRow)
				if parseErr != nil {
					return parseErr
				}
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
