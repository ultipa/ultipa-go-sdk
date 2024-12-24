package api

import (
	"fmt"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) DeleteNodes(filter string, config *configuration.InsertRequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf("delete().nodes(%s)", filter)
	if !config.Silent {
		uql = uql + " as nodes return nodes{*}"
	}
	resp, err := api.Uql(uql, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) DeleteEdges(filter string, config *configuration.InsertRequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf("delete().edges(%s)", filter)
	if !config.Silent {
		uql = uql + " as edges return edges{*}"
	}
	resp, err := api.Uql(uql, config.RequestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
