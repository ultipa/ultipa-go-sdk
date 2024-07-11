package api

import (
	"errors"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) Stats(config *configuration.RequestConfig) (stats *structs.Stat, err error) {
	resp, err := api.Uql("stats()", config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	stats, err = resp.Alias(http.RESP_STATISTIC_KEY).AsStats()

	return stats, err
}
