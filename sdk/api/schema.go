package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		//totalEdges, _ := strconv.ParseInt(v.Get("totalEdges").(string), 10, 64)

		schemas = append(schemas, &http.ResponseSchema{
			Name:        v.Get("name").(string),
			Description: v.Get("description").(string),
			Properties:  nil,
			TotalNodes:  totalNodes,
			//TotalEdges:  totalEdges,
		})
	}
	return &http.ResponseNodeSchemas{
		Status:  res.Status,
		Schemas: schemas,
	}, nil
}

func (api *UltipaAPI) ListSchema(DBType ultipa.DBType, config *configuration.RequestConfig) ([]*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	if DBType == ultipa.DBType_DBNODE {
		resp, err = api.UQL(fmt.Sprintf(`show().node_schema()`), config)
		if err != nil {
			return nil, err
		}

		schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()
	} else if DBType == ultipa.DBType_DBEDGE {
		resp, err = api.UQL(fmt.Sprintf(`show().edge_schema()`), config)
		if err != nil {
			return nil, err
		}

		schemas, err = resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()
	}

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas, err
}

func (api *UltipaAPI) GetSchema(schemaName string, DBType ultipa.DBType, config *configuration.RequestConfig) (*structs.Schema, error) {

	if DBType == ultipa.DBType_DBNODE {
		return api.GetNodeSchema(schemaName, config)
	} else if DBType == ultipa.DBType_DBEDGE {
		return api.GetEdgeSchema(schemaName, config)
	} else {
		return nil, errors.New("GetSchema() error db_type")
	}

	return nil, nil

}

func (api *UltipaAPI) GetNodeSchema(schemaName string, config *configuration.RequestConfig) (*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	resp, err = api.UQL(fmt.Sprintf(`show().node_schema(@%v)`, schemaName), config)
	if err != nil {
		return nil, err
	}

	schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas[0], err
}

func (api *UltipaAPI) GetEdgeSchema(schemaName string, config *configuration.RequestConfig) (*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	resp, err = api.UQL(fmt.Sprintf(`show().edge_schema(@%v)`, schemaName), config)
	if err != nil {
		return nil, err
	}

	schemas, err = resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas[0], err
}

func (api *UltipaAPI) CreateSchema(schema *structs.Schema, isCreateProperties bool, conf *configuration.RequestConfig) (*http.UQLResponse, error) {

	var resp *http.UQLResponse
	var err error

	schemaName := schema.Name
	if !utils.CheckCustomerNonIdName(schemaName) && !strings.HasPrefix(schemaName, "`") && !strings.HasSuffix(schemaName, "`") {
		schemaName = fmt.Sprintf("`%v`", schemaName)
	}

	isEscapedName := utils.CheckIsEscapedName(schemaName)

	api.Logger.Log("Creating Schema : @" + schema.Name)

	if schema.DBType == ultipa.DBType_DBNODE {
		uql := fmt.Sprintf(`create().node_schema("%v","%v")`, schemaName, schema.Desc)
		if isEscapedName {
			uql = fmt.Sprintf(`create().node_schema(%v,"%v")`, schemaName, schema.Desc)
		}
		resp, err = api.UQL(uql, conf)
		if err != nil {
			return nil, err
		}

	} else if schema.DBType == ultipa.DBType_DBEDGE {

		resp, err = api.UQL(fmt.Sprintf(`create().edge_schema("%v","%v")`, schemaName, schema.Desc), conf)
		if err != nil {
			return nil, err
		}

	} else {

		return nil, errors.New("GetSchema() error db_type")

	}

	api.Logger.Log("Created Schema : @" + schemaName)
	// create property of schemas
	if isCreateProperties {

		for _, prop := range schema.Properties {

			if prop.IsIDType() || prop.IsIgnore() {
				continue
			}

			resp, err := api.CreateProperty(schemaName, schema.DBType, prop, conf)

			if err != nil {
				return nil, err
			}

			if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
				return resp, nil
			}
		}

	}

	return resp, err
}

func (api *UltipaAPI) CreateSchemaIfNotExist(schema *structs.Schema, config *configuration.RequestConfig) (exist bool, err error) {

	exist = true
	s, _ := api.GetSchema(schema.Name, schema.DBType, config)

	if s == nil {
		_, err = api.CreateSchema(schema, true, config)
		exist = false
	}

	return exist, err

}
