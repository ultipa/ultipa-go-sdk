package api

import (
	"errors"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
)

//Backup backup ultipa database data to directory backupToDirectory on server
func (api *UltipaAPI) Backup(backupToDirectory string, req *configuration.RequestConfig) (*http.UQLResponse, error) {
	requestConfig := req
	if req != nil {
		requestConfig = &configuration.RequestConfig{
			GraphName:      "global",
			Timeout:        req.Timeout,
			ClusterId:      req.ClusterId,
			Host:           req.Host,
			UseMaster:      true,
			UseControl:     true,
			RequestType:    req.RequestType,
			Uql:            req.Uql,
			Timezone:       req.Timezone,
			TimezoneOffset: req.TimezoneOffset,
			ThreadNum:      req.ThreadNum,
			MaxPkgSize:     req.MaxPkgSize,
		}
	} else {
		requestConfig = &configuration.RequestConfig{
			GraphName:  "global",
			UseMaster:  true,
			UseControl: true,
		}
	}

	client, err := api.GetControlClient(requestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(req)
	if err != nil {
		return nil, err
	}
	defer cancel()

	resp, err := client.Backup(ctx, &ultipa.BackupRequest{BackupPath: backupToDirectory})
	if err != nil {
		return nil, err
	}

	if resp.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		api.Logger.Log("backup failed : " + resp.Status.Msg)
		return nil, errors.New(resp.Status.Msg)
	}

	return &http.UQLResponse{
		Status: &http.Status{
			Message: resp.Status.Msg,
			Code:    resp.Status.ErrorCode,
		},
	}, nil

}
