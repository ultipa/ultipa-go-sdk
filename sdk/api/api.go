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
	Pool   *connection.ConnectionPool
	Config *configuration.UltipaConfig
}

func NewUltipaAPI(pool *connection.ConnectionPool) *UltipaAPI {

	api := &UltipaAPI{
		Pool:   pool,
		Config: pool.Config,
	}

	return api
}

func (api *UltipaAPI) GetClient(config *configuration.RequestConfig) (ultipa.UltipaRpcsClient, *configuration.UltipaConfig, error) {
	var err error
	var conn *connection.Connection

	conf := api.Pool.Config

	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)

		// Check if User set Host Address
		if config.Host != "" {
			conn, err = connection.NewConnection(config.Host, conf)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if conn == nil {
		conn, err = api.Pool.GetConn(conf)
		if err != nil {
			return nil, nil, err
		}
	}

	client := conn.GetClient()

	return client, conf, nil
}

func (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	var err error

	if config == nil { config = &configuration.RequestConfig{}}

	config.Uql = uql
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext(config)

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
	panic("not implemented")
	return nil, nil
}

// test connections
func (api *UltipaAPI) Test() (bool, error) {
	conn, err := api.Pool.GetConn(nil)

	if err != nil {
		return false, err
	}
	client := conn.GetClient()
	ctx, _ := api.Pool.NewContext(nil)
	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "Conn Test",
	})

	if err != nil || resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return false, err
	}

	return true, err
}

func (api *UltipaAPI) SetCurrentGraph(graphName string) error {
	api.Config.CurrentGraph = graphName
	return nil
}

func (api *UltipaAPI) Close() error {
	return api.Pool.Close()
}
