package api

import (
	"fmt"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/printers"
	"ultipa-go-sdk/sdk/structs"
)

func (api *UltipaAPI) ListAlgo(req *configuration.RequestConfig) ([]*structs.Algo, error) {

	resp, err := api.UQL("show().algo()", req)

	if err != nil {
		return nil, err
	}

	table, err := resp.Get(0).AsTable()

	if err != nil {
		return nil, err
	}

	var algos []*structs.Algo
	algoDatas := table.ToKV()

	for _, algoData := range algoDatas {

		algo, err := structs.NewAlgo(algoData.Data["name"].(string), algoData.Data["param"].(string))

		if err != nil {
			printers.PrintError(fmt.Sprint(err.Error(), algoData))
			continue
		}

		algos = append(algos, algo)
	}

	return algos, nil
}
