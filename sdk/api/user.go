package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowUser(requestConfig *configuration.RequestConfig) (users []*structs.User, err error) {
	resp, err := api.Uql("show().user()", requestConfig)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	users, err = resp.Alias(http.RESP_USER_KEY).AsUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (api *UltipaAPI) GetUser(name string, config *configuration.RequestConfig) (user *structs.User, err error) {
	uql := fmt.Sprintf(`show().user("%s")`, name)
	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}
	users, err := resp.Alias(http.RESP_USER_KEY).AsUser()
	if err != nil {
		return nil, err
	}
	user = users[0]

	return user, nil
}

func (api *UltipaAPI) CreateUser(user *structs.User, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := user.ToCreateUserUql()
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) AlterUser(user *structs.User, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := user.ToAlterUserUql()
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) DropUser(userName string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf(`drop().user("%s")`, userName)
	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
