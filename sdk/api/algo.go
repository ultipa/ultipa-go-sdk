package api

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/codingsince1985/checksum"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/http"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
)

// Deprecated: 5.0 not support, should use ShowHDCAlgo
func (api *UltipaAPI) ShowAlgo(config *configuration.RequestConfig) ([]*structs.Algo, error) {
	resp, err := api.Uql("show().algo()", config)
	if err != nil {
		return nil, err
	}

	algos, err := resp.Alias(http.RESP_ALGOS_KEY).AsAlgos()

	if err != nil {
		return nil, err
	}

	return algos, nil
}

// InstallAlgo install algo
func (api *UltipaAPI) InstallAlgo(soFilePath, infoFilePath, hdcName string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {

	chunkSize := 1024 * 1024 * 1 // 2MB

	// check file status

	algoFile, err := os.OpenFile(soFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoFileReader := bufio.NewReader(algoFile)

	algoFileMD5, _ := checksum.MD5sum(soFilePath)

	algoInfoFile, err := os.OpenFile(infoFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoInfoFileReader := bufio.NewReader(algoInfoFile)

	algoInfoFileMD5, _ := checksum.MD5sum(infoFilePath)

	client, err := api.GetControlClient(config)
	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return nil, err
	}
	defer cancel()

	streamClient, err := client.InstallAlgo(ctx)

	if err != nil {
		return nil, err
	}

	// send algo so file

	for {

		chunk := make([]byte, chunkSize)
		n, err := algoFileReader.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		err = streamClient.Send(&ultipa.InstallAlgoRequest{
			FileName:   path.Base(algoFile.Name()),
			Md5:        algoFileMD5,
			Chunk:      chunk[:n],
			WithServer: &ultipa.WithServer{HdcServerName: hdcName},
		})

		if err != nil {
			return nil, err
		}
	}

	// send algo info file
	for {
		chunk := make([]byte, chunkSize)

		n, err := algoInfoFileReader.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		err = streamClient.Send(&ultipa.InstallAlgoRequest{
			FileName: path.Base(algoInfoFile.Name()),
			Md5:      algoInfoFileMD5,
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

	// reply status 暂时没有初始化，返回nil 按成功处理
	if reply.Status == nil {
		return reply, nil
	}

	if reply.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(reply.Status.Msg)
	}

	return reply, nil

}

// UninstallAlgo uninstall algo
func (api *UltipaAPI) UninstallAlgo(algoName, hdcName string, config *configuration.RequestConfig) (*ultipa.UninstallAlgoReply, error) {

	client, err := api.GetControlClient(config)
	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return nil, err
	}
	defer cancel()

	reply, err := client.UninstallAlgo(ctx, &ultipa.UninstallAlgoRequest{
		AlgoName:   algoName,
		WithServer: &ultipa.WithServer{HdcServerName: hdcName},
	})

	if err != nil {
		return nil, err
	}

	if reply.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(reply.Status.Msg)
	}

	return reply, nil
}

// RollbackHDCAlgo Rollback HDC Algo
func (api *UltipaAPI) RollbackHDCAlgo(algoName, hdcName string, config *configuration.RequestConfig) (*ultipa.RollbackAlgoReply, error) {

	client, err := api.GetControlClient(config)
	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(config)
	if err != nil {
		return nil, err
	}
	defer cancel()

	reply, err := client.RollbackAlgo(ctx, &ultipa.RollbackAlgoRequest{
		AlgoName:   algoName,
		WithServer: &ultipa.WithServer{HdcServerName: hdcName},
	})

	if err != nil {
		return nil, err
	}

	if reply.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(reply.Status.Msg)
	}

	return reply, nil
}

// GetAlgo get a specific algorithm by name
func (api *UltipaAPI) GetAlgo(algoName string, config *configuration.RequestConfig) (*structs.Algo, error) {
	algos, err := api.ShowAlgo(config)
	if err != nil {
		return nil, err
	}

	for _, algo := range algos {
		if algo.Name == algoName {
			return algo, nil
		}
	}
	return nil, fmt.Errorf("algo %v not found", algoName)
}
