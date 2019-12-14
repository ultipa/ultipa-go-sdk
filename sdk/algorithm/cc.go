package algorithm

import "ultipa-go-sdk/sdk"

// StartCCTask starts a Page Rank task
func StartCCTask(client sdk.Client) *TaskReply {
	return StartTask(client, TaskCC, TaskValues{})
}
