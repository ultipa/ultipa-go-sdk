package api

import (
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ListAlgo(req *configuration.RequestConfig) ([]*structs.Algo, error) {

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
