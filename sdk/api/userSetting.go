package api

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) SetUserSetting(userName, Type, data string, config *configuration.RequestConfig) (*ultipa.UserSettingReply, error) {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		UserName: userName,
		Opt:      ultipa.UserSettingRequest_OPT_SET,
		Type:     Type,
		Data:     data,
	})

	return resp, err
}

func (api *UltipaAPI) GetUserSetting(userName, Type string, config *configuration.RequestConfig) (*ultipa.UserSettingReply, error) {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.UserSetting(ctx, &ultipa.UserSettingRequest{
		UserName: userName,
		Opt:      ultipa.UserSettingRequest_OPT_GET,
		Type:     Type,
	})

	return resp, err
}
