package api

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"

	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) DownloadAlgoResultFile(fileName string, jobId string, requestConfig *configuration.RequestConfig, receive func(data []byte) error) error {
	var err error

	err, files := api.getFilesByJobId(jobId, requestConfig)
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasSuffix(file, fileName) {
			fileName = file
		}
	}

	client, err := api.GetControlClient(requestConfig)
	if err != nil {
		return err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
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
		} else if record.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
			return errors.New(record.Status.Msg)
		}
		err = receive(record.Chunk)
		if err != nil {
			return err
		}
	}

	return err
}

func (api *UltipaAPI) DownloadAllAlgoResultFile(jobId string, requestConfig *configuration.RequestConfig, receive func(data []byte, fileName string) error) error {
	var err error

	client, err := api.GetControlClient(requestConfig)
	if err != nil {
		return err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
	if err != nil {
		return err
	}
	defer cancel()

	err, files := api.getFilesByJobId(jobId, requestConfig)
	if err != nil {
		return err
	}

	for _, file := range files {
		resp, err := client.DownloadFile(ctx, &ultipa.DownloadFileRequest{
			FileName: file,
			//TaskId:   jobId,
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
			} else if record.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
				return err
			}

			fileName := filepath.Base(file)
			err = receive(record.Chunk, fileName)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (api *UltipaAPI) getFilesByJobId(jobId string, requestConfig *configuration.RequestConfig) (error, []string) {
	var files []string
	jobs, err := api.ShowJob(jobId, requestConfig)
	if err != nil {
		return fmt.Errorf("show job error: %v", err), files
	}
	if len(jobs) == 0 {
		return errors.New("job not found"), files
	}

	job := jobs[0]
	if job.Result == nil {
		return errors.New("job result is empty"), files

	}

	i := 0
	for {
		key := "output_file" + strconv.Itoa(i)
		if i == 0 {
			key = "output_file"
		}

		if file, ok := job.Result[key]; ok {
			files = append(files, file)
			i++
		} else {
			break
		}
	}

	if len(files) == 0 {
		return errors.New("empty files"), nil
	}

	return nil, files
}
