package api

import (
	"fmt"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Kill(processId string, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	if processId == "" {
		return nil, fmt.Errorf("processId can not empty")
	}
	uql := fmt.Sprintf(`kill("%s")`, processId)

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, err
}
