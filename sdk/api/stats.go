package api

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) Stats(requestConfig *configuration.RequestConfig) (stats *structs.Stat, err error) {
	resp, err := api.Uql("stats()", requestConfig)

	if err != nil {
		return nil, err
	}

	stats, err = resp.Alias(http.RESP_STATISTIC_KEY).AsStats()

	return stats, err
}
