package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowPolicy(config *configuration.RequestConfig) (policies []*structs.Policy, err error) {
	resp, err := api.Uql("show().policy()", config)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	policies, err = resp.Alias(http.RESP_POLICY_KEY).AsPolicy()
	if err != nil {
		return nil, err
	}

	return
}

func (api *UltipaAPI) GetPolicy(policyName string, config *configuration.RequestConfig) (policy *structs.Policy, err error) {
	uql := fmt.Sprintf(`show().policy("%s")`, policyName)
	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	policies, err := resp.Alias(http.RESP_POLICY_KEY).AsPolicy()
	if err != nil {
		return nil, err
	}
	policy = policies[0]

	return
}

func (api *UltipaAPI) CreatePolicy(policy *structs.Policy, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := policy.ToCreatePolicyUql()
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) AlterPolicy(policy *structs.Policy, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := policy.ToAlterPolicyUql()
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) DropPolicy(policyName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`drop().policy("%s")`, policyName)
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
