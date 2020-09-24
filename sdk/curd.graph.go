package sdk

import (
	"strconv"
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

type ResponseGraph struct {
	Id         int64
	Name       string
	TotalNodes int64
	TotalEdges int64
}

func (t *Connection) ListGraph(commonReq *types.Request_Common) []*ResponseGraph {

	var graphs []*ResponseGraph
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listGraph)

	res := t.UQLListSample(uql.ToString(), commonReq)

	for _, g := range *res.Data {
		var newGraph ResponseGraph
		graph := *g
		newGraph.Id, _ = strconv.ParseInt(graph["id"].(string), 10, 64)
		newGraph.TotalNodes, _ = strconv.ParseInt(graph["totalNodes"].(string), 10, 64)
		newGraph.TotalEdges, _ = strconv.ParseInt(graph["totalEdges"].(string), 10, 64)
		newGraph.Name, _ = graph["graph"].(string)
		graphs = append(graphs, &newGraph)
	}

	return graphs
}

type ResponseCreateGraph struct {
	*types.ResWithoutData
}

func (t *Connection) CreateGraph(name string, commonReq *types.Request_Common) *ResponseCreateGraph {

	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_createGraph)
	uql.AddParam("name", "\""+name+"\"", false)

	res := t.UQLListSample(uql.ToString(), commonReq)

	return &ResponseCreateGraph{
		res.ResWithoutData,
	}
}

type ResponseDropGraph struct {
	*types.ResWithoutData
}

func (t *Connection) DropGraph(name string, commonReq *types.Request_Common) *ResponseDropGraph {

	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_dropGraph)
	uql.AddParam("name", "\""+name+"\"", false)

	res := t.UQLListSample(uql.ToString(), commonReq)

	return &ResponseDropGraph{
		res.ResWithoutData,
	}
}

type ResponseGetGraph struct {
	*types.ResWithoutData
	Graph *ResponseGraph
}

func (t *Connection) GetGraph(name string, commonReq *types.Request_Common) *ResponseGetGraph {

	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_getGraph)
	uql.AddParam("name", "\""+name+"\"", false)

	res := t.UQL(uql.ToString(), commonReq)

	graph := *res.Data.Values
	var newGraph ResponseGraph
	newGraph.Id, _ = strconv.ParseInt(graph["id"].(string), 10, 64)
	newGraph.TotalNodes, _ = strconv.ParseInt(graph["totalNodes"].(string), 10, 64)
	newGraph.TotalEdges, _ = strconv.ParseInt(graph["totalEdges"].(string), 10, 64)
	newGraph.Name, _ = graph["graph"].(string)

	return &ResponseGetGraph{
		res.ResWithoutData,
		&newGraph,
	}
}
