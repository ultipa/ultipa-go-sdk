package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/utils"
)

func (t *Connection) Stat(commonReq *types.Request_Common) *types.ResStat {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.UQLCommand_stat)

	res := t.UQLListSample(uql.ToString(), commonReq)
	var newData types.Response_Stat
	if res.Status.Code == types.ErrorCode_SUCCESS {
		datas := res.Data
		for _, data := range *datas {
			newData = types.Response_Stat{
				MemUsage:    (*data)["memUsage"].(string),
				CpuUsage:    (*data)["cpuUsage"].(string),
				ExpiredDate: (*data)["expiredDate"].(string),
			}
			break
		}
	}
	return &types.ResStat{
		res.ResWithoutData,
		&newData,
	}
}

func (t *Connection) ClusterInfo(commonReq *types.Request_Common) *types.ResListClusterInfo {
	t.RefreshRaftLeader("", commonReq)
	res := t.GetLeaderReuqest(commonReq)
	var result = []*types.Response_ClusterInfo{}
	if res.Status.ClusterInfo != nil {
		for _, peer := range res.Status.ClusterInfo.RaftPeers {
			info := &types.Response_ClusterInfo{
				RaftPeerInfo: peer,
				Response_Stat: &types.Response_Stat{
					MemUsage:    "",
					CpuUsage:    "",
					ExpiredDate: "",
				},
			}
			//log.Printf("------")
			if peer.Status {
				res := t.Stat(&types.Request_Common{
					UseHost: peer.Host,
				})
				//v, _ := utils.StructToJSONString(res)
				//log.Printf(v)
				if res.Status.Code == types.ErrorCode_SUCCESS {
					info.Response_Stat = res.Data
				}
			}
			result = append(result, info)
		}
	}

	return &types.ResListClusterInfo{
		&types.ResWithoutData{
			Status: &types.Status{
				Code: types.ErrorCode_SUCCESS,
			},
		},
		result,
	}
}
