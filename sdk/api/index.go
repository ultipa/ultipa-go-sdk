/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index
 * @Date: 2022/8/4 3:27 pm
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

// Deprecated: 5.0 not support, should use CreateNodeIndex or CreateEdgeIndex
func (api *UltipaAPI) CreateIndex(dbType ultipa.DBType, schemaName, propertyName string, indexName string, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := ""
	if schemaName == "" {
		return nil, fmt.Errorf("schemaName can not empty %s", schemaName)
	}

	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}

	indexName, err = CheckReplaceSchemaPropertyName(indexName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), indexName))
	}

	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`create().node_index(@%v.%v, %s)`, schemaName, propertyName, indexName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`create().edge_index(@%v.%v, %s)`, schemaName, propertyName, indexName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) CreateEdgeIndex(source string, indexName string, requestConfig *configuration.RequestConfig) (jobResponse *http.JobResponse, err error) {
	uql := ""
	if source == "" {
		return nil, fmt.Errorf("source can not empty %s", source)
	}

	indexName, err = CheckReplaceSchemaPropertyName(indexName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), indexName))
	}

	uql = fmt.Sprintf(`create().edge_index(%s, %s)`, source, indexName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	return http.GetJobResponseFromUqlResponse(resp)
}

func (api *UltipaAPI) CreateNodeIndex(source string, indexName string, requestConfig *configuration.RequestConfig) (jobResponse *http.JobResponse, err error) {
	uql := ""
	if source == "" {
		return nil, fmt.Errorf("source can not empty %s", source)
	}

	indexName, err = CheckReplaceSchemaPropertyName(indexName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), indexName))
	}

	uql = fmt.Sprintf(`create().node_index(%s, %s)`, source, indexName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	return http.GetJobResponseFromUqlResponse(resp)
}

func (api *UltipaAPI) ShowIndex(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error

	resp, err = api.Uql(fmt.Sprintf(`show().index()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
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

func (api *UltipaAPI) ShowEdgeIndex(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().edge_index()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	indexes, err = resp.Alias(http.RESP_EDGE_INDEX_KEY).AsIndexes()

	return indexes, err
}

func (api *UltipaAPI) ShowNodeIndex(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().node_index()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	indexes, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsIndexes()

	return indexes, err
}

func (api *UltipaAPI) DropIndex(dbType ultipa.DBType, indexName string, config *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := ""
	//if schemaName == "" {
	//	schemaName = "*"
	//}
	//
	//schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	//}
	//
	//propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	//}

	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_index("%v")`, indexName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_index("%v")`, indexName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) CreateFullText(dbType ultipa.DBType, schemaName, propertyName, indexName string, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := ""

	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}

	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`create().node_fulltext(@%v.%v, "%v")`, schemaName, propertyName, indexName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`create().edge_fulltext(@%v.%v, "%v")`, schemaName, propertyName, indexName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) CreateNodeFullText(schemaName, propertyName, indexName string, requestConfig *configuration.RequestConfig) (jobResponse *http.JobResponse, err error) {
	uql := ""

	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}

	uql = fmt.Sprintf(`create().node_fulltext(@%v.%v, "%v")`, schemaName, propertyName, indexName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	return http.GetJobResponseFromUqlResponse(resp)
}

func (api *UltipaAPI) CreateEdgeFullText(schemaName, propertyName, indexName string, requestConfig *configuration.RequestConfig) (jobResponse *http.JobResponse, err error) {
	uql := ""

	schemaName, err = CheckReplaceSchemaPropertyName(schemaName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, schemaName = %s", err.Error(), schemaName))
	}

	propertyName, err = CheckReplaceSchemaPropertyName(propertyName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, propertyName = %s", err.Error(), propertyName))
	}

	uql = fmt.Sprintf(`create().edge_fulltext(@%v.%v, "%v")`, schemaName, propertyName, indexName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	return http.GetJobResponseFromUqlResponse(resp)
}

func (api *UltipaAPI) ShowFullText(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().fulltext()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullTexts()
	if err != nil {
		return nil, err
	}

	edgeIndexes, err := resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullTexts()
	if err != nil {
		return nil, err
	}

	indexes = append(indexes, edgeIndexes...)

	return indexes, err
}

func (api *UltipaAPI) ShowEdgeFullText(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().edge_fulltext()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	indexes, err = resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullTexts()

	return indexes, err
}

func (api *UltipaAPI) ShowNodeFullText(requestConfig *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.Response
	var err error
	var indexes []*structs.Index

	resp, err = api.Uql(fmt.Sprintf(`show().node_fulltext()`), requestConfig)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullTexts()

	return indexes, err
}

func (api *UltipaAPI) DropFullText(fullTextName string, dbType ultipa.DBType, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {

	uql := ""
	switch dbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`drop().node_fulltext("%v")`, fullTextName)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`drop().edge_fulltext("%v")`, fullTextName)
	default:
		return nil, errors.New("DBType must be DBType_DBNODE or DBType_DBEDGE")
	}

	resp, err = api.Uql(uql, requestConfig)
	if err != nil {
		return nil, err
	}

	return resp, err
}
