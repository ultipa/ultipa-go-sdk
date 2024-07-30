package api

import (
	"errors"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) Authenticate(authenticateType ultipa.AuthenticateType, uql string, config *configuration.RequestConfig) (*ultipa.AuthenticateReply, error) {

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

	resp, err := client.Authenticate(ctx, &ultipa.AuthenticateRequest{
		Type: authenticateType,
		Uql:  uql,
	})

	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Msg)
	}

	return resp, nil
}
