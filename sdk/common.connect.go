package sdk

import (
	"ultipa-go-sdk/types"
	"ultipa-go-sdk/types/types_response"
	"ultipa-go-sdk/utils"
)
func (t *Connection) Stat(commonReq *SdkRequest_Common) *types_response.ResStat {
	uql := utils.UQLMAKER{}
	uql.SetCommand(utils.CommandList_stat)

	res := t.UQLListSample(uql.ToString(),  commonReq)
	var newData types_response.Stat
	if res.Status.Code == types.ErrorCode_SUCCESS  {
		datas := res.Data
		for _, data := range *datas{
			newData = types_response.Stat{
				MemUsage: (*data)["memUsage"].(string),
				CpuUsage: (*data)["cpuUsage"].(string),
				ExpiredDate: (*data)["expiredDate"].(string),
			}
			break
		}
	}
	return &types_response.ResStat{
		res.ResWithoutData,
		&newData,
	}
}

func (t *Connection) ClusterInfo(commonReq *SdkRequest_Common) *types_response.ResListClusterInfo {
	t.RefreshRaftLeader("",commonReq)
	res := t.GetLeaderReuqest(commonReq)
	var result = []*types_response.ClusterInfo{}
	if res.Status.ClusterInfo != nil {
		for _, peer := range res.Status.ClusterInfo.RaftPeers{
			info := &types_response.ClusterInfo{
				RaftPeerInfo: peer,
				Stat: &types_response.Stat{
					MemUsage: "",
					CpuUsage: "",
					ExpiredDate: "",
				},
			}
			//log.Printf("------")
			if peer.Status {
				res := t.Stat(&SdkRequest_Common{
					UseHost: peer.Host,
				})
				//v, _ := utils.StructToJSONString(res)
				//log.Printf(v)
				if res.Status.Code == types.ErrorCode_SUCCESS {
					info.Stat = res.Data
				}
			}
			result = append(result, info)
		}
	}

	return &types_response.ResListClusterInfo{
		&types.ResWithoutData{
			Status: &types.Status{
				Code: types.ErrorCode_SUCCESS,
			},
		},
		result,
	}
}

