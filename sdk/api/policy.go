package api

import (
	"fmt"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowPolicy(requestConfig *configuration.RequestConfig) (policies []*structs.Policy, err error) {
	resp, err := api.Uql("show().policy()", requestConfig)
	if err != nil {
		return nil, err
	}

	policies, err = resp.Alias(http.RESP_POLICY_KEY).AsPolicies()
	if err != nil {
		return nil, err
	}

	return
}

func (api *UltipaAPI) GetPolicy(policyName string, requestConfig *configuration.RequestConfig) (policy *structs.Policy, err error) {
	uql := fmt.Sprintf(`show().policy("%s")`, policyName)
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.Status.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	policies, err := resp.Alias(http.RESP_POLICY_KEY).AsPolicies()
	if err != nil {
		return nil, err
	}
	if len(policies) != 0 {
		policy = policies[0]
	}

	return
}

func (api *UltipaAPI) CreatePolicy(policy *structs.Policy, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	if policy == nil {
		return nil, fmt.Errorf("policy cant be null")
	}

	if policy.Name == "" {
		return nil, fmt.Errorf("policy name is required")
	}

	uql := policy.ToCreatePolicyUql()
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) AlterPolicy(policy *structs.Policy, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := policy.ToAlterPolicyUql()
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) DropPolicy(policyName string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`drop().policy("%s")`, policyName)
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
