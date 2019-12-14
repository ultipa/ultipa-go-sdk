package algorithm

import "ultipa-go-sdk/sdk"

// NewKHopParams create KHop params for Page Rank task
func NewKHopParams() TaskValues {
	return TaskValues{
		"depth": "1",
	}
}

// StartKHopTask starts a N Hop task
func StartKHopTask(client sdk.Client, params TaskValues) *TaskReply {
	return StartTask(client, TaskKHop, params)
}
