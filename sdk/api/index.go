/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index
 * @Date: 2022/8/4 3:27 下午
 */

package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) CreateIndex(dbType ultipa.DBType, schemaName, propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`create().node_index(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`create().edge_index(@%v.%v)`, schemaName, propertyName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) ShowIndex(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error

	resp, err = api.Uql(fmt.Sprintf(`show().index()`), config)
	if err != nil {
		return nil, err
	}

	var indexes []*structs.Index
	indexes, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsIndexes()
	if err != nil {
		return nil, err
	}
	EdgeIndexes, err := resp.Alias(http.RESP_EDGE_INDEX_KEY).AsIndexes()
	if err != nil {
		return nil, err
	}

	indexes = append(indexes, EdgeIndexes...)

	return indexes, err
}

func (api *UltipaAPI) ShowEdgeIndex(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().edge_index()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_EDGE_INDEX_KEY).AsIndexes()

	return indexes, err
}

func (api *UltipaAPI) ShowNodeIndex(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().node_index()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsIndexes()

	return indexes, err
}

func (api *UltipaAPI) DropIndex(dbType ultipa.DBType, schemaName, propertyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {

	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_index(@%v.%v)`, schemaName, propertyName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_index(@%v.%v)`, schemaName, propertyName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) CreateFullText(dbType ultipa.DBType, schemaName, propertyName, fulltextName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`create().node_fulltext(@%v.%v, "%v")`, schemaName, propertyName, fulltextName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`create().edge_fulltext(@%v.%v, "%v")`, schemaName, propertyName, fulltextName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) ShowFullText(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().fulltext()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullText()
	if err != nil {
		return nil, err
	}

	edgeIndexes, err := resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullText()
	if err != nil {
		return nil, err
	}

	indexes = append(indexes, edgeIndexes...)

	return indexes, err
}

func (api *UltipaAPI) ShowEdgeFullText(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().edge_fulltext()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullText()

	return indexes, err
}

func (api *UltipaAPI) ShowNodeFullText(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().node_fulltext()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullText()

	return indexes, err
}
