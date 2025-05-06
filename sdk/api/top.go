package api

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) Top(requestConfig *configuration.RequestConfig) (tops []*structs.Process, err error) {
	resp, err := api.Uql("top()", requestConfig)

	if err != nil {
		return nil, err
	}

	tops, err = resp.Alias(http.RESP_TOP_KEY).AsTops()

	return tops, err
}
