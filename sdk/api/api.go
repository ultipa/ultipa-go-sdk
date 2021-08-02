package api

import (
	"log"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/connection"
	"ultipa-go-sdk/sdk/http"
)

// UQL, Insert, Export, Download ... API methods

type UltipaAPI struct {
	Pool *connection.ConnectionPool
	Config *configuration.UltipaConfig
}

func NewUltipaAPI(pool *connection.ConnectionPool) *UltipaAPI {
	api := &UltipaAPI{
		Pool: pool,
		Config: pool.Config,
	}
	return api
}

func (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	conf := api.Pool.Config

	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
	}

	conn, err := api.Pool.GetConn()

	if err != nil {
		return nil, err
	}

	client := conn.GetClient()

	ctx, _ := api.Pool.NewContext()

	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout:   conf.Timeout,
		Uql:       uql,
	})

	if err != nil {
		return nil, err
	}

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		log.Fatalln(err)
	}

	return uqlResp, nil
}

//todo:
func (api *UltipaAPI) UQLStream(uql string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	conf := api.Pool.Config
	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
	}

	conn, err := api.Pool.GetConn()

	if err != nil {
		return nil, err
	}

	client := conn.GetClient()

	ctx, _ := api.Pool.NewContext()

	//todo: UqlStream Function
	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout:   conf.Timeout,
		Uql:       uql,
	})

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		log.Fatalln(err)
	}

	return uqlResp, nil
}

// test connections
func (api *UltipaAPI) Test() (bool, error) {
	conn, _ := api.Pool.GetConn()

	ctx, _ := api.Pool.NewContext()
	resp, err := conn.Client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "Conn Test",
	})

	if err != nil || resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return false, err
	}

	return true, err
}

func (api *UltipaAPI) Close() error {
	return api.Pool.Close()
}
