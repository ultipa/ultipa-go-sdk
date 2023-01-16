package api

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.EntityTable, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

	config.UseMaster = true
	client, conf, err := api.GetClient(config.RequestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config.RequestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: config.InsertType,
		//TODO 暂时先设置为false，批量插入不返回ids，后续调整再定
		//Silent:     config.Silent,
		Silent: false,
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

	if config.RequestConfig == nil {
		config.RequestConfig = &configuration.RequestConfig{}
	}

	config.UseMaster = true
	client, conf, err := api.GetClient(config.RequestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config.RequestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	table := &ultipa.EntityTable{}

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

	err, nodeRows := setPropertiesToNodeRow(schema, rows)

	if err != nil {
		return nil, err
	}
	table.EntityRows = nodeRows
	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName:  conf.CurrentGraph,
		NodeTable:  table,
		InsertType: config.InsertType,
		//TODO 暂时先设置为false，批量插入不返回ids，后续调整再定
		//Silent:     config.Silent,
		Silent: false,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return http.NewNodesInsertResponse(resp)
}

func setPropertiesToNodeRow(schema *structs.Schema, rows []*structs.Node) (error, []*ultipa.EntityRow) {
	wg := sync.WaitGroup{}
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	nodeRows := make([]*ultipa.EntityRow, len(rows))

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
			newNode := &ultipa.EntityRow{
				Id:         row.ID,
				Uuid:       row.UUID,
				SchemaName: schema.Name,
			}

			for _, prop := range properties {

				if prop.IsIDType() || prop.IsIgnore() {
					continue
				}

				bs, err := row.GetBytesSafe(prop.Name, prop.Type)

				if err != nil {
					printers.PrintError("Get row bytes value failed  " + prop.Name + " " + err.Error())
					err = errors.New(fmt.Sprintf("node row [%d] error: failed to serialize value of property %s,value=%v", index, prop.Name, row.Values.Get(prop.Name)))
					cancel()
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
	wg.Wait()
	return err, nodeRows
}

type Batch struct {
	Nodes  []*ultipa.EntityRow
	Edges  []*ultipa.EntityRow
	Schema *structs.Schema
}

//InsertNodesBatchAuto Nodes interface values should be string
func (api *UltipaAPI) InsertNodesBatchAuto(nodes []*structs.Node, config *configuration.InsertRequestConfig) (*http.InsertBatchAutoResponse, error) {

	resps := &http.InsertBatchAutoResponse{
		Resps:     map[string]*http.InsertResponse{},
		ErrorItem: map[int]int{},
		Statistic: &http.Statistic{},
	}

	// collect schema and nodes
	m := map[string]map[int]int{}
	schemas, err := api.ListSchema(ultipa.DBType_DBNODE, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	batches := map[string]*Batch{}

	for index, node := range nodes {
		m[node.Schema] = map[int]int{}
		var rows []*ultipa.EntityRow
		// init schema
		if batches[node.Schema] == nil {

			batches[node.Schema] = &Batch{}

			s := utils.Find(schemas, func(i int) bool {
				return schemas[i].Name == node.Schema
			})

			if schema, ok := s.(*structs.Schema); ok {
				batches[node.Schema].Schema = schema

				err, rows = setPropertiesToNodeRow(schema, []*structs.Node{node})
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Data verification failed, index: [%s], caused by:%v", strconv.Itoa(index), err))
				}
			} else {
				// schema not exit
				return nil, errors.New("Schema not found : " + node.Schema)
			}
		}

		batch := batches[node.Schema]
		// add nodes
		//node.UpdateByValueID()
		if len(rows) != 0 {
			batch.Nodes = append(batch.Nodes, rows[0])
			m[node.Schema][len(batch.Nodes)-1] = index
		}
		//batch.Nodes = append(batch.Nodes, node)
	}

	for _, batch := range batches {
		batchSchema := batch.Schema

		if config == nil {
			config = &configuration.InsertRequestConfig{}
		}

		if config.RequestConfig == nil {
			config.RequestConfig = &configuration.RequestConfig{}
		}

		config.UseMaster = true
		client, conf, err := api.GetClient(config.RequestConfig)

		if err != nil {
			return nil, err
		}

		ctx, cancel, err := api.Pool.NewContext(config.RequestConfig)
		if err != nil {
			return nil, err
		}
		defer cancel()

		table := &ultipa.EntityTable{}

		table.Schemas = []*ultipa.Schema{
			{
				SchemaName: batchSchema.Name,
				Properties: []*ultipa.Property{},
			},
		}

		for _, prop := range batchSchema.Properties {

			if prop.IsIDType() || prop.IsIgnore() {
				continue
			}

			table.Schemas[0].Properties = append(table.Schemas[0].Properties, &ultipa.Property{
				PropertyName: prop.Name,
				PropertyType: prop.Type,
			})
		}

		if err != nil {
			return nil, err
		}
		table.EntityRows = batch.Nodes
		resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
			GraphName:  conf.CurrentGraph,
			NodeTable:  table,
			InsertType: config.InsertType,
			//TODO 暂时先设置为false，批量插入不返回ids，后续调整再定
			//Silent:     config.Silent,
			Silent: false,
		})

		if err != nil {
			return nil, err
		}

		if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
			if resps.ErrorCode == "" {
				resps.ErrorCode = ultipa.ErrorCode_name[int32(resp.Status.ErrorCode)]
			}
			resps.Msg += batchSchema.Name + ":" + resp.Status.Msg + "\r\n"
		}

		response, err := http.NewNodesInsertResponse(resp)
		resps.Resps[batchSchema.Name] = response

		for k, v := range response.Data.ErrorItem {
			m3 := m[batchSchema.Name]
			vl := m3[k]
			resps.ErrorItem[vl] = v
		}
		resps.Statistic.TotalCost += response.Statistic.TotalCost
		resps.Statistic.EngineCost += response.Statistic.EngineCost

	}

	return resps, nil
}
