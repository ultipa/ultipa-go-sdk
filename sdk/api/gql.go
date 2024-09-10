package api

import (
	"errors"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Gql(gql string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {

	resp, _, err := api.doExecuteQuery(gql, ultipa.QueryType_GQL, requestConfig)
	if err != nil {
		return nil, err
	}

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		return nil, err
	}

	if uqlResp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(uqlResp.Status.Message)
	}

	if requestConfig != nil && requestConfig.Host != "" {
		return uqlResp, err
	}

	//if uqlResp.NeedRedirect() {
	//    err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)
	//    if err != nil {
	//        return nil, err
	//    }
	//    return api.Uql(gql, requestConfig)
	//}

	return uqlResp, nil
}

func (api *UltipaAPI) GQLStream(gql string, requestConfig *configuration.RequestConfig) (*http.UQLResponseStream, error) {
	resp, _, err := api.doExecuteQuery(gql, ultipa.QueryType_GQL, requestConfig)
	if err != nil {
		return nil, err
	}

	uqlResp, err := http.NewUQLResponseStream(resp)
	if err != nil {
		return nil, err
	}
	if uqlResp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(uqlResp.Status.Message)
	}

	if requestConfig != nil && requestConfig.Host != "" {
		return uqlResp, err
	}

	//if uqlResp.NeedRedirect() {
	//    err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)
	//    if err != nil {
	//        return nil, err
	//    }
	//    return api.UQLStream(gql, requestConfig)
	//}
	return uqlResp, nil
}
