package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowSchema(requestConfig *configuration.RequestConfig) (*structs.Schemas, error) {
	var resp *http.UQLResponse
	var err error
	var schemas = &structs.Schemas{}

	resp, err = api.Uql(fmt.Sprintf(`show().schema()`), requestConfig)
	if err != nil {
		return nil, err
	}

	nodeSchemas, err := resp.Alias(http.RESP_NODE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		return nil, err
	}
	edgesSchemas, err := resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()
	if err != nil {
		return nil, err
	}

	schemas.Schemas = append(nodeSchemas, edgesSchemas...)

	graphCount, err := resp.Alias(http.RESP_GRAPH_COUNT_KEY).AsGraphCount()
	if err != nil {
		return nil, err
	}
	for _, g := range graphCount {
		switch g.Type {
		case "total_nodes":
			schemas.TotalNodes = g.SP.Count
		case "total_edges":
			schemas.TotalEdges = g.SP.Count
		case "node", "edge":
			for _, schema := range schemas.Schemas {
				if schema.Name == g.Schema && schema.Type == g.Type {
					schema.SetTotalByGraphCount(g)
				}
			}
		}

	}

	if len(schemas.Schemas) == 0 {
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

	schemas, err = resp.Alias(http.RESP_EDGE_SCHEMA_KEY).AsSchemas()

	if len(schemas) == 0 {
		return nil, fmt.Errorf("no data return")
	}

	return schemas, err
}

func (api *UltipaAPI) GetSchema(schemaName string, dbType ultipa.DBType, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	if !(dbType == ultipa.DBType_DBNODE || dbType == ultipa.DBType_DBEDGE) {
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	schemas, err := api.ShowSchema(requestConfig)
	if err != nil {
		return nil, err
	}

	for _, schema := range schemas.Schemas {
		if schemaName == schema.Name && schema.DBType == dbType {
			return schema, nil
		}
	}

	return nil, nil
}

func (api *UltipaAPI) GetNodeSchema(schemaName string, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	return api.GetSchema(schemaName, ultipa.DBType_DBNODE, requestConfig)
}

func (api *UltipaAPI) GetEdgeSchema(schemaName string, requestConfig *configuration.RequestConfig) (*structs.Schema, error) {
	return api.GetSchema(schemaName, ultipa.DBType_DBEDGE, requestConfig)
}

func (api *UltipaAPI) CreateSchema(schema *structs.Schema, isCreateProperties bool, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	schemaName, err := CheckReplaceSchemaPropertyName(schema.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schema.Name))
	}

	var resp *http.UQLResponse

	api.Logger.Log("Creating Schema : @" + schemaName)

	uql := ""
	if schema.DBType == ultipa.DBType_DBNODE {
		uql = fmt.Sprintf(`create().node_schema(%v,"%v")`, schemaName, schema.Desc)
	} else if schema.DBType == ultipa.DBType_DBEDGE {
		uql = fmt.Sprintf(`create().edge_schema(%v,"%v")`, schemaName, schema.Desc)
	} else {
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, requestConfig)
	if err != nil {
		return nil, err
	}

	api.Logger.Log("Created Schema : @" + schemaName)
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
	s, err := api.GetSchema(schema.Name, schema.DBType, requestConfig)
	if err != nil {
		return false, fmt.Errorf("GetSchema error, %v", err)
	}

	exist = true
	if s == nil {
		_, err = api.CreateSchema(schema, true, requestConfig)
		exist = false
	}

	return exist, err
}

func (api *UltipaAPI) DropSchema(schema *structs.Schema, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	schemaName, err := CheckReplaceSchemaPropertyName(schema.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schema.Name))
	}

	uql := ""
	switch schema.DBType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_schema(@%s)`, schemaName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_schema(@%s)`, schemaName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) AlterSchema(schema, newSchema *structs.Schema, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	schemaName, err := CheckReplaceSchemaPropertyName(schema.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schema.Name))
	}

	err = CheckName(newSchema.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), newSchema.Name))
	}

	parms := ""
	switch schema.DBType {
	case ultipa.DBType_DBNODE:
		parms = "node_schema"
	case ultipa.DBType_DBEDGE:
		parms = "edge_schema"
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	uql := fmt.Sprintf(`alter().%s(@%s).set({name: "%s", description: "%s"})`, parms, schemaName, newSchema.Name, newSchema.Desc)

	// Only modify the description of the schema
	if newSchema.Name == "" {
		uql = fmt.Sprintf(`alter().%s(@%s).set({description: "%s"})`, parms, schemaName, newSchema.Desc)
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
