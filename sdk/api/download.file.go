package api

import (
	"errors"
	"io"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) DownloadAlgoResultFile(fileName string, taskId string, requestConfig *configuration.RequestConfig, receive func(data []byte) error) error {
	var err error

	client := api.Conn.GetControlClient()

	ctx, cancel, err := api.Conn.NewContext(requestConfig)
	if err != nil {
		return err
	}
	defer cancel()

	resp, err := client.DownloadFile(ctx, &ultipa.DownloadFileRequest{
		FileName: fileName,
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

func (api *UltipaAPI) DownloadAllAlgoResultFile(taskId string, requestConfig *configuration.RequestConfig, receive func(data []byte, fileName string) error) error {
	var err error

	client := api.Conn.GetControlClient()

	ctx, cancel, err := api.Conn.NewContext(requestConfig)
	if err != nil {
		return err
	}
	defer cancel()

	tasks, err := api.ShowTask(taskId, structs.TaskstatusAll, requestConfig)
	if err != nil {
		return errors.New("get task failed, " + err.Error())
	}

	files, err := tasks[0].GetTaskFileName()
	if err != nil {
		return err
	}

	for _, file := range files {
		resp, err := client.DownloadFile(ctx, &ultipa.DownloadFileRequest{
			FileName: file,
			//TaskId:   taskId,
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
