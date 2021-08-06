package api

import (
	"errors"
	"fmt"
	"strconv"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ListNodeSchema(config *configuration.RequestConfig) (*http.ResponseNodeSchemas, error) {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listNodeSchema)
	res, err := api.UQL(uql.ToString(), config)
	if err != nil {
		return nil, err
	}
	table, err := res.GetSingleTable()
	if err != nil {
		return nil, err
	}
	var schemas []*http.ResponseSchema
	if !res.Status.IsSuccess() {
		return &http.ResponseNodeSchemas{
			Status:  res.Status,
			Schemas: schemas,
		}, nil
	}
	values := table.ToKV()
	for _, v := range values {
		totalNodes, _ := strconv.ParseInt(v.Get("totalNodes").(string), 10, 64)
		totalEdges, _ := strconv.ParseInt(v.Get("totalEdges").(string), 10, 64)

		schemas = append(schemas, &http.ResponseSchema{
			Name:        v.Get("name").(string),
			Description: v.Get("description").(string),
			Properties:  nil,
			TotalNodes:  totalNodes,
			TotalEdges:  totalEdges,
		})
	}
	return &http.ResponseNodeSchemas{
		Status:  res.Status,
		Schemas: schemas,
	}, nil
}

func (api *UltipaAPI) GetSchema(schemaName string, DBType ultipa.DBType) (*structs.Schema, error) {

	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	if DBType == ultipa.DBType_DBNODE {

		resp, err = api.UQL(fmt.Sprintf(`show().node_schema("%v")`, schemaName), nil)
		if err != nil {
			return nil, err
		}

	} else if DBType == ultipa.DBType_DBEDGE {
		resp, err = api.UQL(fmt.Sprintf(`show().edge_schema(%v)`, schemaName), nil)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("GetSchema() error db_type")
	}

	schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()

	if err != nil {
		return nil, err
	}

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas[0], err
}

func (api *UltipaAPI) CreateSchema(schema *structs.Schema, isCreateProperties bool) (*http.UQLResponse, error) {

	var resp *http.UQLResponse
	var err error

	if schema.DBType == ultipa.DBType_DBNODE {
		resp, err = api.UQL(fmt.Sprintf(`create().node_schema("%v","%v")`, schema.Name, schema.Desc), nil)
		if err != nil {
			return nil, err
		}

	} else if schema.DBType == ultipa.DBType_DBEDGE {
		resp, err = api.UQL(fmt.Sprintf(`create().edge_schema("%v","%v")`, schema.Name, schema.Desc), nil)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("GetSchema() error db_type")
	}

	// todo create property here!
	if isCreateProperties {

	}

	return resp, err
}
