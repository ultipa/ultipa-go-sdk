package api

import (
	"errors"
	"fmt"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

// CreateProperty create property for schema, schemaName maybe escaped if schemaName contains some special characters.
func (api *UltipaAPI) CreateProperty(dbType ultipa.DBType, schemaName string, prop *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	params := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		params = "node_property"
	case ultipa.DBType_DBEDGE:
		params = "edge_property"
	default:
		errStr := fmt.Sprintf("create property: unknown db type, %v", dbType.String())
		return nil, errors.New(errStr)
	}

	if prop.Type == structs.PropertyType_IGNORE {
		return nil, err
	}

	propertyTypeStr, err := prop.GetStringType()
	if err != nil {
		return nil, err
	}

	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err := CheckReplaceSchemaPropertyName(prop.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), prop.Name))
	}

	api.Logger.Log("Creating Property : @" + schemaName + "." + propertyName)
	uql := fmt.Sprintf(`create().%v(@%v,%s,"%v","%v")`, params, schemaName, propertyName, propertyTypeStr, prop.Desc)
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	api.Logger.Log("Created Property : @" + schemaName + "." + propertyName)
	return resp, nil
}

func (api *UltipaAPI) CreatePropertyIfNotExist(dbType ultipa.DBType, schemaName string, prop *structs.Property, requestConfig *configuration.RequestConfig) (exist bool, resp *http.UQLResponse, err error) {
	property, err := api.GetProperty(dbType, schemaName, prop.Name, requestConfig)

	if err != nil {
		return false, nil, err
	}

	if property == nil {
		resp, err = api.CreateProperty(dbType, schemaName, prop, requestConfig)
		if err != nil {
			return false, resp, err
		}

		return false, resp, nil
	}

	return true, resp, nil
}

func (api *UltipaAPI) GetProperty(dbType ultipa.DBType, schemaName string, propertyName string, requestConfig *configuration.RequestConfig) (property *structs.Property, err error) {
	schema, err := api.GetSchema(schemaName, dbType, requestConfig)

	if err != nil {
		return nil, err
	}

	if schema == nil {
		return nil, nil
	}

	for _, prop := range schema.Properties {
		if prop.Name == propertyName {
			return prop, nil
		}
	}

	//return nil, fmt.Errorf("property %s not found in schema %s", propertyName, schemaName)
	return nil, nil
}

func (api *UltipaAPI) ShowProperty(dbType ultipa.DBType, schemaName string, requestConfig *configuration.RequestConfig) (nodeProperty, edgeProperty []*structs.Property, err error) {
	//resp, err := api.Uql("show().property()", requestConfig)
	////
	//if err != nil {
	//    return nil, nil, err
	//}
	//
	//nodeProperty, err = resp.Alias(http.RESP_NODE_PROPERTY_KEY).AsProperties()
	//if err != nil {
	//    return nil, nil, err
	//}
	//edgeProperty, err = resp.Alias(http.RESP_EDGE_PROPERTY_KEY).AsProperties()
	//if err != nil {
	//    return nil, nil, err
	//}
	//
	//return nodeProperty, edgeProperty, nil

	switch dbType {
	case ultipa.DBType_DBNODE:
		nodeProperty, err = api.ShowNodeProperty(schemaName, requestConfig)
	case ultipa.DBType_DBEDGE:
		edgeProperty, err = api.ShowEdgeProperty(schemaName, requestConfig)
	default:
		nodeProperty, err = api.ShowNodeProperty(schemaName, requestConfig)
		edgeProperty, err = api.ShowEdgeProperty(schemaName, requestConfig)
	}
	return
}

func (api *UltipaAPI) ShowNodeProperty(schemaName string, requestConfig *configuration.RequestConfig) (property []*structs.Property, err error) {
	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	if schemaName == "" || schemaName == "*" {
		schemaName = ""
	} else {
		schemaName = "@" + schemaName
	}
	uql := fmt.Sprintf("show().node_property(%s)", schemaName)
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	property, err = resp.Alias(http.RESP_NODE_PROPERTY_KEY).AsProperties()
	return property, nil
}

func (api *UltipaAPI) ShowEdgeProperty(schemaName string, requestConfig *configuration.RequestConfig) (property []*structs.Property, err error) {
	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	if schemaName == "" || schemaName == "*" {
		schemaName = ""
	} else {
		schemaName = "@" + schemaName
	}
	uql := fmt.Sprintf("show().edge_property(%s)", schemaName)
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	property, err = resp.Alias(http.RESP_EDGE_PROPERTY_KEY).AsProperties()
	return property, nil
}

func (api *UltipaAPI) GetNodeProperty(schemaName string, propertyName string, requestConfig *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(ultipa.DBType_DBNODE, schemaName, propertyName, requestConfig)
}

func (api *UltipaAPI) GetEdgeProperty(schemaName string, propertyName string, requestConfig *configuration.RequestConfig) (property *structs.Property, err error) {
	return api.GetProperty(ultipa.DBType_DBEDGE, schemaName, propertyName, requestConfig)
}

func (api *UltipaAPI) CreateNodeProperty(schemaName string, prop *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.CreateProperty(ultipa.DBType_DBNODE, schemaName, prop, requestConfig)
}

func (api *UltipaAPI) CreateEdgeProperty(schemaName string, prop *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.CreateProperty(ultipa.DBType_DBEDGE, schemaName, prop, requestConfig)
}

func (api *UltipaAPI) AlterNodeProperty(property, newProperty *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.AlterProperty(ultipa.DBType_DBNODE, property, newProperty, requestConfig)
}

func (api *UltipaAPI) AlterEdgeProperty(property, newProperty *structs.Property, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.AlterProperty(ultipa.DBType_DBEDGE, property, newProperty, requestConfig)
}

func (api *UltipaAPI) DropNodeProperty(schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.DropProperty(ultipa.DBType_DBNODE, schemaName, propertyName, requestConfig)
}

func (api *UltipaAPI) DropEdgeProperty(schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	return api.DropProperty(ultipa.DBType_DBEDGE, schemaName, propertyName, requestConfig)
}

func (api *UltipaAPI) DropProperty(dbType ultipa.DBType, schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}

	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_property(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_property(@%v.%v)`, schemaName, propertyName)
	default:
		return nil, errors.New("drop property: unknown db type")
	}

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) AlterProperty(dbType ultipa.DBType, property, newProperty *structs.Property, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	params := ""

	switch dbType {
	case ultipa.DBType_DBNODE:
		params = "node_property"
	case ultipa.DBType_DBEDGE:
		params = "edge_property"
	default:
		return nil, errors.New("alter property: unknown db type")
	}

	schemaName, err := CheckReplaceSchemaPropertyName(property.Schema)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), property.Schema))
	}

	propertyName, err := CheckReplaceSchemaPropertyName(property.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), property.Name))
	}

	uql := fmt.Sprintf(`alter().%v(@%v.%v).set({name: "%v", description: "%v"})`, params, schemaName, propertyName, newProperty.Name, newProperty.Desc)

	// Only modify the description of the property
	if newProperty.Name == "" {
		uql = fmt.Sprintf(`alter().%v(@%v.%v).set({description: "%v"})`, params, schemaName, propertyName, newProperty.Desc)
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
