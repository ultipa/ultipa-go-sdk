package api

import (
	"log"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.NodeTable, config *configuration.RequestConfig) (*ultipa.InsertNodesReply, error) {

	config.UseMaster = true
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

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	config.UseMaster = true
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

		if prop.IsIDType() || prop.IsIgnore() {
			continue
		}

		table.Schemas[0].Properties = append(table.Schemas[0].Properties, &ultipa.Property{
			PropertyName: prop.Name,
			PropertyType: prop.Type,
		})
	}

	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}

	for _, row := range rows {

		if row == nil {
			continue
		}

		wg.Add(1)
		go func(row *structs.Node) {
			defer wg.Done()
			newnode := &ultipa.NodeRow{
				Id:         row.ID,
				Uuid:       row.UUID,
				SchemaName: schema.Name,
			}

			for _, prop := range schema.Properties {

				if prop.IsIDType() || prop.IsIgnore() {
					continue
				}

				bs, err := row.GetBytes(prop.Name)

				if err != nil {
					log.Fatal("Get row bytes value failed  ", prop.Name, " ", err)
				}

				newnode.Values = append(newnode.Values, bs)
			}

			mtx.Lock()
			table.NodeRows = append(table.NodeRows, newnode)
			mtx.Unlock()
		}(row)

	}

	wg.Wait()

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: config.InsertType,
		Silent:     true,
	})

	return resp, err
}
