package structs

import (
	"errors"
	"strings"
)

type TaskStatus int

const (
	TaskstatusAll TaskStatus = iota
	TaskstatusPending
	TaskStatusComputing
	TaskStatusWriting
	TaskStatusFailed
	TaskStatusDone
	TaskStatusStopped
)

// Deprecated: 5.0 not support, should use CreateNodeIndex or CreateEdgeIndex
var taskStatusDescriptions = map[TaskStatus]string{
	TaskstatusAll:       "*",
	TaskstatusPending:   "pending",
	TaskStatusComputing: "computing",
	TaskStatusWriting:   "writing",
	TaskStatusFailed:    "failed",
	TaskStatusDone:      "done",
	TaskStatusStopped:   "stopped",
}

func (ts TaskStatus) String() string {
	return taskStatusDescriptions[ts]
}

type Task struct {
	Param    map[string]string `json:"param"`
	TaskInfo TaskInfo          `json:"task_info"`
	Result   map[string]string `json:"result"`
	ErrorMsg string            `json:"error_msg"`
}

type TaskInfo struct {
	TaskID           int        `json:"task_id"`
	ServerID         int        `json:"server_id"`
	AlgoName         string     `json:"algo_name"`
	StartTime        int        `json:"start_time"`
	WritingStartTime int        `json:"writing_start_time"`
	EndTime          int        `json:"end_time"`
	TimeCost         int        `json:"time_cost"`
	TaskStatus       TaskStatus `json:"TASK_STATUS"`
	ReturnType       int        `json:"return_type"`
}

type ReturnType struct {
	IsRealtime      bool `json:"is_realtime"`
	IsVisualization bool `json:"is_visualization"`
	IsWriteBack     bool `json:"is_write_back"`
}

func (t *Task) GetTaskFileName() ([]string, error) {
	resultFiles, ok := t.Result["result_files"]
	if !ok {
		return nil, errors.New("get task fileName error")
	}
	return strings.Split(resultFiles, ","), nil
}
