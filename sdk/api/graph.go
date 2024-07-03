package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
	"reflect"
	"strconv"
	"time"
)

func (api *UltipaAPI) ShowGraph(config *configuration.RequestConfig) (*http.ResponseGraphs, error) {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listGraph)
	res, err := api.Uql(uql.ToString(), config)
	if err != nil {
		return nil, err
	}
	if res.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(res.Status.Message)
	}
	table, err := res.GetSingleTable()
	if err != nil {
		return nil, err
	}
	var graphs []*http.ResponseGraph
	if !res.Status.IsSuccess() {
		return &http.ResponseGraphs{
			Status: res.Status,
			Graphs: graphs,
		}, nil
	}
	values := table.ToKV()
	for _, v := range values {
		id, _ := strconv.ParseInt(v.Get("id").(string), 10, 64)
		totalNodes, _ := strconv.ParseInt(v.Get("totalNodes").(string), 10, 64)
		totalEdges, _ := strconv.ParseInt(v.Get("totalEdges").(string), 10, 64)
		status := v.Get("status").(string)
		description := v.Get("description").(string)
		clusterId := ""
		if v := v.Get("clusterId"); v != nil {
			clusterId = v.(string)
		}
		graphs = append(graphs, &http.ResponseGraph{
			Id:          id,
			ClusterId:   clusterId,
			Name:        v.Get("name").(string),
			TotalNodes:  totalNodes,
			TotalEdges:  totalEdges,
			Status:      status,
			Description: description,
		})
	}
	return &http.ResponseGraphs{
		Status: res.Status,
		Graphs: graphs,
	}, nil
}

func (api *UltipaAPI) CreateGraphIfNotExit(graph *structs.GraphInfo, config *configuration.RequestConfig) (resp *http.UQLResponse, exist bool, err error) {
	exist, err = api.HasGraph(graph.Name, config)

	if exist {
		return nil, exist, err
	}

	resp, err = api.CreateGraph(graph, config)
	return resp, exist, err
}

func (api *UltipaAPI) CreateGraph(graph *structs.GraphInfo, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	resp, err := api.Uql(fmt.Sprintf(`create().graph("%v", "%v")`, graph.Name, graph.Description), config)

	if err != nil {
		return nil, err
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		api.Logger.Log("create graph failed : " + graph.Name + " " + resp.Status.Message)
		return resp, errors.New(resp.Status.Message)
	}

	api.Logger.Log("Creating Graph Request OK! - " + graph.Name)

	// Try to detect the graph is created, default times is 600
	times := 60
	for {
		if times < 0 {
			break
		}

		api.Logger.Log("Detecting New Graph - " + graph.Name + " Leader")
		clusterErr := api.Pool.RefreshClusterInfo(graph.Name)

		if clusterErr != nil {
			if reflect.TypeOf(clusterErr).Elem().String() != "utils.LeaderNotYetElectedError" {
				api.Logger.Log(fmt.Sprintf("failed to detect New Graph - %s Leader", graph.Name))
				return nil, clusterErr
			}
			continue
		}

		conn := api.Pool.GraphMgr.GetLeader(graph.Name)

		if conn != nil {
			api.Logger.Log("Detected New Graph - " + graph.Name + " Leader - OK")
			break
		}

		time.Sleep(time.Second)
		times--
	}

	return resp, err
}

func (api *UltipaAPI) DropGraph(graphName string, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	resp, err := api.Uql(fmt.Sprintf(`drop().graph("%v")`, graphName), config)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (api *UltipaAPI) HasGraph(graphName string, config *configuration.RequestConfig) (bool, error) {
	resp, err := api.ShowGraph(config)

	if err != nil {
		return false, err
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return false, errors.New(resp.Status.Message)
	}

	for _, graph := range resp.Graphs {

		if graph.Name == graphName {
			return true, nil
		}
	}

	return false, nil
}

func (api *UltipaAPI) GetGraph(graphName string, config *configuration.RequestConfig) (*http.ResponseGraphs, error) {
	resp, err := api.ShowGraph(config)

	if err != nil {
		return nil, err
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	for _, graph := range resp.Graphs {
		if graph.Name == graphName {
			resp.Graphs = []*http.ResponseGraph{graph}
			return resp, nil
		}
	}

	return nil, errors.New("graph not found")
}

func (api *UltipaAPI) AlterGraph(graphName, newGraphName, description string, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`alter().graph("%v").set({name: "%v", description: "%v"})`, graphName, newGraphName, description)
	// Only modify the description of the graphSet
	if newGraphName == "" {
		uql = fmt.Sprintf(`alter().graph("%v").set({description: "%v"})`, graphName, description)
	}

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) Truncate(truncate *structs.Truncate, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := ""
	if truncate.Schema != "*" {
		truncate.Schema = "@" + truncate.Schema
	} else {
		truncate.Schema = `"*"`
	}

	switch truncate.DbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`truncate().graph("%v").node(%v)`, truncate.GraphName, truncate.Schema)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`truncate().graph("%v").edge(%v)`, truncate.GraphName, truncate.Schema)
	case ultipa.DBType_DBGLOBAL:
		uql = fmt.Sprintf(`truncate().graph("%v")`, truncate.GraphName)
	}

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) Compact(graphName string, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`compact().graph("%v")`, graphName)

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, err
}

func (api *UltipaAPI) MountGraph(graphName string, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`mount().graph("%v")`, graphName)

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) UnmountGraph(graphName string, config *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`unmount().graph("%v")`, graphName)

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
