package algorithm

import (
	"context"
	"log"
	"time"
	"ultipa-go-sdk/rpc"
)

// TaskValues stores params of a task
type TaskValues map[string]string

// TaskType includes TaskLouvain, TaskCC, TaskLPA, TaskPageRank
type TaskType = ultipa.TaskRequest_TASK_TYPE

type TaskReply = ultipa.TaskReply

const (
	// TaskLouvain runs louvain algorithm
	TaskLouvain TaskType = ultipa.TaskRequest_TASK_LOUVAIN
	// TaskCC runs Connected Component algorithm
	TaskCC TaskType = ultipa.TaskRequest_TASK_CC
	// TaskLPA runs Label Propagation algorithm
	TaskLPA TaskType = ultipa.TaskRequest_TASK_LPA
	// TaskPageRank runs Page rank algorithm
	TaskPageRank TaskType = ultipa.TaskRequest_TASK_PAGERANK
	// TaskKHop runs KHop algorithm
	TaskKHop TaskType = ultipa.TaskRequest_TASK_KHOP
)

//StartTask runs a user defined task by algorithm.[TaskName]
func StartTask(client ultipa.UltipaRpcsClient, taskName TaskType, params map[string]string) *TaskReply {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	_p := []*ultipa.Value{}

	for k, v := range params {
		_p = append(_p, &ultipa.Value{Key: k, Value: v})
	}

	msg, err := client.Task(ctx, &ultipa.TaskRequest{
		TaskOpt:  ultipa.TaskRequest_OPT_START,
		TaskType: taskName,
		Params:   _p,
	})

	if err != nil {
		log.Printf("[Error] start task error: %v", "connection refused")
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
		log.Printf("[Error] get task status error: %v", err)
	}

	return msg
}
