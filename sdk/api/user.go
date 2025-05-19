package api

import (
	"fmt"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowUser(requestConfig *configuration.RequestConfig) (users []*structs.User, err error) {
	resp, err := api.Uql("show().user()", requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	users, err = resp.Alias(http.RESP_USER_KEY).AsUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (api *UltipaAPI) GetUser(userName string, requestConfig *configuration.RequestConfig) (user *structs.User, err error) {
	uql := fmt.Sprintf(`show().user("%s")`, userName)
	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	users, err := resp.Alias(http.RESP_USER_KEY).AsUsers()
	if err != nil {
		return nil, err
	}

	// if user no exists， uql will return err : user [userName] does not exist
	user = users[0]

	return user, nil
}

func (api *UltipaAPI) CreateUser(request *structs.CreateUser, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := request.ToCreateUserUql()
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) AlterUser(request *structs.AlterUser, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := request.ToAlterUserUql()
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) DropUser(userName string, requestConfig *configuration.RequestConfig) (resp *http.Response, err error) {
	uql := fmt.Sprintf(`drop().user("%s")`, userName)
	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
