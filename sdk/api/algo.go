package api

import (
	"bufio"
	"github.com/codingsince1985/checksum"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"github.com/ultipa/ultipa-go-sdk/sdk/structs"
	"io"
	"os"
	"path"
)

func (api *UltipaAPI) ShowAlgo(config *configuration.RequestConfig) ([]*structs.Algo, error) {
	resp, err := api.Uql("show().algo()", config)

	if err != nil {
		return nil, err
	}

	algos, err := resp.Get(0).AsAlgos()

	if err != nil {
		return nil, err
	}

	return algos, nil
}

func (api *UltipaAPI) InstallAlgo(soFilePath string, infoFilePath string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {

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

func (api *UltipaAPI) UninstallAlgo(algoName string, config *configuration.RequestConfig) (*ultipa.UninstallAlgoReply, error) {

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
		AlgoName: algoName,
	})

	if err != nil {
		return nil, err
	}

	return reply, nil
}
