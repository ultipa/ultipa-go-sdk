package structs

//type TaskStatus int
//
//const (
//    TaskstatusAll TaskStatus = iota
//    TaskstatusPending
//    TaskStatusComputing
//    TaskStatusWriting
//    TaskStatusFailed
//    TaskStatusDone
//    TaskStatusStopped
//)

//var taskStatusDescriptions = map[TaskStatus]string{
//    TaskstatusAll:       "*",
//    TaskstatusPending:   "pending",
//    TaskStatusComputing: "computing",
//    TaskStatusWriting:   "writing",
//    TaskStatusFailed:    "failed",
//    TaskStatusDone:      "done",
//    TaskStatusStopped:   "stopped",
//}

//func (ts TaskStatus) String() string {
//    return taskStatusDescriptions[ts]
//}

//type Task struct {
//    Param    map[string]string `json:"param"`
//    TaskInfo TaskInfo          `json:"task_info"`
//    Result   map[string]string `json:"result"`
//    ErrorMsg string            `json:"error_msg"`
//}

// Job represents the job data structure
type Job struct {
	Id        string            `json:"job_id"`     // "job_id" field
	GraphName string            `json:"graph_name"` // "graph_name" field
	Type      string            `json:"type"`       // "type" field
	Query     string            `json:"query"`      // "uql" field
	Status    string            `json:"status"`     // "status" field
	ErrMsg    string            `json:"err_msg"`    // "err_msg" field
	Result    map[string]string `json:"result"`     // "result" field
	StartTime string            `json:"start_time"` // "start_time" field, converted to time.Time type
	EndTime   string            `json:"end_time"`   // "end_time" field, converted to time.Time type
	Progress  string            `json:"progress"`   // "progress" field
}

//type ReturnType struct {
//    IsRealtime      bool `json:"is_realtime"`
//    IsVisualization bool `json:"is_visualization"`
//    IsWriteBack     bool `json:"is_write_back"`
//}
//
//func (t *Task) GetTaskFileName() ([]string, error) {
//    resultFiles, ok := t.Result["result_files"]
//    if !ok {
//        return nil, errors.New("get task fileName error")
//    }
//    return strings.Split(resultFiles, ","), nil
//}
