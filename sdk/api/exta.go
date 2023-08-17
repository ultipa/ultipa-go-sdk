/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  exta
 * @Date: 2022/8/5 7:35 下午
 */

package api

import (
	"bufio"
	"github.com/codingsince1985/checksum"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"io"
	"os"
	"path"
)

func (api *UltipaAPI) InstallExta(extaFilePath string, extaInfoFilePath string, req *configuration.RequestConfig) (*ultipa.InstallExtaReply, error) {

	chunkSize := 1024 * 1024 * 1 // 2MB

	// check file status

	extaFile, err := os.OpenFile(extaFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	extaFileReader := bufio.NewReader(extaFile)

	extaFileMD5, _ := checksum.MD5sum(extaFilePath)

	extaInfoFile, err := os.OpenFile(extaInfoFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	extaInfoFileReader := bufio.NewReader(extaInfoFile)

	extaInfoFileMD5, _ := checksum.MD5sum(extaInfoFilePath)

	client, err := api.GetControlClient(req)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(req)
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

func (api *UltipaAPI) UninstallExta(extaName string, req *configuration.RequestConfig) (*ultipa.UninstallExtaReply, error) {

	client, err := api.GetControlClient(req)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(req)
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
