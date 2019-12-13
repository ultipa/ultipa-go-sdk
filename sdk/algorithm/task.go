package algorithm

import (
	"context"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

// TaskType includes TaskLouvain, TaskCC, TaskLPA, TaskPageRank
type TaskType = ultipa.TaskRequest_TASK_TYPE

const (
	// TaskLouvain runs louvain algorithm
	TaskLouvain TaskType = ultipa.TaskRequest_TASK_LOUVAIN
	// TaskCC runs Connected Component algorithm
	TaskCC TaskType = ultipa.TaskRequest_TASK_CC
	// TaskLPA runs Label Propagation algorithm
	TaskLPA TaskType = ultipa.TaskRequest_TASK_LPA
	// TaskPageRank runs Page rank algorithm
	TaskPageRank TaskType = ultipa.TaskRequest_TASK_PAGERANK
)

//StartTask runs a user defined task by algorithm.[TaskName]
func StartTask(client ultipa.UltipaRpcsClient, taskName TaskType, params map[string]string) *ultipa.TaskReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.Task(ctx, &ultipa.TaskRequest{
		TaskOpt:  ultipa.TaskRequest_OPT_START,
		TaskType: taskName,
	})

	if err != nil {
		log.Fatalf("[Error] start task error: %v", err)
	}

	return msg
}

//GetTask run a query to get the tasks status
func GetTask(client ultipa.UltipaRpcsClient) *ultipa.TaskReply {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	msg, err := client.Task(ctx, &ultipa.TaskRequest{
		TaskOpt: ultipa.TaskRequest_OPT_SEARCH,
	})

	if err != nil {
		log.Fatalf("[Error] get task status error: %v", err)
	}

	return msg
}
