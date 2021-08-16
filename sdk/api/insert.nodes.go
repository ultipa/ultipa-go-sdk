package api

import (
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.NodeTable, config *configuration.RequestConfig) (*ultipa.InsertNodesReply, error) {

	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext(config)

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName: conf.CurrentGraph,
		NodeTable: table,
		Silent:    true,
	})

	return resp, err
}

func (api *UltipaAPI) InsertNodesBatchBySchema(schema *structs.Schema, rows []*structs.Node, config *configuration.RequestConfig) (*ultipa.InsertNodesReply, error) {
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext(config)

	table := &ultipa.NodeTable{}

	table.Schemas = []*ultipa.Schema{
		{
			SchemaName: schema.Name,
			Properties: []*ultipa.Property{},
		},
	}

	for _, prop := range schema.Properties {

		if prop.IsIDType() {
			continue
		}

		table.Schemas[0].Properties = append(table.Schemas[0].Properties, &ultipa.Property{
			PropertyName: prop.Name,
			PropertyType: prop.Type,
		})
	}

	for _, row := range rows {

		newnode := &ultipa.NodeRow{
			Id:         row.ID,
			SchemaName: schema.Name,
		}

		for _, prop := range schema.Properties {

			if prop.IsIDType() {
				continue
			}

			bs, err := row.GetBytes(prop.Name)

			if err != nil {
				log.Fatal("Get row bytes value failed", prop.Name)
			}

			newnode.Values = append(newnode.Values, bs)
		}

		table.NodeRows = append(table.NodeRows, newnode)
	}

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: ultipa.InsertType_OVERWRITE,
		Silent:     true,
	})

	return resp, err
}
