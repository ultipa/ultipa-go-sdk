package api

import (
	"errors"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.NodeTable, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

	config.UseMaster = true
	client, conf, err := api.GetClient(config.RequestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config.RequestConfig)
	defer cancel()

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName: conf.CurrentGraph,
		NodeTable: table,
		InsertType: config.InsertType,
		Silent:    config.Silent,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return http.NewNodesInsertResponse(resp)
}

func (api *UltipaAPI) InsertNodesBatchBySchema(schema *structs.Schema, rows []*structs.Node, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

	if config == nil {
		config = &configuration.InsertRequestConfig{}
	}

	config.UseMaster = true
	client, conf, err := api.GetClient(config.RequestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config.RequestConfig)
	defer cancel()

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
	nodeRows := make([]*ultipa.NodeRow, len(rows))
	for index, row := range rows {

		wg.Add(1)
		go func(index int, row *structs.Node) {
			defer wg.Done()

			if row.Get("_id") != "" {

			}

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
					 printers.PrintError("Get row bytes value failed  " + prop.Name + " " + err.Error())
					 return
				}

				newnode.Values = append(newnode.Values, bs)
			}
			nodeRows[index] = newnode
		}(index,row)
	}

	wg.Wait()
	table.NodeRows =nodeRows
	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: config.InsertType,
		Silent:     config.Silent,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return http.NewNodesInsertResponse(resp)
}

type Batch struct {
	Nodes []*structs.Node
	Edges []*structs.Edge
	Schema *structs.Schema
}


//InsertNodesBatchAuto Nodes interface values should be string
func(api *UltipaAPI)  InsertNodesBatchAuto(nodes []*structs.Node, config *configuration.InsertRequestConfig) (resps []*http.InsertResponse ,err error) {

	// collect schema and nodes

	schemas, err := api.ListSchema(ultipa.DBType_DBNODE, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	batches := map[string]*Batch{}

	for _, node := range nodes {

		// init schema
		if batches[node.Schema] == nil {

			batches[node.Schema] = &Batch{}

			s := utils.Find(schemas, func(i int) bool {
				return schemas[i].Name == node.Schema
			})

			if schema, ok := s.(*structs.Schema); ok {
				batches[node.Schema].Schema = schema
			} else {
				// schema not exit
				return nil, errors.New("Schema not found : " + node.Schema)
			}
		}

		batch := batches[node.Schema]
		// add nodes
		//node.UpdateByValueID()
		batch.Nodes = append(batch.Nodes, node)
	}

	for _, batch := range batches {

		structs.ConvertStringNodes(batch.Schema, batch.Nodes)

		resp , err := api.InsertNodesBatchBySchema(batch.Schema, batch.Nodes, config)

		if err != nil {
			return nil, err
		}

		resps = append(resps, resp)
	}

	return resps, nil
}


