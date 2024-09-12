package api

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) Stats(requestConfig *configuration.RequestConfig) (stats *structs.License, err error) {
	return api.license("stats()", requestConfig)
}
