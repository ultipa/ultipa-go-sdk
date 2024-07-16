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

// CreateProperty create property for schema, schemaName maybe escaped if schemaName contains some special characters.
func (api *UltipaAPI) CreateProperty(dbType ultipa.DBType, schemaName string, prop *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	err = CheckName(prop.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}
	escapedSchemaName := schemaName
	if utils.IsNeedToEscapeSchemaNameForProperty(schemaName) {
		escapedSchemaName = fmt.Sprintf("`%v`", schemaName)
	}

	return api.doCreateProperty(escapedSchemaName, dbType, prop, requestConfig)
}

func (api *UltipaAPI) doCreateProperty(schemaName string, dbType ultipa.DBType, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	api.Logger.Log("Creating Property : @" + schemaName + "." + prop.Name)
	switch dbType {
	case ultipa.DBType_DBNODE:
		resp, err = api.doCreateNodeProperty(schemaName, prop, config)
	case ultipa.DBType_DBEDGE:
		resp, err = api.doCreateEdgeProperty(schemaName, prop, config)
	default:
		return nil, errors.New("create property: unknown db type")
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return resp, errors.New(resp.Status.Message)
	}
	api.Logger.Log("Created Property : @" + schemaName + "." + prop.Name)
	return resp, err
}

func (api *UltipaAPI) CreatePropertyIfNotExist(dbType ultipa.DBType, schemaName string, prop *structs.Property, config *configuration.RequestConfig) (exist bool, err error) {
	err = CheckName(schemaName)
	if err != nil {
		return false, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(prop.Name)
	if err != nil {
		return false, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}
	property, err := api.GetProperty(dbType, schemaName, prop.Name, config)

	if err != nil {
		return false, err
	}

	if property == nil {
		_, err = api.CreateProperty(dbType, schemaName, prop, config)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (api *UltipaAPI) GetProperty(dbType ultipa.DBType, schemaName string, propertyName string, requestConfig *configuration.RequestConfig) (property *structs.Property, err error) {
	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}
	schema, err := api.GetSchema(schemaName, dbType, requestConfig)

	if err != nil {
		return nil, err
	}

	for _, prop := range schema.Properties {
		if prop.Name == propertyName {
			return prop, nil
		}
	}

	return nil, fmt.Errorf("property %s not found in schema %s", propertyName, schemaName)
}

func (api *UltipaAPI) ShowProperty(dbType ultipa.DBType, schemaName string, requestConfig *configuration.RequestConfig) (property []*structs.Property, err error) {
	var schemas []*structs.Schema

	// get all schema property
	if schemaName == "" {
		if dbType == ultipa.DBType_DBNODE {
			schemas, err = api.ShowNodeSchema(requestConfig)
		} else if dbType == ultipa.DBType_DBEDGE {
			schemas, err = api.ShowEdgeSchema(requestConfig)
		}

		if err != nil {
			return nil, err
		}

		for _, schema := range schemas {
			property = append(property, schema.Properties...)
		}
		return property, nil
	}

	// get one schema property
	schema, err := api.GetSchema(schemaName, dbType, requestConfig)

	if err != nil {
		return nil, err
	}

	return schema.Properties, nil
}

func (api *UltipaAPI) ShowNodeProperty(schemaName string, requestConfig *configuration.RequestConfig) (property []*structs.Property, err error) {
	var schemas []*structs.Schema

	// get all schema property
	if schemaName == "" {
		schemas, err = api.ShowNodeSchema(requestConfig)
		if err != nil {
			return nil, err
		}

		for _, schema := range schemas {
			property = append(property, schema.Properties...)
		}
		return property, nil
	}

	// get one schema property
	schema, err := api.GetSchema(schemaName, ultipa.DBType_DBNODE, requestConfig)

	if err != nil {
		return nil, err
	}

	return schema.Properties, nil
}

func (api *UltipaAPI) ShowEdgeProperty(schemaName string, requestConfig *configuration.RequestConfig) (property []*structs.Property, err error) {
	var schemas []*structs.Schema

	// get all schema property
	if schemaName == "" {
		schemas, err = api.ShowEdgeSchema(requestConfig)

		if err != nil {
			return nil, err
		}

		for _, schema := range schemas {
			property = append(property, schema.Properties...)
		}
		return property, nil
	}

	// get one schema property
	schema, err := api.GetSchema(schemaName, ultipa.DBType_DBEDGE, requestConfig)

	if err != nil {
		return nil, err
	}

	return schema.Properties, nil
}

func (api *UltipaAPI) GetNodeProperty(schemaName string, propertyName string, config *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(ultipa.DBType_DBNODE, schemaName, propertyName, config)
}

func (api *UltipaAPI) GetEdgeProperty(schemaName string, propertyName string, config *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(ultipa.DBType_DBEDGE, schemaName, propertyName, config)
}

func (api *UltipaAPI) CreateNodeProperty(schemaName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	if prop.Type == structs.PropertyType_IGNORE {
		return nil, err
	}
	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(prop.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}
	escapeSchemaName := schemaName
	if utils.IsNeedToEscapeSchemaNameForProperty(schemaName) {
		escapeSchemaName = fmt.Sprintf("`%v`", schemaName)
	}

	return api.doCreateNodeProperty(escapeSchemaName, prop, config)
}

func (api *UltipaAPI) doCreateNodeProperty(schemaName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	if prop.Type == structs.PropertyType_IGNORE {
		return nil, err
	}
	propertyTypeStr, err := prop.GetStringType()
	if err != nil {
		return nil, err
	}
	propName := prop.Name
	if utils.IsNeedToEscapeName(propName) {
		propName = fmt.Sprintf("`%v`", propName)
	} else {
		propName = fmt.Sprintf(`"%v"`, propName)
	}

	uql := fmt.Sprintf(`create().node_property(@%v,%s,"%v","%v")`, schemaName, propName, propertyTypeStr, prop.Desc)

	resp, err = api.Uql(uql, config)
	return resp, err
}

func (api *UltipaAPI) CreateEdgeProperty(schemaName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	if prop.Type == structs.PropertyType_IGNORE {
		return nil, err
	}

	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(prop.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}

	escapeSchemaName := schemaName
	if utils.IsNeedToEscapeSchemaNameForProperty(schemaName) {
		escapeSchemaName = fmt.Sprintf("`%v`", schemaName)
	}

	return api.doCreateEdgeProperty(escapeSchemaName, prop, config)
}

func (api *UltipaAPI) doCreateEdgeProperty(schemaName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	if prop.Type == structs.PropertyType_IGNORE {
		return nil, err
	}
	propertyTypeStr, err := prop.GetStringType()
	if err != nil {
		return nil, err
	}

	propName := prop.Name
	if utils.IsNeedToEscapeName(propName) {
		propName = fmt.Sprintf("`%s`", propName)
	} else {
		propName = fmt.Sprintf(`"%s"`, propName)
	}

	uql := fmt.Sprintf(`create().edge_property(@%v,%v,"%v","%v")`, schemaName, propName, propertyTypeStr, prop.Desc)
	resp, err = api.Uql(uql, config)
	return resp, err
}

func (api *UltipaAPI) AlterNodeProperty(property, newProperty *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.AlterProperty(ultipa.DBType_DBNODE, property, newProperty, config)
}

func (api *UltipaAPI) AlterEdgeProperty(property, newProperty *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.AlterProperty(ultipa.DBType_DBEDGE, property, newProperty, config)
}

// Usage: DropNodeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropNodeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.Uql(fmt.Sprintf(`drop().node_property(%v)`, propertyName), config)

	return resp, err
}

// Usage: DropEdgeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropEdgeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.Uql(fmt.Sprintf(`drop().edge_property(%v)`, propertyName), config)

	return resp, err
}

func (api *UltipaAPI) DropProperty(dbType ultipa.DBType, schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_property(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_property(@%v.%v)`, schemaName, propertyName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) AlterProperty(dbType ultipa.DBType, property, newProperty *structs.Property, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	params := ""

	switch dbType {
	case ultipa.DBType_DBNODE:
		params = "node_property"
	case ultipa.DBType_DBEDGE:
		params = "edge_property"
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	uql := fmt.Sprintf(`alter().%v("@%v.%v").set({name: "%v", description: "%v"})`, params, property.Schema, property.Name, newProperty.Name, newProperty.Desc)

	// Only modify the description of the property
	if newProperty.Name == "" {
		uql = fmt.Sprintf(`alter().%v("@%v.%v").set({description: "%v"})`, params, property.Schema, property.Name, newProperty.Desc)
	}

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
