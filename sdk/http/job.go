package http

import "fmt"

type JobResponse struct {
	JobId uint32
}

func GetJobResponseFromUqlResponse(response *UQLResponse) (*JobResponse, error) {
	if response == nil {
		return nil, fmt.Errorf("cannot convert nil UQLResponse to JobResponse")
	}
	table, err := response.Alias(RESP_JOB_KEY).AsTable()
	if err != nil {
		return nil, fmt.Errorf("parse uqlResponse to jobResponse error : %v", err)
	}

	values := table.ToKV()
	jobId := values[0].Get("new_job_id").(uint32)

	return &JobResponse{jobId}, nil
}
