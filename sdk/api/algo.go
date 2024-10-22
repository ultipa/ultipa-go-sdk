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

//// InstallHDCAlgos install algos, file: { "soFile": "ymlFile"}
//func (api *UltipaAPI) InstallHDCAlgos(files map[string]string, hdcName string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {
//	if files == nil || len(files) == 0 {
//		return nil, errors.New("empty files")
//	}
//
//	var result *ultipa.InstallAlgoReply
//	var err error
//	for soFile, ymlFile := range files {
//		if result, err = api.InstallHDCAlgo(soFile, ymlFile, hdcName, config); err != nil {
//			return nil, err
//		}
//	}
//
//	return result, nil
//}

// Deprecated:  should use InstallHDCAlgos
func (api *UltipaAPI) InstallHDCAlgo(soFile, ymlFile, hdcName string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {

	chunkSize := 1024 * 1024 * 1 // 2MB

	// check file status

	algoFile, err := os.OpenFile(soFile, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoFileReader := bufio.NewReader(algoFile)

	algoFileMD5, _ := checksum.MD5sum(soFile)

	algoInfoFile, err := os.OpenFile(ymlFile, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	algoInfoFileReader := bufio.NewReader(algoInfoFile)

	algoInfoFileMD5, _ := checksum.MD5sum(ymlFile)

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

	// reply status, if nil , mark success
	if reply.Status == nil {
		return reply, nil
	}

	if reply.Status.ErrorCode != ultipa.ErrorCode_SUCCESS {
		return nil, errors.New(reply.Status.Msg)
	}

	return reply, nil

}

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

// InstallHDCAlgos  install algos, files : [...so, yml]
func (api *UltipaAPI) InstallHDCAlgos(files []string, hdcName string, config *configuration.RequestConfig) (*ultipa.InstallAlgoReply, error) {
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
		algoFile, err := os.OpenFile(file, os.O_RDONLY, 0644)
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
