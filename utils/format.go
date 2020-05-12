package utils

import (
	"fmt"
	"strings"
	ultipa "ultipa-go-sdk/rpc"
)

func FormatPaths(paths []*ultipa.Path) Paths {
	var ps Paths
	for _, v := range paths {
		var newPath Path
		for _, nv := range v.Nodes {
			newPath.Nodes = append(newPath.Nodes, FormatNode(nv))
		}
		for _, ev := range v.Edges {
			newPath.Edges = append(newPath.Edges, FormatEdge(ev))
		}
		ps = append(ps, &newPath)
	}
	return ps
}
func FormatNodes(nodes []*ultipa.Node) []*Node {
	var newNodes []*Node
	for _, nv := range nodes {
		newNodes = append(newNodes, FormatNode(nv))
	}
	return newNodes
}
func FormatNode(node *ultipa.Node) *Node {
	newNode := Node{}
	newNode.ID = node.GetId()
	newNode.Values = FormatValues(node.GetValues())
	return &newNode
}

func FormatEdges(edges []*ultipa.Edge) []*Edge {
	var newEdges []*Edge
	for _, ev := range edges {
		newEdges = append(newEdges, FormatEdge(ev))
	}
	return newEdges
}
func FormatValues(values []*ultipa.Value)  *map[string]string{
	_values := map[string]string{}
	for _, ev := range values {
		_values[ev.GetKey()] = ev.GetValue()
	}
	return &_values
}
func FormatEdge(edge *ultipa.Edge) *Edge {
	newEdge := Edge{}
	newEdge.ID = edge.GetId()
	newEdge.From = edge.GetFromId()
	newEdge.To = edge.GetToId()
	newEdge.Values = FormatValues(edge.GetValues())
	return &newEdge
}
func FormatStatus(status *ultipa.Status, err error) *Status {
	if err != nil {
		return &Status{
			Code: ErrorCode_FAILED,
			Message: fmt.Sprint(err),
		}
	}
	clusterInfo := ClusterInfo{}
	newStatus := Status{
		Code:    ErrorCode_SUCCESS,
		Message: "",
		ClusterInfo: &clusterInfo,
	}
	newStatus.Code = status.ErrorCode
	if status.GetErrorCode() != ultipa.ErrorCode_SUCCESS {
		newStatus.Message = status.GetMsg()
	}
	_clusterInfo := status.ClusterInfo
	if _clusterInfo != nil {
		clusterInfo.Redirect = _clusterInfo.Redirect
		clusterInfo.RaftPeers = strings.Split(_clusterInfo.RaftPeers, ",")
	}
	return &newStatus
}

func TableToValues(table *Table) *map[string][]string {
	res := map[string][]string{}
	for index, key := range table.Headers {
		res[key] = table.TableRows[index]
	}
	return &res
}

func TableToArray(table *Table) *[]*map[string]string {
	var res []*map[string]string
	for _, rows := range table.TableRows {
		item := map[string]string{}
		for index, row := range rows {
			item[table.Headers[index]] = row
		}
		res = append(res, &item)
	}
	return &res
}
