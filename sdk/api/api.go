package api

import (
	"context"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/connection"
	"ultipa-go-sdk/sdk/http"
)

// UQL, Insert, Export, Download ... API methods

type UltipaAPI struct {
	Pool *connection.ConnectionPool
}

func NewUltipaAPI(pool *connection.ConnectionPool) *UltipaAPI {
	return &UltipaAPI{
		Pool: pool,
	}
}

func  (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) (res http.DataItem, err error) {

	conf := api.Pool.Config
	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
	}

	conn, _ := api.Pool.GetConn()
	client := conn.GetClient()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(api.Pool.Config.Timeout) * time.Second)

	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout: conf.Timeout,
	})

	res.Data = resp

	return res,nil

}