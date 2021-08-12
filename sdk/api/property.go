package api

import (
	"errors"
	"fmt"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/structs"
)
func (api *UltipaAPI) CreateProperty(schemaName string, dbType ultipa.DBType, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	switch dbType {
	case ultipa.DBType_DBNODE:
		return api.CreateNodeProperty(schemaName, prop, conf)
	case ultipa.DBType_DBEDGE:
		return api.CreateEdgeProperty(schemaName, prop, conf)
	default:
		return nil, errors.New("create property: unknown db type")
	}
}

func (api *UltipaAPI) CreateNodeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	uql := fmt.Sprintf(`create().node_property(@%v,"%v",%v,"%v")`, schemaName, prop.Name, prop.GetStringType(), prop.Desc)

	resp, err = api.UQL(uql, conf)

	return resp, err
}

func (api *UltipaAPI) CreateEdgeProperty(schemaName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	resp, err = api.UQL(fmt.Sprintf(`create().edge_property(@%v,"%v",%v,"%v")`, schemaName, prop.Name, prop.GetStringType(), prop.Desc), conf)

	return resp, err
}

// Usage: AlterNodeProperty("@schemaName.propertyName", dbType *ultipa.DBType, &*structs.Property{Name, Desc}, *RequestConfig)
func (api *UltipaAPI) AlterNodeProperty(propertyName string, prop *structs.Property, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	resp, err = api.UQL(fmt.Sprintf(`alter().node_property(@%v.%v).set({name: "%v", description: "%v"})`, propertyName, prop.Name, prop.Desc), config)

	return resp, err
}

// Usage: AlterEdgeProperty("@schemaName.propertyName", dbType *ultipa.DBType, &*structs.Property{Name, Desc}, *RequestConfig)
func (api *UltipaAPI) AlterEdgeProperty(propertyName string, prop *structs.Property, conf *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	resp, err = api.UQL(fmt.Sprintf(`alter().edge_property(@%v.%v).set({name: "%v", description: "%v"})`, propertyName, prop.Name, prop.Desc), conf)

	return resp, err
}

// Usage: DropNodeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropNodeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.UQL(fmt.Sprintf(`drop().node_property(@%v.%v)`, propertyName), config)

	return resp, err
}

// Usage: DropEdgeProperty("@schemaName.propertyName", *RequestConfig)
func (api *UltipaAPI) DropEdgeProperty(propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	resp, err = api.UQL(fmt.Sprintf(`drop().node_property(@%v.%v)`, propertyName), config)

	return resp, err
}
