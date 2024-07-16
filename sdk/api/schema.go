package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ShowSchema(requestConfig *configuration.RequestConfig) ([]*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	resp, err = api.Uql(fmt.Sprintf(`show().schema()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		return nil, err
	}
	edgesSchemas, err := resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		return nil, err
	}

	schemas = append(schemas, edgesSchemas...)

	if len(schemas) == 0 {
		return nil, fmt.Errorf("no data return")
	}

	return schemas, err
}

func (api *UltipaAPI) ShowNodeSchema(requestConfig *configuration.RequestConfig) ([]*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	resp, err = api.Uql(fmt.Sprintf(`show().node_schema()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, fmt.Errorf("no data return")
	}

	return schemas, err
}

func (api *UltipaAPI) ShowEdgeSchema(requestConfig *configuration.RequestConfig) ([]*structs.Schema, error) {
	var resp *http.UQLResponse
	var err error
	var schemas []*structs.Schema

	resp, err = api.Uql(fmt.Sprintf(`show().edge_schema()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	schemas, err = resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, fmt.Errorf("no data return")
	}

	return schemas, err
}

func (api *UltipaAPI) GetSchema(schemaName string, dbType ultipa.DBType, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	err := CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	if dbType == ultipa.DBType_DBNODE {
		return api.GetNodeSchema(schemaName, requestConfig)
	} else if dbType == ultipa.DBType_DBEDGE {
		return api.GetEdgeSchema(schemaName, requestConfig)
	}

	return nil, errors.New("GetSchema() error db_type")

}

func (api *UltipaAPI) GetNodeSchema(schemaName string, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	err := CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	var resp *http.UQLResponse
	var schemas []*structs.Schema
	escapedSchemaName := schemaName
	if utils.IsNeedToEscapeName(schemaName) {
		escapedSchemaName = fmt.Sprintf("`%v`", schemaName)
	}
	resp, err = api.Uql(fmt.Sprintf(`show().node_schema(@%v)`, escapedSchemaName), requestConfig)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	schemas, err = resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas[0], err
}

func (api *UltipaAPI) GetEdgeSchema(schemaName string, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	err := CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	var resp *http.UQLResponse
	var schemas []*structs.Schema
	escapedSchemaName := schemaName
	if utils.IsNeedToEscapeName(schemaName) {
		escapedSchemaName = fmt.Sprintf("`%v`", schemaName)
	}
	resp, err = api.Uql(fmt.Sprintf(`show().edge_schema(@%v)`, escapedSchemaName), requestConfig)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	schemas, err = resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, err
	}

	return schemas[0], err
}

func (api *UltipaAPI) CreateSchema(schema *structs.Schema, isCreateProperties bool, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	err := CheckName(schema.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schema.Name))
	}
	var resp *http.UQLResponse

	schemaName := schema.Name
	if utils.IsNeedToEscapeName(schemaName) {
		schemaName = fmt.Sprintf("`%v`", schemaName)
	} else {
		schemaName = fmt.Sprintf(`"%v"`, schemaName)
	}
	api.Logger.Log("Creating Schema : @" + schema.Name)

	if schema.DBType == ultipa.DBType_DBNODE {
		uql := fmt.Sprintf(`create().node_schema(%v,"%v")`, schemaName, schema.Desc)

		resp, err = api.Uql(uql, requestConfig)
		if err != nil {
			return nil, err
		}
		if !resp.Status.IsSuccess() {
			return nil, errors.New(resp.Status.Message)
		}

	} else if schema.DBType == ultipa.DBType_DBEDGE {
		uql := fmt.Sprintf(`create().edge_schema(%v,"%v")`, schemaName, schema.Desc)
		resp, err = api.Uql(uql, requestConfig)
		if err != nil {
			return nil, err
		}
		if !resp.Status.IsSuccess() {
			return nil, errors.New(resp.Status.Message)
		}

	} else {

		return nil, errors.New("GetSchema() error db_type")

	}

	api.Logger.Log("Created Schema : @" + schema.Name)
	// create property of schemas
	if isCreateProperties {

		for _, prop := range schema.Properties {

			if prop.IsIDType() || prop.IsIgnore() {
				continue
			}

			resp, err := api.CreateProperty(schema.DBType, schema.Name, prop, requestConfig)

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

func (api *UltipaAPI) CreateSchemaIfNotExist(schema *structs.Schema, requestConfig *configuration.RequestConfig) (exist bool, err error) {
	err = CheckName(schema.Name)
	if err != nil {
		return false, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schema.Name))
	}
	exist = true
	s, _ := api.GetSchema(schema.Name, schema.DBType, requestConfig)

	if s == nil {
		_, err = api.CreateSchema(schema, true, requestConfig)
		exist = false
	}

	return exist, err

}

func (api *UltipaAPI) DropSchema(schema *structs.Schema, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := ""
	switch schema.DBType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_schema(@%s)`, schema.Name)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_schema(@%s)`, schema.Name)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) AlterSchema(schema, newSchema *structs.Schema, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	parms := ""

	switch schema.DBType {
	case ultipa.DBType_DBNODE:
		parms = "node_schema"
	case ultipa.DBType_DBEDGE:
		parms = "edge_schema"
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	uql := fmt.Sprintf(`alter().%s("@%s").set({name: "%s", description: "%s"})`, parms, schema.Name, newSchema.Name, newSchema.Desc)

	// Only modify the description of the schema
	if newSchema.Name == "" {
		uql = fmt.Sprintf(`alter().%s("@%s").set({description: "%s"})`, parms, schema.Name, newSchema.Desc)
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
