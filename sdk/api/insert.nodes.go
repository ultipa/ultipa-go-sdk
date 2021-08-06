package api

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/configuration"
)

func (api *UltipaAPI) InsertNodesBatch(table *ultipa.NodeTable, config *configuration.RequestConfig) (*ultipa.InsertNodesReply, error) {

	client, conf, err := api.GetClient(config)

	if err != nil {
		return nil, err
	}

	ctx, _ := api.Pool.NewContext()

	resp, err := client.InsertNodes(ctx, &ultipa.InsertNodesRequest{
		GraphName: conf.CurrentGraph,
		NodeTable: table,
		Silent:    true,
	})

	return resp, err
}
