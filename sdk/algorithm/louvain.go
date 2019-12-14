package algorithm

import (
	"ultipa-go-sdk/sdk"
)

func NewLouvainParams() TaskValues {
	return TaskValues{
		"edge_property_name":      "",
		"min_modularity_increase": "0.01",
		"phase1_loop":             "5",
	}
}

func StartLouvainTask(client sdk.Client, params TaskValues) *TaskReply {
	return StartTask(client, TaskLouvain, params)
}
