package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Lte(dbType ultipa.DBType, schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	if schemaName == "" {
		schemaName = "*"
	}

	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`LTE().node_property(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`LTE().edge_property(@%v.%v)`, schemaName, propertyName)
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

func (api *UltipaAPI) Ufe(dbType ultipa.DBType, schemaName, propertyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	if schemaName == "" {
		schemaName = "*"
	}

	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`UFE().node_property(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`UFE().edge_property(@%v.%v)`, schemaName, propertyName)
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
