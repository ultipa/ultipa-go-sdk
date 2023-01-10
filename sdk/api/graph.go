package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/structs"
	"ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ListGraph(config *configuration.RequestConfig) (*http.ResponseGraphs, error) {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listGraph)
	res, err := api.UQL(uql.ToString(), config)
	if err != nil {
		return nil, err
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

func (api *UltipaAPI) CreateGraph(graph *structs.Graph, config *configuration.RequestConfig) (*http.UQLResponse, error) {

	resp, err := api.UQL(fmt.Sprintf(`create().graph("%v", "%v")`, graph.Name, graph.Description), config)

	if err != nil {
		return nil, err
	}

	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		api.Logger.Log("create graph failed : " + graph.Name + " " + resp.Status.Message)
		return resp, errors.New(resp.Status.Message)
	}

	api.Logger.Log("Creating Graph Request OK! - " + graph.Name)

	// Try to detect the graph is created, default times is 600
	times := 600
	for {
		if times < 0 {
			break
		}

		api.Logger.Log("Detecting New Graph - " + graph.Name + " Leader")
		err := api.Pool.RefreshClusterInfo(graph.Name)

		if err != nil && !strings.Contains(err.Error(), "error graph name") {
			return nil, err
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

	resp, err := api.UQL(fmt.Sprintf(`drop().graph("%v")`, graphName), config)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (api *UltipaAPI) HasGraph(graphName string, config *configuration.RequestConfig) (bool, error) {
	resp, err := api.ListGraph(config)

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
