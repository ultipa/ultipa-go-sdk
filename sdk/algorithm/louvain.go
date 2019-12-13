package algorithm

import "ultipa-go-sdk/rpc"

type louvainParams map[string]string

func NewLouvainParams() louvainParams {
	return louvainParams{
		"edge_property_name":      "",
		"min_modularity_increase": "0.01",
		"phase1_loop":             "5",
	}
}

func StartLouvainTask(client ultipa.UltipaRpcsClient, params louvainParams) *ultipa.TaskReply {

	return StartTask(client, TaskLouvain, params)
}
