package api

import (
	"context"
	"errors"
	"fmt"
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

	ctx, cancel, err := api.Pool.NewContext(config.RequestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName:            conf.CurrentGraph,
		EdgeTable:            table,
		CreateNodeIfNotExist: config.CreateNodeIfNotExist,
		InsertType:           config.InsertType,
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

	return http.NewEdgesInsertResponse(resp)
}

func (api *UltipaAPI) InsertEdgesBatchBySchema(schema *structs.Schema, rows []*structs.Edge, config *configuration.InsertRequestConfig) (*http.InsertResponse, error) {

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

	err, edgeRows := setPropertiesToEdgeRow(schema, rows)

	if err != nil {
		return nil, err
	}
	table.EdgeRows = edgeRows
	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName:            conf.CurrentGraph,
		EdgeTable:            table,
		InsertType:           config.InsertType,
		CreateNodeIfNotExist: config.CreateNodeIfNotExist,
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

	return http.NewEdgesInsertResponse(resp)
}

func setPropertiesToEdgeRow(schema *structs.Schema, rows []*structs.Edge) (error, []*ultipa.EdgeRow) {
	wg := sync.WaitGroup{}
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	edgeRows := make([]*ultipa.EdgeRow, len(rows))

	for index, row := range rows {
		err = checkEdgeProperties(schema, row, index)
		if err != nil {
			return err, edgeRows
		}

		wg.Add(1)

		go func(index int, row *structs.Edge) {
			defer wg.Done()
			var newEdge *ultipa.EdgeRow
			newEdge, err = doConvertSdkEdgeRowToUltipaEdgeRow(schema, row, index)
			if err != nil {
				cancel()
				return
			}
			edgeRows[index] = newEdge
		}(index, row)
		select {
		case <-ctx.Done():
			return err, edgeRows
		default:
		}
	}
	wg.Wait()
	return err, edgeRows
}

func checkEdgeProperties(schema *structs.Schema, row *structs.Edge, index int) error {
	if row == nil {
		return errors.New(fmt.Sprintf("node row [%d] error: node row is nil.", index))
	}
	err := CheckEdgeRows(row, schema.Properties, index)
	if err != nil {
		return err
	}
	return nil
}

func convertSdkEdgeRowToUltipaEdgeRow(schema *structs.Schema, row *structs.Edge, index int) (*ultipa.EdgeRow, error) {
	err := checkEdgeProperties(schema, row, index)
	if err != nil {
		return nil, err
	}
	return doConvertSdkEdgeRowToUltipaEdgeRow(schema, row, index)
}

func doConvertSdkEdgeRowToUltipaEdgeRow(schema *structs.Schema, row *structs.Edge, index int) (*ultipa.EdgeRow, error) {
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

		if !row.Values.Contain(prop.Name) {
			return nil, errors.New(fmt.Sprintf("edge row [%d] error: values doesn't contain property [%s]", index, prop.Name))
		}

		bs, err := row.GetBytesSafe(prop.Name, prop.Type)

		if err != nil {
			printers.PrintError("Get row bytes value failed " + prop.Name + " " + err.Error())
			err = errors.New(fmt.Sprintf("edge row [%d] error: failed to serialize value of property %s,value=%v", index, prop.Name, row.Values.Get(prop.Name)))
			return nil, err
		}

		newEdge.Values = append(newEdge.Values, bs)
	}
	return newEdge, nil
}

//InsertEdgesBatchAuto Nodes interface values should be string
func (api *UltipaAPI) InsertEdgesBatchAuto(edges []*structs.Edge, config *configuration.InsertRequestConfig) (*http.InsertBatchAutoResponse, error) {

	resps := &http.InsertBatchAutoResponse{
		Resps:     map[string]*http.InsertResponse{},
		ErrorItem: map[int]int{},
		Statistic: &http.Statistic{},
	}

	// collect schema and nodes
	m := map[string]map[int]int{}
	schemas, err := api.ListSchema(ultipa.DBType_DBEDGE, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	batches := map[string]*Batch{}

	for index, edge := range edges {

		m[edge.Schema] = map[int]int{}
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
		row, err := convertSdkEdgeRowToUltipaEdgeRow(batch.Schema, edge, index)
		if err != nil {
			return nil, err
		}

		if row != nil {
			batch.Edges = append(batch.Edges, row)
			m[edge.Schema][len(batch.Edges)-1] = index
		}
		//batch.Edges = append(batch.Edges, edge)
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

		table := &ultipa.EdgeTable{}

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
		table.EdgeRows = batch.Edges
		resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
			GraphName:            conf.CurrentGraph,
			EdgeTable:            table,
			InsertType:           config.InsertType,
			CreateNodeIfNotExist: config.CreateNodeIfNotExist,
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

		response, err := http.NewEdgesInsertResponse(resp)
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
