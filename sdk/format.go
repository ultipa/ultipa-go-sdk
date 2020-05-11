package sdk

import (
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
	var values map[string]string
	for _, _v := range node.GetValues() {
		values[_v.GetKey()] = _v.GetValue()
	}
	newNode := Node{}
	newNode.ID = node.GetId()
	newNode.Values = values
	return &newNode
}

func FormatEdges(edges []*ultipa.Edge) []*Edge {
	var newEdges []*Edge
	for _, ev := range edges {
		newEdges = append(newEdges, FormatEdge(ev))
	}
	return newEdges
}

func FormatEdge(edge *ultipa.Edge) *Edge {
	var values map[string]string
	for _, ev := range edge.GetValues() {
		values[ev.GetKey()] = ev.GetValue()
	}
	newEdge := Edge{}
	newEdge.ID = edge.GetId()
	newEdge.From = edge.GetFromId()
	newEdge.To = edge.GetToId()
	newEdge.Values = values
	return &newEdge
}

func FormatStatus(status *ultipa.Status) *Status {
	newStatus := Status{
		Code:    ErrorCode_SUCCESS,
		Message: "",
	}
	if status.GetErrorCode() != ultipa.ErrorCode_SUCCESS {
		newStatus.Code = ErrorCode_FAILED
		newStatus.Message = status.GetMsg()
	}
	return &newStatus
}

func TableToValues(table *Table) *map[string][]string {
	var res map[string][]string
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
