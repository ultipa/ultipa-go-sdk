package api

import (
	"io"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) DownloadFileV2(fileName string, taskId string, config *configuration.RequestConfig, receive func(data []byte) error) error {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return err
	}

	ctx, cancel := api.Pool.NewContext(config)
	defer cancel()

	resp, err := client.DownloadFileV2(ctx, &ultipa.DownloadFileRequestV2{
		FileName: fileName,
		TaskId:   taskId,
	})

	if err != nil {
		return err
	}

	for {
		record, err := resp.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		err = receive(record.Chunk)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return err
}
