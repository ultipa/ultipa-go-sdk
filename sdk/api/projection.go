package api

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowProjection(config *configuration.RequestConfig) (projecctions []*structs.Projection, err error) {
	resp, err := api.Uql("show().projection()", config)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	projecctions, err = resp.Alias(http.RESP_PROJECTION_KEY).AsProjections()

	return projecctions, err
}
