/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  exta
 * @Date: 2022/8/5 7:35 下午
 */

package api

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codingsince1985/checksum"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"io"
	"os"
	"path"
)

func (api *UltipaAPI) InstallExta(soFilePath string, infoFilePath string, requestConfig *configuration.RequestConfig) (*ultipa.InstallExtaReply, error) {

	chunkSize := 1024 * 1024 * 1 // 2MB

	// check file status

	extaFile, err := os.OpenFile(soFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	extaFileReader := bufio.NewReader(extaFile)

	extaFileMD5, _ := checksum.MD5sum(soFilePath)

	extaInfoFile, err := os.OpenFile(infoFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	extaInfoFileReader := bufio.NewReader(extaInfoFile)

	extaInfoFileMD5, _ := checksum.MD5sum(infoFilePath)

	client, err := api.GetControlClient(requestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	streamClient, err := client.InstallExta(ctx)

	if err != nil {
		return nil, err
	}

	// send exta so file

	for {

		chunk := make([]byte, chunkSize)
		n, err := extaFileReader.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		err = streamClient.Send(&ultipa.InstallExtaRequest{
			FileName: path.Base(extaFile.Name()),
			Md5:      extaFileMD5,
			Chunk:    chunk[:n],
		})

		if err != nil {
			return nil, err
		}
	}

	// send exta info file
	for {
		chunk := make([]byte, chunkSize)

		n, err := extaInfoFileReader.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		err = streamClient.Send(&ultipa.InstallExtaRequest{
			FileName: path.Base(extaInfoFile.Name()),
			Md5:      extaInfoFileMD5,
			Chunk:    chunk[:n],
		})

		if err != nil {
			return nil, err
		}
	}

	reply, err := streamClient.CloseAndRecv()

	if err != nil {
		return nil, err
	}

	return reply, nil

}

func (api *UltipaAPI) UninstallExta(extaName string, requestConfig *configuration.RequestConfig) (*ultipa.UninstallExtaReply, error) {

	client, err := api.GetControlClient(requestConfig)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(requestConfig)
	if err != nil {
		return nil, err
	}
	defer cancel()

	reply, err := client.UninstallExta(ctx, &ultipa.UninstallExtaRequest{
		ExtaName: extaName,
	})

	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (api *UltipaAPI) ShowExta(config *configuration.RequestConfig) ([]*structs.Exta, error) {
	var resp *http.UQLResponse
	var err error
	var extas []*structs.Exta

	resp, err = api.Uql(fmt.Sprintf(`show().exta()`), config)
	if err != nil {
		return nil, err
	}
	if resp.Status.Code != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(resp.Status.Message)
	}

	extas, err = resp.Alias(http.RESP_EXTAS_KEY).AsExta()

	if len(extas) == 0 {
		return nil, err
	}

	return extas, err
}
