package api

import (
	"errors"
	"fmt"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Kill(processId string, all bool, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := "kill()"
	if all {
		uql = `kill("*")`
	} else {
		uql = fmt.Sprintf(`kill("%s")`, processId)
	}

	resp, err = api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, err
}
