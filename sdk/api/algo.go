package api

import (
	"bufio"
	"errors"
	"github.com/codingsince1985/checksum"
	ultipa "github.com/ultipa/ultipa-go-sdk/rpc"
	"github.com/ultipa/ultipa-go-sdk/sdk/configuration"
	"io"
	"os"
	"path"
)

// UninstallHDCAlgo uninstall algo
func (api *UltipaAPI) UninstallHDCAlgo(algoName, hdcName string, config *configuration.RequestConfig) (*ultipa.UninstallAlgoReply, error) {

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

// InstallHDCAlgo  install algos, files : [...so, yml]
func (api *UltipaAPI) InstallHDCAlgo(files []string, hdcName string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {
	if files == nil || len(files) == 0 {
		return nil, errors.New("empty files")
	}
	if len(files) < 2 {
		return nil, errors.New("lack files, At least two files are required: so + yml")
	}

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

	// Send each so/yml file
	for i, file := range files {
		algoFile, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer algoFile.Close()

		algoFileReader := bufio.NewReader(algoFile)
		algoFileMD5, _ := checksum.MD5sum(file)

		isYml := i == len(files)-1

		// Send  file in chunks
		if err := sendFileChunks(algoFileReader, streamClient, path.Base(algoFile.Name()), algoFileMD5, hdcName, isYml); err != nil {
			return nil, err
		}
	}

	reply, err := streamClient.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	// Reply status check, handle empty reply
	if reply.Status == nil {
		return reply, nil
	}

	if reply.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(reply.Status.Msg)
	}

	return reply, nil
}

// sendFileChunks sends the file in chunks to the server
func sendFileChunks(fileReader *bufio.Reader, streamClient ultipa.UltipaControls_InstallAlgoClient, fileName, md5, hdcName string, isYml bool) error {
	chunkSize := 1024 * 1024 * 1 // 1MB

	for {
		chunk := make([]byte, chunkSize)
		n, err := fileReader.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		request := &ultipa.InstallAlgoRequest{
			FileName: fileName,
			Md5:      md5,
			Chunk:    chunk[:n],
		}

		// Only include server name if provided
		if !isYml && hdcName != "" {
			request.WithServer = &ultipa.WithServer{HdcServerName: hdcName}
		}

		if err := streamClient.Send(request); err != nil {
			return err
		}
	}
	return nil
}
