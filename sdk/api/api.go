package api

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
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

func  (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) ( *http.UQLResponse,  error) {

	conf := api.Pool.Config
	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
	}

	conn, err := api.Pool.GetConn()

	if err != nil {
		return nil, err
	}

	client := conn.GetClient()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(api.Pool.Config.Timeout) * time.Second)

	ctx = metadata.AppendToOutgoingContext(ctx, conf.ToMetaKV()...)

	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout: conf.Timeout,
		Uql: uql,
	})

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		log.Fatalln(err)
	}

	return uqlResp, nil
}


func  (api *UltipaAPI) UQLStream(uql string, config *configuration.RequestConfig) ( *http.UQLResponse,  error) {

	conf := api.Pool.Config
	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
	}

	conn, err := api.Pool.GetConn()

	if err != nil {
		return nil, err
	}

	client := conn.GetClient()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(api.Pool.Config.Timeout) * time.Second)

	ctx = metadata.AppendToOutgoingContext(ctx, conf.ToMetaKV()...)

	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout: conf.Timeout,
		Uql: uql,
	})

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		log.Fatalln(err)
	}

	return uqlResp, nil
}