package api

import (
	"errors"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.NodeTable, config *configuration.RequestConfig) (*ultipa.InsertNodesReply, error) {

	config.UseMaster = true
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config)
	defer cancel()

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName: conf.CurrentGraph,
		NodeTable: table,
		Silent:    true,
	})

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

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

	ctx, cancel := api.Pool.NewContext(config)
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
					 printers.PrintError("Get row bytes value failed  " + prop.Name + " " + err.Error())
					 return
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

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return resp, err
}

type Batch struct {
	Nodes []*structs.Node
	Edges []*structs.Edge
	Schema *structs.Schema
}


//InsertNodesBatchAuto Nodes interface values should be string
func(api *UltipaAPI)  InsertNodesBatchAuto(nodes []*structs.Node, config *configuration.RequestConfig) error {

	// collect schema and nodes

	schemas, err := api.ListSchema(ultipa.DBType_DBNODE,config)

	if err != nil {
		return err
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
				continue
			}
		}

		batch := batches[node.Schema]
		// add nodes
		batch.Nodes = append(batch.Nodes, node)
	}

	for _, batch := range batches {

		structs.ConvertStringNodes(batch.Schema, batch.Nodes)

		_ , err := api.InsertNodesBatchBySchema(batch.Schema, batch.Nodes, config)

		if err != nil {
			return err
		}
	}

	return nil
}


