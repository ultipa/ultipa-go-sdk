package api

import (
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
)

func (api *UltipaAPI) Gql(gql string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	return api.query(gql, ultipa.QueryType_GQL, requestConfig)

}

func (api *UltipaAPI) GQLStream(gql string, cb func(*http.UQLResponse) error, requestConfig *configuration.RequestConfig) error {
	return api.queryStream(gql, ultipa.QueryType_GQL, cb, requestConfig)
}
