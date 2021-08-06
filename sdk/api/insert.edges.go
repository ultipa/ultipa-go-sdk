package api

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) InsertEdgesBatch(table *ultipa.EdgeTable, config *configuration.RequestConfig) (*ultipa.InsertEdgesReply, error) {
	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext()

	resp, err := client.InsertEdges(ctx, &ultipa.InsertEdgesRequest{
		GraphName: conf.CurrentGraph,
		EdgeTable: table,
		Silent:    true,
	})

	return resp, err
}
