package api

import (
	"errors"
	"fmt"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

//CreateProperty create property for schema, schemaName maybe escaped if schemaName contains some special characters.
func (api *UltipaAPI) CreateProperty(schemaName string, dbType ultipa.DBType, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
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

	return api.doCreateProperty(escapedSchemaName, dbType, prop, conf)
}

func (api *UltipaAPI) doCreateProperty(schemaName string, dbType ultipa.DBType, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	api.Logger.Log("Creating Property : @" + schemaName + "." + prop.Name)
	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(prop.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}
	switch dbType {
	case ultipa.DBType_DBNODE:
		resp, err = api.doCreateNodeProperty(schemaName, prop, conf)
	case ultipa.DBType_DBEDGE:
		resp, err = api.doCreateEdgeProperty(schemaName, prop, conf)
	default:
		return nil, errors.New("create property: unknown db type")
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return resp, errors.New(resp.Status.Message)
	}
	api.Logger.Log("Created Property : @" + schemaName + "." + prop.Name)
	return resp, err
}

func (api *UltipaAPI) CreatePropertyIfNotExist(schemaName string, dbType ultipa.DBType, prop *structs.Property, config *configuration.RequestConfig) (exist bool, err error) {
	err = CheckName(schemaName)
	if err != nil {
		return false, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(prop.Name)
	if err != nil {
		return false, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}
	property, err := api.GetProperty(schemaName, prop.Name, dbType, config)

	if err != nil {
		return false, err
	}

	if property == nil {
		_, err = api.CreateProperty(schemaName, dbType, prop, config)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (api *UltipaAPI) GetProperty(schemaName string, propertyName string, dbType ultipa.DBType, config *configuration.RequestConfig) (property *structs.Property, err error) {
	err = CheckName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}
	err = CheckName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}
	schema, err := api.GetSchema(schemaName, dbType, config)

	if err != nil {
		return nil, err
	}

	for _, prop := range schema.Properties {
		if prop.Name == propertyName {
			return prop, nil
		}
	}

	return nil, nil
}

func (api *UltipaAPI) GetNodeProperty(schemaName string, propertyName string, config *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(schemaName, propertyName, ultipa.DBType_DBNODE, config)
}

func (api *UltipaAPI) GetEdgeProperty(schemaName string, propertyName string, config *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(schemaName, propertyName, ultipa.DBType_DBEDGE, config)
}

func (api *UltipaAPI) CreateNodeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

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

	return api.doCreateNodeProperty(escapeSchemaName, prop, conf)
}

func (api *UltipaAPI) doCreateNodeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

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

	resp, err = api.UQL(uql, conf)
	return resp, err
}

func (api *UltipaAPI) CreateEdgeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

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

	return api.doCreateEdgeProperty(escapeSchemaName, prop, conf)
}

func (api *UltipaAPI) doCreateEdgeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

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
	resp, err = api.UQL(uql, conf)
	return resp, err
}

// Usage: AlterNodeProperty("@schemaName.propertyName", dbType *ultipa.DBType, &*structs.Property{Name, Desc}, *RequestConfig)
func (api *UltipaAPI) AlterNodeProperty(propertyName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	resp, err = api.UQL(fmt.Sprintf(`alter().node_property(%v).set({name: "%v", description: "%v"})`, propertyName, prop.Name, prop.Desc), config)

	return resp, err
}

// Usage: AlterEdgeProperty("@schemaName.propertyName", dbType *ultipa.DBType, &*structs.Property{Name, Desc}, *RequestConfig)
func (api *UltipaAPI) AlterEdgeProperty(propertyName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	resp, err = api.UQL(fmt.Sprintf(`alter().edge_property(%v).set({name: "%v", description: "%v"})`, propertyName, prop.Name, prop.Desc), conf)

	return resp, err
}

// Usage: DropNodeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropNodeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.UQL(fmt.Sprintf(`drop().node_property(%v)`, propertyName), config)

	return resp, err
}

// Usage: DropEdgeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropEdgeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.UQL(fmt.Sprintf(`drop().edge_property(%v)`, propertyName), config)

	return resp, err
}
