package api

import (
	"context"
	"errors"
	"fmt"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) InsertEdgesBatch(table *ultipa.EdgeTable, config *configuration.RequestConfig) (*ultipa.InsertEdgesReply, error) {

	config.UseMaster = true
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config)

	defer cancel()

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName: conf.CurrentGraph,
		EdgeTable: table,
		Silent:    true,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return resp, err
}

func (api *UltipaAPI) InsertEdgesBatchBySchema(schema *structs.Schema, rows []*structs.Edge, config *configuration.RequestConfig) (*ultipa.InsertEdgesReply, error) {

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

	err = setPropertiesToEdgeRow(schema, wg, edgeRows, rows)

	wg.Wait()
	if err != nil {
		return nil, err
	}
	table.EdgeRows = edgeRows
	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName:            conf.CurrentGraph,
		EdgeTable:            table,
		InsertType:           config.InsertType,
		CreateNodeIfNotExist: config.CreateNodeIfNotExist,
		Silent:               true,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return resp, err
}

func setPropertiesToEdgeRow(schema *structs.Schema, wg sync.WaitGroup, edgeRows []*ultipa.EdgeRow, rows []*structs.Edge) (err error) {
	ctx, cancel := context.WithCancel(context.Background())

	for index, row := range rows {
		if row == nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("edge row [%d] error: node row is nil.", index))
			}
			return err
		}

		properties := schema.Properties
		err = CheckEdgeRows(row, properties, index)
		if err != nil {
			return err
		}

		wg.Add(1)

		go func(index int, row *structs.Edge) {
			defer wg.Done()

			newEdge := &ultipa.EdgeRow{
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

				if !row.Values.Has(prop.Name) {
					cancel()
					err = errors.New(fmt.Sprintf("edge row [%d] error: values doesn't contain property [%s]", index, prop.Name))
				}

				bs, err := row.GetBytes(prop.Name)

				if err != nil {
					printers.PrintError("Get row bytes value failed " + prop.Name + " " + err.Error())
					err = errors.New(fmt.Sprintf("edge row [%d] error: failed to serialize value of property %s,value=%v", index, prop.Name, row.Values.Get(prop.Name)))
					return
				}

				newEdge.Values = append(newEdge.Values, bs)
			}
			edgeRows[index] = newEdge
		}(index, row)
		select {
		case <-ctx.Done():
			return err
		default:
		}
	}
	return err
}

//InsertNodesBatchAuto Nodes interface values should be string
func (api *UltipaAPI) InsertEdgesBatchAuto(edges []*structs.Edge, config *configuration.RequestConfig) error {

	// collect schema and nodes

	schemas, err := api.ListSchema(ultipa.DBType_DBEDGE, config)

	if err != nil {
		return err
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
				continue
			}
		}

		batch := batches[edge.Schema]
		// add edges
		batch.Edges = append(batch.Edges, edge)
	}

	for _, batch := range batches {

		structs.ConvertStringEdges(batch.Schema, batch.Edges)

		_, err := api.InsertEdgesBatchBySchema(batch.Schema, batch.Edges, config)

		if err != nil {
			return err
		}
	}

	return nil
}
