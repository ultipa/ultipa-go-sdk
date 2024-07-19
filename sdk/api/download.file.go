package api

import (
	"errors"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"io"
)

func (api *UltipaAPI) DownloadAlgoResultFile(fileName string, taskId string, config *configuration.RequestConfig, receive func(data []byte) error) error {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return err
	}
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

	return err
}

func (api *UltipaAPI) DownloadAllAlgoResultFile(taskId string, config *configuration.RequestConfig, receive func(data []byte) error) error {
	tasks, err := api.ShowTask(taskId, structs.TaskstatusAll, config)
	if err != nil {
		return errors.New("get task failed, " + err.Error())
	}

	files, err := tasks[0].GetTaskFileName()
	if err != nil {
		return err
	}

	for _, file := range files {
		err = api.DownloadAlgoResultFile(file, taskId, config, receive)
		if err != nil {
			return err
		}
	}

	return nil
}

func (api *UltipaAPI) DownloadAllAlgoResultFile2(taskId string, config *configuration.RequestConfig, receive func(data []byte, fileName string) error) error {
	var err error

	client, err := api.GetControlClient(config)

	if err != nil {
		return err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return err
	}
	defer cancel()

	tasks, err := api.ShowTask(taskId, structs.TaskstatusAll, config)
	if err != nil {
		return errors.New("get task failed, " + err.Error())
	}

	files, err := tasks[0].GetTaskFileName()
	if err != nil {
		return err
	}

	for _, file := range files {
		resp, err := client.DownloadFileV2(ctx, &ultipa.DownloadFileRequestV2{
			FileName: file,
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
			err = receive(record.Chunk, file)
			if err != nil {
				return err
			}
		}
	}

	return err
}
