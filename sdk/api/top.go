package api

import (
	"fmt"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) Top(requestConfig *configuration.RequestConfig) (tops []*structs.Process, err error) {
	resp, err := api.Uql("top()", requestConfig)

	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf(resp.Status.Message)
	}

	di := resp.Alias(http.RESP_TOP_KEY)
	if di.Data == nil {
		return nil, fmt.Errorf("no data return")
	}
	tops, err = di.AsTops()

	return tops, err
}
