/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  index
 * @Date: 2022/8/4 3:27 下午
 */

package api

import (
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ListIndex(config *configuration.RequestConfig) ([]*http.ResponseIndex, error) {
	var resp *http.UQLResponse
	var err error
	var responseIndexes []*http.ResponseIndex

	resp, err = api.UQL(fmt.Sprintf(`show().index()`), config)
	if err != nil {
		return nil, err
	}

	for _, alias := range resp.AliasList {
		var indexes []*structs.Index
		var r *http.ResponseIndex
		if alias == http.RESP_NODE_INDEX_KEY {
			indexes, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsIndexes()
			if err != nil {
				return nil, err
			}
			r = &http.ResponseIndex{
				Type:    ultipa.DBType_DBNODE,
				Indexes: indexes,
			}

		}

		if alias == http.RESP_EDGE_INDEX_KEY {
			indexes, err = resp.Alias(http.RESP_EDGE_INDEX_KEY).AsIndexes()
			if err != nil {
				return nil, err
			}
			r = &http.ResponseIndex{
				Type:    ultipa.DBType_DBEDGE,
				Indexes: indexes,
			}
		}
		responseIndexes = append(responseIndexes, r)

	}

	return responseIndexes, err
}

func (api *UltipaAPI) ListEdgeIndex(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.UQL(fmt.Sprintf(`show().edge_index()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_EDGE_INDEX_KEY).AsIndexes()

	if len(indexes) == 0 {
		return nil, err
	}

	return indexes, err
}

func (api *UltipaAPI) ListNodeIndex(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.UQL(fmt.Sprintf(`show().node_index()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsIndexes()

	if len(indexes) == 0 {
		return nil, err
	}

	return indexes, err
}

func (api *UltipaAPI) ListFullText(config *configuration.RequestConfig) ([]*http.ResponseIndex, error) {
	var resp *http.UQLResponse
	var err error
	var responseIndexes []*http.ResponseIndex

	resp, err = api.UQL(fmt.Sprintf(`show().fulltext()`), config)
	if err != nil {
		return nil, err
	}

	for _, alias := range resp.AliasList {
		var indexes []*structs.Index
		var r *http.ResponseIndex
		if alias == http.RESP_NODE_FULLTEXT_KEY {
			indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullText()
			if err != nil {
				return nil, err
			}
			r = &http.ResponseIndex{
				Type:    ultipa.DBType_DBNODE,
				Indexes: indexes,
			}

		}

		if alias == http.RESP_EDGE_FULLTEXT_KEY {
			indexes, err = resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullText()
			if err != nil {
				return nil, err
			}
			r = &http.ResponseIndex{
				Type:    ultipa.DBType_DBEDGE,
				Indexes: indexes,
			}
		}
		responseIndexes = append(responseIndexes, r)

	}

	return responseIndexes, err
}

func (api *UltipaAPI) ListEdgeFullText(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.UQL(fmt.Sprintf(`show().edge_fulltext()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_EDGE_FULLTEXT_KEY).AsFullText()

	if len(indexes) == 0 {
		return nil, err
	}

	return indexes, err
}

func (api *UltipaAPI) ListNodeFullText(config *configuration.RequestConfig) ([]*structs.Index, error) {
	var resp *http.UQLResponse
	var err error
	var indexes []*structs.Index

	resp, err = api.UQL(fmt.Sprintf(`show().node_fulltext()`), config)
	if err != nil {
		return nil, err
	}

	indexes, err = resp.Alias(http.RESP_NODE_FULLTEXT_KEY).AsFullText()

	if len(indexes) == 0 {
		return nil, err
	}

	return indexes, err
}
