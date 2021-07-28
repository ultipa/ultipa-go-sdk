package api

import (
	"strconv"
	"ultipa-go-sdk/sdk/configuration"
	"ultipa-go-sdk/sdk/http"
	"ultipa-go-sdk/sdk/utils"
)

func  (api *UltipaAPI) ListNodeSchema(config *configuration.RequestConfig) ( *http.ResponseNodeSchemas,  error) {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_listNodeSchema)
	res, err := api.UQL(uql.ToString(), config)
	if err != nil {
		return nil, err
	}
	table, err := res.GetSingleTable()
	if err != nil {
		return nil, err
	}
	var schemas []*http.ResponseSchema
	if !res.Status.IsSuccess() {
		return &http.ResponseNodeSchemas{
			Status: res.Status,
			Schemas: schemas,
		}, nil
	}
	values := table.ToKV()
	for _, v := range values {
		totalNodes, _ := strconv.ParseInt(v.Get("totalNodes").(string), 10, 64)
		totalEdges, _ := strconv.ParseInt(v.Get("totalEdges").(string), 10, 64)

		schemas = append(schemas, &http.ResponseSchema{
			Name:       v.Get("name").(string),
			Description: v.Get("description").(string),
			Properties: nil,
			TotalNodes: totalNodes,
			TotalEdges: totalEdges,
		})
	}
	return &http.ResponseNodeSchemas{
		Status: res.Status,
		Schemas: schemas,
	}, nil
}
