package api

import (
	"errors"
	"ultipa-go-sdk/sdk/connection"
	"ultipa-go-sdk/sdk/utils"
)

/**
* Utils for SDK Clients.
 */

type UQLType = int

const (
	UQLType_Master = 1
	UQLType_Normal = 2
	UQLType_Global = 3
)

func (api *UltipaAPI) RefreshClusterInfo(graphName string) error {
	return api.Pool.RefreshClusterInfo(graphName)
}

func  (api *UltipaAPI) GetConnByUQL(uql string, graphName string) (uqlType UQLType, leader *connection.Connection, followers []*connection.Connection, global *connection.Connection, err error) {

	graph := api.Pool.GraphMgr.GetGraph(graphName)

	if graph == nil {
		err = api.Pool.RefreshClusterInfo(graphName)
		if err != nil {
			return 0, nil, nil, nil, err
		}

	}

	// refresh , but not get graph info
	if graph == nil {
		return 0, nil, nil, nil, errors.New("unavailable to get graph cluster infos : " + graphName)
	}

	leader = api.Pool.GraphMgr.GetLeader(graphName)
	followers = api.Pool.GraphMgr.GetGraph(graphName).Followers
	global, err = api.Pool.GetGlobalMasterConn(nil)

	uqlItem := utils.NewUql(uql)

	uqlType = UQLType_Normal

	if uqlItem.HasWrite() {
		uqlType = UQLType_Master
	}

	if uqlItem.IsGlobal() {
		uqlType = UQLType_Global
	}

	return uqlType, leader, followers, global, err

}