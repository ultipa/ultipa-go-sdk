package api

import (
	"bufio"
	"github.com/codingsince1985/checksum"
	"io"
	"os"
	"path"
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ShowAlgo(req *configuration.RequestConfig) ([]*structs.Algo, error) {

	resp, err := api.UQL("show().algo()", req)

	if err != nil {
		return nil, err
	}

	algos, err := resp.Get(0).AsAlgos()

	if err != nil {
		return nil, err
	}

	return algos, nil
}

func (api *UltipaAPI) InstallAlgo(algoFilePath string, algoInfoFilePath string, req *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {

	chunkSize := 1024 * 1024 * 1 // 2MB

	// check file status

	algoFile, err := os.OpenFile(algoFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoFileReader := bufio.NewReader(algoFile)

	algoFileMD5, _ := checksum.MD5sum(algoFilePath)

	algoInfoFile, err := os.OpenFile(algoInfoFilePath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoInfoFileReader := bufio.NewReader(algoInfoFile)

	algoInfoFileMD5, _ := checksum.MD5sum(algoInfoFilePath)

	client, err := api.GetControlClient(req)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(req)
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
			FileName: path.Base(algoFile.Name()),
			Md5:      algoFileMD5,
			Chunk:    chunk[:n],
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

	return reply, nil

}

func (api *UltipaAPI) UninstallAlgo(algoName string, req *configuration.RequestConfig) (*ultipa.UninstallAlgoReply, error) {

	client, err := api.GetControlClient(req)

	if err != nil {
		return nil, err
	}

	ctx, cancel, err := api.Pool.NewContext(req)
	if err != nil {
		return nil, err
	}
	defer cancel()

	reply, err := client.UninstallAlgo(ctx, &ultipa.UninstallAlgoRequest{
		AlgoName: algoName,
	})

	if err != nil {
		return nil, err
	}

	return reply, nil
}
