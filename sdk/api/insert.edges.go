package api

import (
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) InsertEdgesBatch(table *ultipa.EdgeTable, config *configuration.RequestConfig) (*ultipa.InsertEdgesReply, error) {
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext(config)

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName: conf.CurrentGraph,
		EdgeTable: table,
		Silent:    true,
	})

	return resp, err
}

func (api *UltipaAPI) InsertEdgesBatchBySchema(schema *structs.Schema, rows []*structs.Edge, config *configuration.RequestConfig) (*ultipa.InsertEdgesReply, error) {
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext(config)

	table := &ultipa.EdgeTable{}

	table.Headers = []*ultipa.SchemaHeader{
		{
			SchemaName: schema.Name,
			Headers:    []*ultipa.Header{},
		},
	}

	for _, prop := range schema.Properties {

		if prop.IsIDType() {
			continue
		}

		table.Headers[0].Headers = append(table.Headers[0].Headers, &ultipa.Header{
			PropertyName: prop.Name,
			PropertyType: prop.Type,
		})
	}

	for _, row := range rows {

		newnode := &ultipa.EdgeRow{
			FromId: row.From,
			ToId: row.To,
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

		table.EdgeRows = append(table.EdgeRows, newnode)
	}

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName: conf.CurrentGraph,
		EdgeTable: table,
		InsertType: ultipa.InsertType_OVERWRITE,
		Silent:    true,
	})

	return resp, err
}
