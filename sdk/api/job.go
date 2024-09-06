package api

import (
	"fmt"

	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowJob(jobId string, requestConfig *configuration.RequestConfig) (jobs []*structs.Job, err error) {
	uql := fmt.Sprintf("show().job(%v)", jobId)

	resp, err := api.Uql(uql, requestConfig)

	if err != nil {
		return nil, err
	}

	jobs, err = resp.Alias(http.RESP_JOB_KEY).AsJobs()

	return jobs, err
}

func (api *UltipaAPI) ClearJob(jobId string, status structs.TaskStatus, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf("clear().job(%v)", jobId)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *UltipaAPI) StopJob(jobId string, status structs.TaskStatus, config *configuration.RequestConfig) (resp *http.UQLResponse, err error) {
	uql := fmt.Sprintf("stop().job(%v)", jobId)

	resp, err = api.Uql(uql, config)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
