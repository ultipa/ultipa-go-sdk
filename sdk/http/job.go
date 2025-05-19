package http

import "fmt"

type JobResponse struct {
	JobId     uint32
	Statistic *Statistic
	Status    *Status
}

func GetJobResponseFromUqlResponse(response *Response) (*JobResponse, error) {
	if response == nil {
		return nil, fmt.Errorf("cannot convert nil Response to JobResponse")
	}
	table, err := response.Alias(RESP_RESULT_KEY).AsTable()
	if err != nil {
		return nil, fmt.Errorf("parse uqlResponse to jobResponse error : %v", err)
	}

	values := table.ToKV()
	jobId := values[0].Get("new_job_id").(uint32)

	return &JobResponse{JobId: jobId, Statistic: response.Statistic, Status: response.Status}, nil
}
