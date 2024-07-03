package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowPolicy(config *configuration.RequestConfig) (*http.ResponsePolicy, error) {
	resp, err := api.Uql("show().policy()", config)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	var policies []*structs.Policy
	policies, err = resp.Alias(http.RESP_NODE_INDEX_KEY).AsPolicy()
	if err != nil {
		return nil, err
	}

	r := &http.ResponsePolicy{
		Status:   resp.Status,
		Policies: policies,
	}

	return r, nil
}

func (api *UltipaAPI) GetPolicy(policyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`show().policy("%s")`, policyName)
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
