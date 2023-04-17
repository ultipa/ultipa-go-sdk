package api

import (
	"fmt"
	"strconv"
	"time"
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
				if UqlItem.IsGlobal() {
					conf.CurrentGraph = "global"
				}
			} else if UqlItem.HasWrite() || config.UseMaster || conf.Consistency {
				ok, graph := UqlItem.ParseGraph()
				if ok && graph != "" {
					conf.CurrentGraph = graph
				}
				conn, err = api.Pool.GetMasterConn(conf)
			} else if UqlItem.HasExecTask() {
				conn, err = api.Pool.GetAnalyticsConn(conf)
			}
		}

	}

	if err != nil {
		return nil, conf, err
	}

	if conn == nil {
		conn, err = api.Pool.GetConn(conf)
		if err != nil {
			return nil, nil, err
		}
	}
	return conn, conf, nil
}

func (api *UltipaAPI) GetClient(config *configuration.RequestConfig) (ultipa.UltipaRpcsClient, *configuration.UltipaConfig, error) {

	conn, conf, err := api.GetConn(config)

	if err != nil {
		return nil, conf, err
	}

	client := conn.GetClient()
	api.Logger.Log(fmt.Sprintf("fetch client,  hit host:[%s], role [%v], graph=[%s]", conn.Host, conn.Role, conf.CurrentGraph))
	return client, conf, nil
}

func (api *UltipaAPI) GetControlClient(config *configuration.RequestConfig) (ultipa.UltipaControlsClient, error) {

	client, _, err := api.GetControlClientAndConfig(config)
	return client, err
}

func (api *UltipaAPI) GetControlClientAndConfig(config *configuration.RequestConfig) (ultipa.UltipaControlsClient, *configuration.UltipaConfig, error) {

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	config.UseControl = true

	conn, conf, err := api.GetConn(config)

	if err != nil {
		return nil, conf, err
	}
	client := conn.GetControlClient()
	api.Logger.Log(fmt.Sprintf("fetch control client, hit host:[%s], role [%v], graph=[%s]", conn.Host, conn.Role, conf.CurrentGraph))
	return client, conf, nil
}

// UQL send a uql string to ultipa graph, and return a http UQL Response
// get Alias from UQL Response and convert to any type you need by asNodes, asEdges, asPaths, asTable, as asArray...
// Check DataItem to learn more about UQL Response
func (api *UltipaAPI) UQL(uql string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	resp, conf, err := api.doExecuteUql(uql, config)
	if err != nil {
		return nil, err
	}

	uqlResp, err := http.NewUQLResponse(resp)

	if err != nil {
		return nil, err
	}

	if config != nil && config.Host != "" {
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

func (api *UltipaAPI) UQLStream(uql string, config *configuration.RequestConfig) (*http.UQLResponseStream, error) {
	resp, conf, err := api.doExecuteUql(uql, config)
	if err != nil {
		return nil, err
	}
	uqlResp, err := http.NewUQLResponseStream(resp)
	if config != nil && config.Host != "" {
		return uqlResp, err
	}
	if uqlResp.NeedRedirect() {
		err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)
		if err != nil {
			return nil, err
		}
		return api.UQLStream(uql, config)
	}
	return uqlResp, nil
}

func (api *UltipaAPI) doExecuteUql(uql string, config *configuration.RequestConfig) (ultipa.UltipaRpcs_UqlClient, *configuration.UltipaConfig, error) {
	var err error

	if config == nil {
		config = &configuration.RequestConfig{}
	}

	config.Uql = uql
	uqlItem := utils.NewUql(uql)
	isExtra := uqlItem.IsExtra()
	var client ultipa.UltipaRpcsClient
	var uqlExClient ultipa.UltipaControlsClient
	var conf *configuration.UltipaConfig
	if isExtra {
		uqlExClient, conf, err = api.GetControlClientAndConfig(config)
	} else {
		client, conf, err = api.GetClient(config)
	}

	if err != nil {
		return nil, conf, err
	}
	//CurrentGraph of conf may be changed by uql
	config.GraphName = conf.CurrentGraph
	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		defer cancel()
		return nil, conf, err
	}
	uqlRequest := api.buildUqlRequest(uql, config, conf)
	var resp ultipa.UltipaRpcs_UqlClient
	if isExtra {
		resp, err = uqlExClient.UqlEx(ctx, uqlRequest)
	} else {
		resp, err = client.Uql(ctx, uqlRequest)
	}

	if err != nil {
		// if get error, ex: unavailable
		err = api.Pool.RefreshClusterInfo(conf.CurrentGraph)

		if err != nil {
			return nil, conf, err
		}

		if isExtra {
			resp, err = uqlExClient.UqlEx(ctx, uqlRequest)
		} else {
			resp, err = client.Uql(ctx, uqlRequest)
		}

		if err != nil {
			return nil, conf, err
		}
	}
	return resp, conf, nil
}

// buildUqlRequest build uqlRequest according to requestConfig and configuration
func (api *UltipaAPI) buildUqlRequest(uql string, config *configuration.RequestConfig, conf *configuration.UltipaConfig) *ultipa.UqlRequest {
	uqlRequest := &ultipa.UqlRequest{
		GraphName: conf.CurrentGraph,
		Timeout:   uint32(conf.Timeout),
		Uql:       uql,
	}
	if config.ThreadNum > 0 {
		uqlRequest.ThreadNum = config.ThreadNum
	}
	if config.TimezoneOffset == 0 && config.Timezone == "" {
		_, offset := time.Now().Zone()
		uqlRequest.TzOffset = strconv.Itoa(offset)
	} else if config.TimezoneOffset != 0 {
		uqlRequest.TzOffset = strconv.FormatInt(config.TimezoneOffset, 10)
	} else if config.Timezone != "" {
		uqlRequest.Tz = config.Timezone
	}
	return uqlRequest
}

// test connections
func (api *UltipaAPI) Test() (bool, error) {
	conn, err := api.Pool.GetConn(nil)

	if err != nil {
		return false, err
	}
	client := conn.GetClient()
	ctx, cancel, err := api.Pool.NewContext(nil)
	if err != nil {
		return false, err
	}
	defer cancel()
	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "Conn Test",
	})

	if err != nil || resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return false, err
	}

	return true, err
}
func (api *UltipaAPI) GetActiveClientTest() (bool, *connection.Connection, error) {
	conn, err := api.Pool.GetConn(nil)

	if err != nil {
		return false, nil, err
	}
	client := conn.GetClient()
	ctx, cancel, err := api.Pool.NewContext(nil)
	if err != nil {
		return false, nil, err
	}
	defer cancel()
	resp, err := client.SayHello(ctx, &ultipa.HelloUltipaRequest{
		Name: "Conn Test",
	})

	if err != nil || resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return false, nil, err
	}

	return true, conn, err
}

func (api *UltipaAPI) SetCurrentGraph(graphName string) error {
	api.Config.CurrentGraph = graphName
	return nil
}

func (api *UltipaAPI) Close() error {
	return api.Pool.Close()
}

func (api *UltipaAPI) SafelyClose() error {
	if api != nil && api.Pool != nil {
		return api.Pool.Close()
	}
	return nil
}
