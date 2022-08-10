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

	if err != nil {
		return nil, err
	}

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

	err, nodeRows := setPropertiesToNodeRow(schema, wg, rows)

	wg.Wait()
	if err != nil {
		return nil, err
	}
	table.NodeRows = nodeRows
	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: config.InsertType,
		Silent:     true,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return resp, err
}

func setPropertiesToNodeRow(schema *structs.Schema, wg sync.WaitGroup, rows []*structs.Node) (error, []*ultipa.NodeRow) {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	nodeRows := make([]*ultipa.NodeRow, len(rows))

	for index, row := range rows {
		if row == nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("node row [%d] error: node row is nil.", index))
			}
			return err, nodeRows
		}
		values := row.GetValues()
		properties := schema.Properties
		err = CheckValuesAndProperties(properties, values, index)
		if err != nil {
			return err, nodeRows
		}

		wg.Add(1)
		go func(index int, row *structs.Node) {
			defer wg.Done()
			newNode := &ultipa.NodeRow{
				Id:         row.ID,
				Uuid:       row.UUID,
				SchemaName: schema.Name,
			}

			for _, prop := range properties {

				if prop.IsIDType() || prop.IsIgnore() {
					continue
				}

				if !row.Values.Has(prop.Name) {
					cancel()
					err = errors.New(fmt.Sprintf("node row [%d] error: values doesn't contain property [%s]", index, prop.Name))
				}

				bs, err := row.GetBytes(prop.Name)

				if err != nil {
					printers.PrintError("Get row bytes value failed  " + prop.Name + " " + err.Error())
					err = errors.New(fmt.Sprintf("node row [%d] error: failed to serialize value of property %s,value=%v", index, prop.Name, row.Values.Get(prop.Name)))
					return
				}

				newNode.Values = append(newNode.Values, bs)
			}
			nodeRows[index] = newNode
		}(index, row)
		select {
		case <-ctx.Done():
			return err, nodeRows
		default:
		}
	}
	return err, nodeRows
}

type Batch struct {
	Nodes  []*structs.Node
	Edges  []*structs.Edge
	Schema *structs.Schema
}

//InsertNodesBatchAuto Nodes interface values should be string
func (api *UltipaAPI) InsertNodesBatchAuto(nodes []*structs.Node, config *configuration.RequestConfig) error {

	// collect schema and nodes

	schemas, err := api.ListSchema(ultipa.DBType_DBNODE, config)

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

		_, err := api.InsertNodesBatchBySchema(batch.Schema, batch.Nodes, config)

		if err != nil {
			return err
		}
	}

	return nil
}
