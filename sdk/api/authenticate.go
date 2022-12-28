package api

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) Authenticate(authenticateType ultipa.AuthenticateType, uql string, requestConfig *configuration.RequestConfig) (*ultipa.AuthenticateReply, error) {

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

	resp, err := client.Authenticate(ctx, &ultipa.AuthenticateRequest{
		Type: authenticateType,
		Uql:  uql,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
