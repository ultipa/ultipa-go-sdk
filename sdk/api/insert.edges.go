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

func (api *UltipaAPI) InsertEdgesBatch(table *ultipa.EdgeTable, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

	config.UseMaster = true
	client, conf, err := api.GetClient(config.RequestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config.RequestConfig)

	defer cancel()

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName: conf.CurrentGraph,
		EdgeTable: table,
		CreateNodeIfNotExist: config.CreateNodeIfNotExist,
		InsertType: config.InsertType,
		Silent:    config.Silent,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return http.NewEdgesInsertResponse(resp)
}

func (api *UltipaAPI) 	InsertEdgesBatchBySchema(schema *structs.Schema, rows []*structs.Edge, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

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

	table := &ultipa.EdgeTable{}

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
	edgeRows := make([]*ultipa.EdgeRow, len(rows))
	for index, row := range rows {

		wg.Add(1)

		go func(index int, row *structs.Edge) {
			defer wg.Done()

			newnode := &ultipa.EdgeRow{
				FromId:     row.From,
				FromUuid:   row.FromUUID,
				ToId:       row.To,
				ToUuid:     row.ToUUID,
				SchemaName: schema.Name,
				Uuid:       row.UUID,
			}

			for _, prop := range schema.Properties {

				if prop.IsIDType() || prop.IsIgnore() {
					continue
				}

				bs, err := row.GetBytes(prop.Name)

				if err != nil {
					printers.PrintError("Get row bytes value failed " + prop.Name + " " + err.Error())
				}

				newnode.Values = append(newnode.Values, bs)
			}
			edgeRows[index] = newnode
		}(index, row)
	}
	wg.Wait()
	table.EdgeRows =edgeRows

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName:            conf.CurrentGraph,
		EdgeTable:            table,
		InsertType:           config.InsertType,
		CreateNodeIfNotExist: config.CreateNodeIfNotExist,
		Silent:               config.Silent,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return http.NewEdgesInsertResponse(resp)
}

//InsertEdgesBatchAuto Nodes interface values should be string
func (api *UltipaAPI) InsertEdgesBatchAuto(edges []*structs.Edge, config *configuration.InsertRequestConfig) (resps []*http.InsertResponse, err error) {

	// collect schema and nodes

	schemas, err := api.ListSchema(ultipa.DBType_DBEDGE, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	batches := map[string]*Batch{}

	for _, edge := range edges {

		// init schema
		if batches[edge.Schema] == nil {

			batches[edge.Schema] = &Batch{}

			s := utils.Find(schemas, func(i int) bool {
				return schemas[i].Name == edge.Schema
			})

			if schema, ok := s.(*structs.Schema); ok {
				batches[edge.Schema].Schema = schema
			} else {
				// schema not exit
				return nil, errors.New("Edge Schema not found : " + edge.Schema)
			}
		}

		batch := batches[edge.Schema]
		// add edges
		batch.Edges = append(batch.Edges, edge)
	}

	for _, batch := range batches {

		structs.ConvertStringEdges(batch.Schema, batch.Edges)

		resp, err := api.InsertEdgesBatchBySchema(batch.Schema, batch.Edges, config)

		if err != nil {
			return nil,err
		}

		resps = append(resps, resp)

	}

	return resps,nil
}
