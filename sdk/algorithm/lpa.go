package algorithm

import "ultipa-go-sdk/rpc"

// NewLPAParams create LPA params for LPA task
func NewLPAParams() TaskValues {
	return TaskValues{
		"loop_num":           "5",
		"node_property_name": "",
	}
}

// StartLPATask starts a LPA task
func StartLPATask(client ultipa.UltipaRpcsClient, params TaskValues) *ultipa.TaskReply {
	return StartTask(client, TaskLouvain, params)
}
