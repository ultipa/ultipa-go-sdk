package api

import (
	"strconv"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/utils"
)

func  (api *UltipaAPI) ListGraph(config *configuration.RequestConfig) ( *http.ResponseGraphs,  error) {
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
		clusterId := ""
		if v := v.Get("clusterId"); v != nil {
			clusterId = v.(string)
		}
		graphs = append(graphs, &http.ResponseGraph{
			Id:         id,
			ClusterId:  clusterId,
			Name:       v.Get("name").(string),
			TotalNodes: totalNodes,
			TotalEdges: totalEdges,
		})
	}
	return &http.ResponseGraphs{
		Status: res.Status,
		Graphs: graphs,
	}, nil
}
