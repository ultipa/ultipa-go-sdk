package api

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/connection"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/utils"
	"ultipa-go-sdk/sdk/utils/logger"
)

// UQL, Insert, Export, Download ... API methods

type UltipaAPI struct {
	Pool   *connection.ConnectionPool
	Config *configuration.UltipaConfig
	Logger *logger.Logger
}

type ClientType int

const (
	ClientTypeGeneral ClientType = 1
	ClientTypeControl ClientType = 2
)

func NewUltipaAPI(pool *connection.ConnectionPool) *UltipaAPI {

	api := &UltipaAPI{
		Pool:   pool,
		Config: pool.Config,
		Logger: logger.NewLogger(pool.Config.Debug),
	}

	return api
}

func (api *UltipaAPI) GetConn(config *configuration.RequestConfig) (*connection.Connection, *configuration.UltipaConfig, error) {
	var err error
	var conn *connection.Connection

	conf := api.Pool.Config

	if config != nil {
		conf = api.Pool.Config.MergeRequestConfig(config)
		UqlItem := utils.NewUql(config.Uql)

		// Check if User set Host Address
		if config.Host != "" {
			conn, err = connection.NewConnection(config.Host, conf)
			if err != nil {
				return nil, nil, err
			}

			// if is raft mode, check if contains CUD ops or exec task
		} else if api.Pool.IsRaft {
			if UqlItem.IsGlobal() || config.UseControl {
				conn, err = api.Pool.GetGlobalMasterConn(conf)
			} else if UqlItem.HasWrite() || config.UseMaster || conf.Consistency  {
				conn, err = api.Pool.GetMasterConn(conf)
			} else if UqlItem.HasExecTask() {
				conn, err = api.Pool.GetAnalyticsConn(conf)
			}
		}

	}

	if conn == nil {
		conn, err = api.Pool.GetConn(conf)
		if err != nil {
			return nil, nil, err
		}
	}

	if err != nil {
		return nil, conf, err
	}

	return conn, conf, nil
}

func (api *UltipaAPI) GetClient(config *configuration.RequestConfig) (ultipa.UltipaRpcsClient, *configuration.UltipaConfig, error) {

	conn, conf, err := api.GetConn(config)

	if err != nil {
		return nil, conf, err
	}

	client := conn.GetClient()

	return client, conf, nil
}

func (api *UltipaAPI) GetControlClient(config *configuration.RequestConfig) (ultipa.UltipaControlsClient,  error) {

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	config.UseControl = true

	conn, _, err := api.GetConn(config)

	if err != nil {
		return nil, err
	}
	client := conn.GetControlClient()

	return client, nil
}

// UQL send a uql string to ultipa graph, and return a http UQL Response
// get Alias from UQL Response and convert to any type you need by asNodes, asEdges, asPaths, asTable, as asArray...
// Check DataItem to learn more about UQL Response
func (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	var err error

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	config.Uql = uql
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, cancel := api.Pool.NewContext(config)
	defer cancel()

	resp, err := client.Uql(ctx, &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout:   conf.Timeout,
		Uql:       uql,
	})

	if err != nil {
		// if get error, ex: unavailable
		err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)
		resp, err = client.Uql(ctx, &ultipa.UqlRequest{
			GraphName: conf.CurrentGraph,
			Timeout:   conf.Timeout,
			Uql:       uql,
		})
	}

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		return nil, err
	}

	if config.Host != "" {
		return uqlResp, err
	}

	if uqlResp.NeedRedirect() {
		err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)
		if err != nil {
			return nil, err
		}
		return api.UQL(uql, config)
	}

	return uqlResp, nil
}

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
