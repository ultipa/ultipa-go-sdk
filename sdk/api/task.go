package api

import (
	"errors"
	"fmt"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"github.com/ultipa/ultipa-go-sdk/sdk/utils"
)

func (api *UltipaAPI) ShowTask(algoNameOrId string, TaskStatus structs.TaskStatus, config *configuration.RequestConfig) (tasks []*structs.Task, err error) {
	uql := ""
	if len(algoNameOrId) == 0 {
		uql = "show().task()"
	} else if utils.IsAllDigits(algoNameOrId) {
		uql = fmt.Sprintf("show().task(%v)", algoNameOrId)
	} else {
		uql = fmt.Sprintf(`show().task("%v","%v")`, algoNameOrId, TaskStatus.String())
	}

	resp, err := api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	tasks, err = resp.Alias(http.RESP_TASK_KEY).AsTask()

	return tasks, err
}

func (api *UltipaAPI) ClearTask(algoNameOrId string, TaskStatus structs.TaskStatus, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	if len(algoNameOrId) == 0 {
		uql = `clear().task("*")`
	} else if utils.IsAllDigits(algoNameOrId) {
		uql = fmt.Sprintf("clear().task(%v)", algoNameOrId)
	} else {
		uql = fmt.Sprintf(`clear().task("%v","%v")`, algoNameOrId, TaskStatus.String())
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}

func (api *UltipaAPI) StopTask(id string, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := ""
	if len(id) == 0 || id == "*" {
		uql = `stop().task("*")`
	} else if utils.IsAllDigits(id) {
		uql = fmt.Sprintf("stop().task(%v)", id)
	} else {
		uql = fmt.Sprintf(`stop().task("%v")`, id)
	}

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	return resp, nil
}
