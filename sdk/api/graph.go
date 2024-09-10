package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"strconv"
)

func (api *UltipaAPI) ShowGraph(requestConfig *configuration.RequestConfig) (graphSets []*structs.GraphSet, err error) {
	res, err := api.Uql("show().graph().more()", requestConfig)
	if err != nil {
		return nil, err
	}

	g, err := res.Alias(http.RESP_GRAPH_KEY).AsTable()
	if err != nil {
		return nil, err
	}

	values := g.ToKV()
	for _, v := range values {
		id := v.Get("id").(string)
		name := v.Get("name").(string)
		status := v.Get("status").(string)
		description := v.Get("description").(string)
		shards := v.Get("shards").(string)
		slotNum := v.Get("slot_num").(string)
		replicaNum := v.Get("replica_num").(string)
		partitionBy := v.Get("partition_by").(string)

		var totalNodes uint64 = 0
		if v := v.Get("total_nodes"); v != nil {
			totalNodes, _ = strconv.ParseUint(v.(string), 10, 64)

		}
		var totalEdges uint64 = 0
		if v := v.Get("total_edges"); v != nil {
			totalEdges, err = strconv.ParseUint(v.(string), 10, 64)
		}

		//clusterId := ""
		//if v := v.Get("clusterId"); v != nil {
		//    clusterId = v.(string)
		//}
		graphSets = append(graphSets, &structs.GraphSet{
			ID: id,
			//ClusterId:   clusterId,
			Name:        name,
			TotalNodes:  totalNodes,
			TotalEdges:  totalEdges,
			Status:      status,
			Description: description,
			Shards:      shards,
			SlotNum:     slotNum,
			ReplicaNum:  replicaNum,
			PartitionBy: partitionBy,
		})
	}

	return graphSets, nil
}

func (api *UltipaAPI) CreateGraphIfNotExist(graph *structs.GraphSet, requestConfig *configuration.RequestConfig) (resp *http.UQLResponse, exist bool, err error) {
	exist, err = api.HasGraph(graph.Name, requestConfig)

	if exist {
		return nil, exist, err
	}

	resp, err = api.CreateGraph(graph, requestConfig)
	return resp, exist, err
}

// CreateGraph 5.0 partitionByHash:Crc32/CityHash64
func (api *UltipaAPI) CreateGraph(graph *structs.GraphSet, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	resp, err := api.Uql(fmt.Sprintf(`create().graph("%v", "%v").shards([%v]).partitionByHash('%v',_id)`, graph.Name, graph.Description, graph.Shards, graph.PartitionBy), requestConfig)

	if err != nil {
		api.Logger.Log("create graph failed : " + graph.Name + " " + err.Error())
		return nil, err
	}

	api.Logger.Log("Creating Graph Request OK! - " + graph.Name)

	// Try to detect the graph is created, default times is 600
	//times := 60
	//for {
	//    if times < 0 {
	//        break
	//    }
	//
	//    api.Logger.Log("Detecting New Graph - " + graph.Name + " Leader")
	//    clusterErr := api.Conn.RefreshClusterInfo(graph.Name)
	//
	//    if clusterErr != nil {
	//        if reflect.TypeOf(clusterErr).Elem().String() != "utils.LeaderNotYetElectedError" {
	//            api.Logger.Log(fmt.Sprintf("failed to detect New Graph - %s Leader", graph.Name))
	//            return nil, clusterErr
	//        }
	//        continue
	//    }
	//
	//    conn := api.Conn.GraphMgr.GetLeader(graph.Name)
	//
	//    if conn != nil {
	//        api.Logger.Log("Detected New Graph - " + graph.Name + " Leader - OK")
	//        break
	//    }
	//
	//    time.Sleep(time.Second)
	//    times--
	//}

	return resp, err
}

func (api *UltipaAPI) DropGraph(graphName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	resp, err := api.Uql(fmt.Sprintf(`drop().graph("%v")`, graphName), requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (api *UltipaAPI) HasGraph(graphName string, requestConfig *configuration.RequestConfig) (bool, error) {
	graphSets, err := api.ShowGraph(requestConfig)

	if err != nil {
		return false, err
	}

	for _, graph := range graphSets {
		if graph.Name == graphName {
			return true, nil
		}
	}

	return false, nil
}

func (api *UltipaAPI) GetGraph(graphName string, requestConfig *configuration.RequestConfig) (*structs.GraphSet, error) {
	grapSets, err := api.ShowGraph(requestConfig)

	if err != nil {
		return nil, err
	}

	for _, graph := range grapSets {
		if graph.Name == graphName {
			return graph, nil
		}
	}

	return nil, errors.New("graph not found")
}

func (api *UltipaAPI) AlterGraph(oldGraph, newGraph *structs.GraphSet, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`alter().graph("%s").set({name: "%s", description: "%s"})`, oldGraph.Name, newGraph.Name, newGraph.Description)
	// Only modify the description of the graphSet
	if newGraph.Name == "" {
		uql = fmt.Sprintf(`alter().graph("%s").set({description: "%s"})`, oldGraph.Name, newGraph.Description)
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) Truncate(request *structs.Truncate, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := ""
	if request.DbType == nil {
		t := ultipa.DBType_DBGLOBAL
		request.DbType = &t
	}

	schemaName, err := CheckReplaceSchemaPropertyName(request.Schema)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s, graphName = %s", err.Error(), request.GraphName))
	}

	if request.Schema != "" {
		if !(*request.DbType == ultipa.DBType_DBNODE || *request.DbType == ultipa.DBType_DBEDGE) {
			//return nil, fmt.Errorf("to truncate schema, dbType must be DBType_DBNODE or DBType_DBEDGE")
			return nil, fmt.Errorf("to truncate schema, DbType is required in the parameters")
		}

		if request.Schema == "*" {
			request.Schema = `"*"`
		} else {
			request.Schema = "@" + schemaName
		}

	} else {
		if !(*request.DbType == ultipa.DBType_DBGLOBAL) {
			return nil, fmt.Errorf("to truncate graph, dbType must be DBType_DBGLOBAL or nil")
		}
	}

	switch *request.DbType {
	case ultipa.DBType_DBNODE:
		uql = fmt.Sprintf(`truncate().graph("%v").nodes(%v)`, request.GraphName, request.Schema)
	case ultipa.DBType_DBEDGE:
		uql = fmt.Sprintf(`truncate().graph("%v").edges(%v)`, request.GraphName, request.Schema)
	default:
		uql = fmt.Sprintf(`truncate().graph("%v")`, request.GraphName)
	}

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) Compact(graphName string, requestConfig *configuration.RequestConfig) (*http.JobResponse, error) {
	uql := fmt.Sprintf(`compact().graph("%v")`, graphName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return http.GetJobResponseFromUqlResponse(resp)
}

func (api *UltipaAPI) MountGraph(graphName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`mount().graph("%v")`, graphName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) UnmountGraph(graphName string, requestConfig *configuration.RequestConfig) (*http.UQLResponse, error) {
	uql := fmt.Sprintf(`unmount().graph("%v")`, graphName)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
