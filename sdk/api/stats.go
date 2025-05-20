package api

import (
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Stats(requestConfig *configuration.RequestConfig) (stats *http.Response, err error) {
	//return api.license("stats()", requestConfig)
	return api.Uql("stats()", requestConfig)
}
