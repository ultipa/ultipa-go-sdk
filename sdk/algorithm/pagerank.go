package algorithm

import "ultipa-go-sdk/sdk"

// NewPageRankParams create Page Rank params for Page Rank task
func NewPageRankParams() TaskValues {
	return TaskValues{
		"loop_num":          "5",
		"page_rank_damping": "0.8",
		"page_rank_default": "1",
	}
}

// StartPageRankTask starts a Page Rank task
func StartPageRankTask(client sdk.Client, params TaskValues) *TaskReply {
	return StartTask(client, TaskPageRank, params)
}
