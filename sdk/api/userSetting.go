package api

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) SetUserSetting(setUserSetting *structs.SetUserSetting, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	var err error

	client, err := api.GetControlClient(requestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	res, err := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		UserName: setUserSetting.UserName,
		Opt:      ultipa.UserSettingRequest_OPT_SET,
		Type:     setUserSetting.Type,
		Data:     setUserSetting.Data,
	})
	if err != nil {
		return nil, err
	}

	//data := &http.DataItem{
	//	Alias: setUserSetting.UserName,
	//	Data:  res.Data,
	//}

	resp := &http.UQLResponse{
		Status: &http.Status{
			Message: res.Status.Msg,
			Code:    res.Status.ErrorCode,
		},
		//DataItemMap: map[string]struct {
		//	DataItem *http.DataItem
		//	Index    int
		//}{setUserSetting.UserName: {DataItem: data}},
	}

	return resp, err
}

func (api *UltipaAPI) GetUserSetting(getUserSetting *structs.GetUserSetting, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	var err error

	client, err := api.GetControlClient(requestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	res, err := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		UserName: getUserSetting.UserName,
		Opt:      ultipa.UserSettingRequest_OPT_GET,
		Type:     getUserSetting.Type,
	})
	if err != nil {
		return nil, err
	}

	data := &http.DataItem{
		Alias: "GetUserSetting",
		Data:  res.Data,
	}

	resp := &http.UQLResponse{
		Status: &http.Status{
			Message: res.Status.Msg,
			Code:    res.Status.ErrorCode,
		},
		DataItemMap: map[string]struct {
			DataItem *http.DataItem
			Index    int
		}{getUserSetting.UserName: {DataItem: data}},
	}

	return resp, err
}
